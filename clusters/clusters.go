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
	Id            int
	Label, Parent string
	Score         float64
	Size          int
	IsTerminal    bool
	Root          *nwk.Node
}

func parse(r io.Reader, args ...interface{}) {
	optM := args[0].(*bool)
	optN := args[1].(*bool)
	optT := args[2].(*bool)
	optS := args[3].(*int)
	optC := args[4].(*bool)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		branchLengths := make(map[int][]float64)
		traverse(tree, branchLengths)
		clusters := make(map[int]*Cluster)
		findCl(tree, branchLengths, optM, clusters)
		for _, cluster := range clusters {
			it := true
			cluster.IsTerminal = isTerminal(cluster.Root,
				clusters, it)
		}
		for k, v := range clusters {
			if v.Size < *optS || (!v.IsTerminal && !*optN) {
				delete(clusters, k)
			}
		}
		if *optT {
			markClusters(tree, clusters)
			fmt.Println(tree)
		} else {
			clusterSlice := make([]*Cluster, 0)
			for _, cluster := range clusters {
				clusterSlice = append(clusterSlice, cluster)
			}
			if *optC {
				slices.SortFunc(clusterSlice, func(a, b *Cluster) int {
					if a.Score != b.Score {
						if a.Score == b.Score {
							return 0
						} else if a.Score < b.Score {
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
					} else if a.Score != b.Score {
						if a.Score == b.Score {
							return 0
						} else if a.Score < b.Score {
							return 1
						}
						return -1
					}
					return a.Id - b.Id
				})
			}
			w := tabwriter.NewWriter(os.Stdout, 2, 1, 1, ' ', 0)
			if len(clusterSlice) > 0 {
				fmt.Fprintf(w, "#Cluster\tParent\tSize\tC\n")
			}
			for _, cluster := range clusterSlice {
				fmt.Fprintf(w, "%s\t%s\t%d\t%.3g\n",
					cluster.Label, cluster.Parent,
					cluster.Size, cluster.Score)
			}
			w.Flush()
		}
	}
}
func traverse(v *nwk.Node, branchLengths map[int][]float64) {
	if v == nil {
		return
	}
	traverse(v.Child, branchLengths)
	if v.Child != nil {
		traverse(v.Child.Sib, branchLengths)
	}
	if v.Parent != nil {
		blv := branchLengths[v.Id]
		blp := branchLengths[v.Parent.Id]
		blp = append(blp, blv...)
		blp = append(blp, v.Length)
		branchLengths[v.Parent.Id] = blp
	}
}
func findCl(v *nwk.Node, branchLengths map[int][]float64,
	optM *bool, clusters map[int]*Cluster) {
	if v == nil {
		return
	}
	findCl(v.Child, branchLengths, optM, clusters)
	findCl(v.Sib, branchLengths, optM, clusters)
	bl := branchLengths[v.Id]
	if len(bl) > 3 && v.Parent != nil {
		q := util.Quartiles(bl)
		t := q.UpperOuterFence
		if *optM {
			t = q.UpperInnerFence
		}
		if v.Length > t {
			cluster := new(Cluster)
			cluster.Id = v.Id
			cluster.Label = v.Label
			if v.Parent != nil {
				cluster.Parent = v.Parent.Label
			}
			cluster.Root = v.Child
			n := 0
			cluster.Size = size(cluster.Root, n)
			r := q.UpperQuartile - q.LowerQuartile
			cluster.Score = v.Length / r
			clusters[cluster.Id] = cluster
		}
	}
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
func markClusters(v *nwk.Node, clusters map[int]*Cluster) {
	if v == nil {
		return
	}
	markClusters(v.Child, clusters)
	markClusters(v.Sib, clusters)
	if clusters[v.Id] != nil {
		v.Label += "c"
	}
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
	m = "print nested clusters (default terminal)"
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
	clio.ParseFiles(files, parse, optM, optN, optT, optS, optC)
}
