package climt

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"text/tabwriter"
)

func scan(r io.Reader, args ...interface{}) {
	start := args[0].(*regexp.Regexp)
	down := args[1].(int)
	delim := args[2].(string)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		root := sc.Tree()
		findStart(root, down, delim, start)
	}
}
func findStart(root *nwk.Node, down int, delim string,
	start *regexp.Regexp) {
	if root == nil {
		return
	}
	if start.MatchString(root.Label) {
		v := root
		if down > 0 {
			sib := v.Sib
			v.Sib = nil
			writeTree(v, down, delim, 0)
			v.Sib = sib
		} else {
			ancestors := make([]*nwk.Node, 0)
			np := v
			for np != nil {
				ancestors = append(ancestors, np)
				np = np.Parent
			}
			cumLen := 0.0
			np = v
			for np != nil {
				cumLen += np.Length
				np = np.Parent
			}
			w := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)
			fmt.Fprint(w, "# Up\tNode\tBranch Length\t"+
				"Cumulative Branch Length\n")
			n := len(ancestors)
			for i := n - 1; i >= 0; i-- {
				f := cumLen
				s := 15.0
				x := math.Pow(10, s)
				f = math.Round(f*x) / x
				if math.Signbit(f) {
					f *= -1.0
				}
				fmt.Fprintf(w, "%d\t%s\t%g\t%g\n",
					i,
					ancestors[i].Label,
					ancestors[i].Length,
					f)
				if i > 0 {
					cumLen -= ancestors[i-1].Length
				}
			}
			w.Flush()
		}
	}
	findStart(root.Child, down, delim, start)
	findStart(root.Sib, down, delim, start)
}
func writeTree(v *nwk.Node, down int, delim string,
	level int) {
	if v == nil || level > down {
		return
	}
	for i := 0; i < level; i++ {
		fmt.Printf("%s", delim)
	}
	fmt.Printf("%s\n", v.Label)
	writeTree(v.Child, down, delim, level+1)
	writeTree(v.Sib, down, delim, level)
}
func Run() {
	util.SetName("climt")
	u := "climt [option]... v [inputFile]..."
	p := "Climb a phylogenetic tree starting at v."
	e := "climt -r 303 foo.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optD := flag.Int("d", 0, "number of levels climbed "+
		"down")
	optDD := flag.String("D", "   ", "delimiter for "+
		"indentation in down climb")
	optR := flag.Bool("r", false, "v is a regular expression")
	flag.Parse()
	if *optV {
		util.PrintInfo("climt")
	}
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("please provide a starting node")
	}
	expr := args[0]
	if !*optR {
		expr = "^" + expr + "$"
	}
	start, err := regexp.Compile(expr)
	util.Check(err)
	files := args[1:]
	clio.ParseFiles(files, scan, start, *optD, *optDD)
}
