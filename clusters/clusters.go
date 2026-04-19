package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"os"
	"slices"
	"text/tabwriter"
)

type Cluster struct {
	Size       int
	IsTerminal bool
	C          float64
	Id         int
}

func parse(r io.Reader, args ...interface{}) {
	optM := args[0].(*bool)
	optB := args[1].(*string)
	optN := args[2].(*bool)
	optT := args[3].(*bool)
	optS := args[4].(*int)
	optC := args[5].(*bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		var nodes []*nwk.Node
		branchLengths := make(map[int][]float64)
		clusters := make(map[int]*Cluster)
		nodes = traverse(tree, nodes)
		for i, node := range nodes {
			node.Id = i
		}
		for _, node := range nodes {
			if node.Parent == nil {
				continue
			}
			bln := branchLengths[node.Id]
			blp := branchLengths[node.Parent.Id]
			blp = append(blp, bln...)
			blp = append(blp, node.Length)
			branchLengths[node.Parent.Id] = blp
		}
		for i, bl := range branchLengths {
			if bl == nil {
				continue
			}
			if len(bl) > 3 && nodes[i].Parent != nil {
				q := util.Quartiles(bl)
				t := q.UpperOuterFence
				if *optM {
					t = q.UpperInnerFence
				}
				if nodes[i].Length > t {
					cluster := &Cluster{Id: i}
					clusters[i] = cluster
				}
			}
		}
		for _, cluster := range clusters {
			v := nodes[cluster.Id].Child
			cluster.Size = size(v, 0)
			cluster.IsTerminal = isTerminal(v, clusters, true)
			bl := branchLengths[cluster.Id]
			q := util.Quartiles(bl)
			r := q.UpperQuartile - q.LowerQuartile
			l := nodes[cluster.Id].Length
			cluster.C = l / r
		}
		for k, cluster := range clusters {
			if cluster.Size < *optS ||
				(!cluster.IsTerminal && !*optN) {
				delete(clusters, k)
			}
		}
		if *optB != "" {
			w := tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', 0)
			fmt.Fprintf(w, "#Len\tType\n")
			for i, bl := range branchLengths {
				v := nodes[i]
				if v.Label == *optB {
					if v.Parent != nil {
						pl := v.Length
						fmt.Fprintf(w, "%.3g\tp\n", pl)
					}
					for _, dl := range bl {
						fmt.Fprintf(w, "%.3g\td\n", dl)
					}
				}
			}
			w.Flush()
		} else if *optT {
			for _, cluster := range clusters {
				nodes[cluster.Id].Label += "c"
			}
			fmt.Println(tree)
		} else {
			var clusterSlice []*Cluster
			for _, cluster := range clusters {
				clusterSlice = append(clusterSlice, cluster)
			}
			if *optC {
				slices.SortFunc(clusterSlice, func(a, b *Cluster) int {
					if a.C != b.C {
						if a.C == b.C {
							return 0
						} else if a.C < b.C {
							return 1
						}
						return -1

					} else if a.Size != b.Size {
						return b.Size - a.Size
					}
					return a.Id - b.Id
				})
			} else {
				slices.SortFunc(clusterSlice, func(a, b *Cluster) int {
					if a.Size != b.Size {
						return b.Size - a.Size
					} else if a.C != b.C {
						if a.C == b.C {
							return 0
						} else if a.C < b.C {
							return 1
						}
						return -1
					}
					return a.Id - b.Id
				})
			}
			w := tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', 0)
			if len(clusters) > 0 {
				fmt.Fprintf(w, "#Cluster\tParent\tSize\tC\n")
			}
			for _, cluster := range clusterSlice {
				label := nodes[cluster.Id].Label
				parent := nodes[cluster.Id].Parent.Label
				fmt.Fprintf(w, "%s\t%s\t%d\t%.3g\n",
					label, parent,
					cluster.Size, cluster.C)
			}
			w.Flush()
		}
	}
}
func traverse(v *nwk.Node, nodes []*nwk.Node) []*nwk.Node {
	if v == nil {
		return nodes
	}
	nodes = traverse(v.Child, nodes)
	nodes = traverse(v.Sib, nodes)
	nodes = append(nodes, v)
	return nodes
}
func size(v *nwk.Node, c int) int {
	if v == nil {
		return c
	}
	c = size(v.Child, c)
	c = size(v.Sib, c)
	if v.Child == nil {
		c++
	}
	return c
}
func isTerminal(v *nwk.Node, clusters map[int]*Cluster,
	it bool) bool {
	if v == nil {
		return it
	}
	it = isTerminal(v.Child, clusters, it)
	it = isTerminal(v.Sib, clusters, it)
	if clusters[v.Id] != nil {
		it = false
	}
	return it
}
func main() {
	util.SetName("clusters")
	u := "clusters [option]... [foo.nwk]..."
	p := "Find nodes with long parental " +
		"branches in Newick trees."
	e := "clusters -n myTree.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	m := "include mild clusters (default extreme)"
	optM := flag.Bool("m", false, m)
	m = "print branch lengths of given node (default clusters)"
	optB := flag.String("b", "", m)
	m = "nested clusters (default terminal)"
	optN := flag.Bool("n", false, m)
	m = "print tree with marked clusters"
	optT := flag.Bool("t", false, m)
	m = "minimum size of cluster"
	optS := flag.Int("s", 10, m)
	m = "sort by cluster score, C (default sort by size)"
	optC := flag.Bool("c", false, m)
	flag.Parse()
	if *optV {
		util.PrintInfo("clusters")
	}
	files := flag.Args()
	clio.ParseFiles(files, parse, optM, optB,
		optN, optT, optS, optC)
}
