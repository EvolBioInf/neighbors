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
	"sort"
	"strings"
	"text/tabwriter"
)

type Count struct {
	label, parent string
	vn, vt        int
	sv, dp        float64
}
type countsSlice []*Count

func (c countsSlice) Len() int {
	return len(c)
}
func (c countsSlice) Less(i, j int) bool {
	if c[i].sv == c[j].sv {
		return c[i].label < c[j].label
	}
	return c[i].sv > c[j].sv
}
func (c countsSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func parse(r io.Reader, args ...interface{}) {
	optA := args[0].(*bool)
	optN := args[1].(*string)
	optT := args[2].(*string)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		counts := make(map[int]*Count)
		traverseTree(tree, counts, *optN, *optT)
		nt := counts[tree.Id].vt
		nn := counts[tree.Id].vn
		for _, count := range counts {
			van := nn - count.vn
			count.sv = float64(count.vt+van) /
				float64(nt+nn) * 100
		}
		cs := make([]*Count, 0)
		for _, count := range counts {
			cs = append(cs, count)
		}
		sort.Sort(countsSlice(cs))
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "#Clade\tTargets\tNeighbors\tSplit (%)\t"+
			"Parent\tDist(Parent)\n")
		i := 0
		for ; i < len(cs) && cs[0].sv == cs[i].sv; i++ {
			pl := cs[i].parent
			if pl == "" {
				pl = "-"
			}
			fmt.Fprintf(w, "%s\t%d\t%d\t%.1f\t%s\t%g\n",
				cs[i].label, cs[i].vt, cs[i].vn,
				cs[i].sv, pl, cs[i].dp)
		}
		if *optA {
			for ; i < len(cs); i++ {
				pl := cs[i].parent
				if pl == "" {
					pl = "-"
				}
				fmt.Fprintf(w, "%s\t%d\t%d\t%.1f\t%s\t%g\n",
					cs[i].label, cs[i].vt, cs[i].vn,
					cs[i].sv, pl, cs[i].dp)
			}
		}
		w.Flush()
	}
}
func traverseTree(v *nwk.Node, counts map[int]*Count,
	np, tp string) {
	if v == nil {
		return
	}
	count := new(Count)
	count.label = v.Label
	if v.Parent != nil {
		count.dp = v.Length
		count.parent = v.Parent.Label
	}
	counts[v.Id] = count
	traverseTree(v.Child, counts, np, tp)
	traverseTree(v.Sib, counts, np, tp)
	if v.Child == nil {
		if strings.HasPrefix(v.Label, np) {
			counts[v.Id].vn = 1.0
		} else if strings.HasPrefix(v.Label, tp) {
			counts[v.Id].vt = 1.0
		} else {
			log.Fatalf("%q is neither target nor neighbor",
				v.Label)
		}
	} else {
		if v.Label == "" {
			log.Fatal("please label internal nodes " +
				"using land")
		}
	}
	if v.Parent != nil {
		counts[v.Parent.Id].vt += counts[v.Id].vt
		counts[v.Parent.Id].vn += counts[v.Id].vn
	}
}
func main() {
	u := "fintac [option]... [foo.nwk]..."
	p := "Find target clade in Newick tree."
	e := "fintac foo.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optA := flag.Bool("a", false, "all splits (default maximal)")
	optN := flag.String("n", "n", "neighbor prefix")
	optT := flag.String("t", "t", "target prefix")
	flag.Parse()
	if *optV {
		util.PrintInfo("fintac")
	}
	if *optN == *optT {
		log.Fatal("Please use distinct target " +
			"and neighbor prefixes.")
	}
	files := flag.Args()
	clio.ParseFiles(files, parse, optA, optN, optT)
}
