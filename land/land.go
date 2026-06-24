package land

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"strconv"
)

func parse(r io.Reader, args ...interface{}) {
	pr := args[0].(string)
	su := args[1].(string)
	ri := args[2].(bool)
	rl := args[3].(bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		labelTree(tree, 1, pr, su, ri, rl)
		fmt.Println(tree)
	}
}
func labelTree(v *nwk.Node, c int, pr, su string,
	ri, rl bool) int {
	if v == nil {
		return c
	}
	l := v.Label
	if v.Child != nil {
		if ri {
			l = ""
		} else {
			l = pr + strconv.Itoa(c) + su
			c++
		}
	} else if rl {
		l = ""
	}
	v.Label = l
	c = labelTree(v.Child, c, pr, su, ri, rl)
	c = labelTree(v.Sib, c, pr, su, ri, rl)
	return c
}
func Run() {
	util.SetName("land")
	u := "land [option]... [treeFile]..."
	p := "Label the internal nodes in Newick trees."
	e := "land -p n foo.nwk"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optP = flag.String("p", "", "prefix")
	var optS = flag.String("s", "", "suffix")
	var optR = flag.Bool("r", false, "remove labels "+
		"from internal nodes")
	var optL = flag.Bool("l", false, "remove leaf labels")
	flag.Parse()
	if *optV {
		util.PrintInfo("land")
	}
	files := flag.Args()
	clio.ParseFiles(files, parse, *optP, *optS, *optR, *optL)
}
