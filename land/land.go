package main

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
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		labelTree(tree, 1, pr, su)
		fmt.Println(tree)
	}
}
func labelTree(v *nwk.Node, c int, pr, su string) int {
	if v == nil {
		return c
	}
	l := v.Label
	if v.Child != nil {
		l = pr + strconv.Itoa(c) + su
		c++
	}
	v.Label = l
	c = labelTree(v.Child, c, pr, su)
	c = labelTree(v.Sib, c, pr, su)
	return c
}
func main() {
	u := "land [option]... [treeFile]..."
	p := "Label the internal nodes in Newick trees."
	e := "land -p n foo.nwk"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optP = flag.String("p", "", "prefix")
	var optS = flag.String("s", "", "suffix")
	flag.Parse()
	if *optV {
		util.PrintInfo("land")
	}
	files := flag.Args()
	clio.ParseFiles(files, parse, *optP, *optS)
}
