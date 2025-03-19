package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"log"
	"os"
	"strings"
)

func parse(r io.Reader, args ...interface{}) {
	labels := args[0].([]string)
	optT := args[1].(*bool)
	optC := args[2].(*bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		origRoot := sc.Tree()
		fmt.Printf("# Selected clade")
		if len(labels) > 1 {
			fmt.Printf("s")
		}
		fmt.Printf("\n")
		for _, label := range labels {
			t := origRoot.CopyClade()
			fmt.Printf("## ")
			if *optC {
				fmt.Printf("Complement of ")
			}
			fmt.Printf("%s\n", label)
			var nodes []*nwk.Node
			nodes = tree2slice(t, nodes)
			found := false
			var clade *nwk.Node
			for _, node := range nodes {
				if node.Label == label {
					clade = node
					found = true
					break
				}
			}
			if !found {
				log.Fatalf("Couldn't find node %q.\n", label)
			}
			if *optC {
				if clade.Parent == nil {
					t = nil
				} else {
					clade.RemoveClade()
				}
			}
			if *optT {
				if *optC {
					if t != nil {
						fmt.Printf("%s\n", t)
					}
				} else {
					clade = clade.Child
					if clade != nil {
						fmt.Printf("(%s%s;\n", clade, label)
					}
				}
			} else {
				if *optC {
					listLeaves(t)
				} else {
					listLeaves(clade.Child)
				}
			}
		}
	}
}
func tree2slice(v *nwk.Node, ns []*nwk.Node) []*nwk.Node {
	if v == nil {
		return ns
	}
	ns = append(ns, v)
	ns = tree2slice(v.Child, ns)
	ns = tree2slice(v.Sib, ns)
	return ns
}
func listLeaves(v *nwk.Node) {
	if v == nil {
		return
	}
	if v.Child == nil {
		fmt.Printf("%s\n", v.Label)
	}
	listLeaves(v.Child)
	listLeaves(v.Sib)
}
func main() {
	u := "pickle [option]... [foo.nwk]..."
	p := "Pick clades in Newick trees."
	e := "pickle 3,5 foo.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optT := flag.Bool("t", false, "print tree")
	optC := flag.Bool("c", false, "complement")
	flag.Parse()
	if *optV {
		util.PrintInfo("pickle")
	}
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "please enter a clade identifier\n")
		os.Exit(-1)
	}
	labels := strings.Split(args[0], ",")
	clio.ParseFiles(args[1:], parse, labels, optT, optC)
}
