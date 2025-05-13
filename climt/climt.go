package main

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
	down := args[1].(bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		root := sc.Tree()
		findStart(root, down, start)
	}
}
func findStart(root *nwk.Node, down bool, start *regexp.Regexp) {
	if root == nil {
		return
	}
	if start.MatchString(root.Label) {
		v := root
		if down {
			children := make([]*nwk.Node, 0)
			np := v.Child
			for np != nil {
				children = append(children, np)
				np = np.Sib
			}
			if len(children) > 0 {
				w := tabwriter.NewWriter(os.Stdout, 0,
					1, 3, ' ', 0)
				fmt.Fprint(w, "# Parent\tChild")
				if len(children) > 1 {
					fmt.Fprint(w, "ren")
				}
				fmt.Fprint(w, "\n")
				fmt.Fprintf(w, "%s\t", v.Label)
				for i, child := range children {
					if i > 0 {
						fmt.Fprint(w, " ")
					}
					fmt.Fprintf(w, "%s", child.Label)
				}
				fmt.Fprint(w, "\n")
				w.Flush()
			}
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
	findStart(root.Child, down, start)
	findStart(root.Sib, down, start)
}
func main() {
	util.SetName("climt")
	u := "climt [option]... v [inputFile]..."
	p := "Climb a phylogenetic tree starting at v."
	e := "climt -r 303 foo.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optD := flag.Bool("d", false, "climb down one level")
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
	clio.ParseFiles(files, scan, start, *optD)
}
