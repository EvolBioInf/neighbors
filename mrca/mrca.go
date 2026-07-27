package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"log"
	"regexp"
	"sort"
)

func parse(r io.Reader, args ...interface{}) {
	nodeNames := args[0].(*regexp.Regexp)
	optL := args[1].(*bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		nodes := make([]*nwk.Node, 0)
		mrca := 0
		nodes = storeNodes(tree, nodes)
		for i, node := range nodes {
			node.Id = i
		}
		startNodes := []int{}
		for i, node := range nodes {
			if nodeNames.MatchString(node.Label) {
				startNodes = append(startNodes, i)
			}
		}
		if len(startNodes) == 0 {
			log.Fatalf("couldn't find any nodes matching %s",
				nodeNames)
		}
		children := make(map[int]bool)
		parents := make(map[int]bool)
		for _, startNode := range startNodes {
			children[startNode] = true
		}
		desc := make(map[int]map[int]bool)
		for _, startNode := range startNodes {
			desc[startNode] = make(map[int]bool)
			desc[startNode][startNode] = true
		}
		counts := make(map[int]map[int]bool)
		for _, startNode := range startNodes {
			counts[startNode] = make(map[int]bool)
			counts[startNode][startNode] = true
		}
		ns := len(startNodes)
		for true {
			found := false
			for child, _ := range children {
				nd := len(desc[child])
				if ns == nd {
					mrca = child
					found = true
					break
				}
			}
			if found {
				break
			}
			for child, _ := range children {
				if nodes[child].Parent != nil {
					parent := nodes[child].Parent.Id
					if desc[parent] == nil {
						desc[parent] = make(map[int]bool)
					}
					for d, _ := range desc[child] {
						desc[parent][d] = true
					}
					parents[parent] = true
					if counts[parent] == nil {
						counts[parent] = make(map[int]bool)
					}
					counts[parent][child] = true
				}
			}
			reset(children)
			for parent, _ := range parents {
				children[parent] = true
			}
			reset(parents)
		}
		fmt.Printf("MRCA %s\n", nodes[mrca].Label)
		fmt.Printf("#Type\tn")
		if *optL {
			fmt.Printf("\tLabel")
			if len(startNodes) > 1 {
				fmt.Printf("s")
			}
		}
		fmt.Printf("\n")
		fmt.Printf("Start\t%d", len(startNodes))
		if *optL {
			fmt.Printf("\t%s", nodes[startNodes[0]].Label)
			for i := 1; i < len(startNodes); i++ {
				l := nodes[startNodes[i]].Label
				fmt.Printf(" %s", l)
			}
		}
		fmt.Printf("\n")
		foundCA := false
		for k, v := range desc {
			if k != mrca && len(v) > 1 {
				foundCA = true
				break
			}
		}
		if foundCA {
			fmt.Printf("#CA\tn")
			if *optL {
				fmt.Printf("\tLabels")
			}
			fmt.Printf("\n")
			reset(children)
			for _, startNode := range startNodes {
				children[startNode] = true
			}
			done := make(map[int]bool)
			for len(children) > 0 {
				reset(parents)
				for child, _ := range children {
					if nodes[child].Parent != nil {
						parent := nodes[child].Parent.Id
						parents[parent] = true
					}
				}
				keys := sortKeys(parents)
				for _, parent := range keys {
					if len(counts[parent]) > 1 && parent != mrca &&
						!done[parent] {
						fmt.Printf("%s", nodes[parent].Label)
						n := len(desc[parent])
						fmt.Printf("\t%d", n)
						if *optL {
							keys2 := sortKeys(desc[parent])
							fmt.Printf("\t%s", nodes[keys2[0]].Label)
							for i := 1; i < len(keys2); i++ {
								d := keys2[i]
								fmt.Printf(" %s", nodes[d].Label)
							}
						}
						fmt.Printf("\n")
						done[parent] = true
					}
				}
				reset(children)
				for parent, _ := range parents {
					children[parent] = true
				}
				reset(parents)
			}
		}
	}
}
func storeNodes(v *nwk.Node, nodes []*nwk.Node) []*nwk.Node {
	if v == nil {
		return nodes
	}
	nodes = append(nodes, v)
	nodes = storeNodes(v.Child, nodes)
	nodes = storeNodes(v.Sib, nodes)
	return nodes
}
func reset(m map[int]bool) {
	for k, _ := range m {
		delete(m, k)
	}
}
func sortKeys(m map[int]bool) []int {
	keys := make([]int, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
func main() {
	util.SetName("mrca")
	u := "mrca [-v] regex [foo.nwk]..."
	p := "Calculate the MRCA of nodes matching regex."
	e := "mrca 941322 eco7k.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optL := flag.Bool("l", false, "print node labels")
	flag.Parse()
	if *optV {
		util.PrintInfo("mrca")
	}
	args := flag.Args()
	if len(args) == 0 {
		m := "Please provide a regex specifying " +
			"the starting nodes."
		log.Fatal(m)
	}
	nodeNames, err := regexp.Compile(args[0])
	util.Check(err)
	files := args[1:]
	clio.ParseFiles(files, parse, nodeNames, optL)
}
