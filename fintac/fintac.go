package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Count struct {
	label, parent string
	vn, vt, vu    int
	sv, dp        float64
}

func parse(r io.Reader, args ...interface{}) {
	optA := args[0].(*bool)
	tregex := args[1].(*regexp.Regexp)
	nregex := args[2].(*regexp.Regexp)
	uregex := args[3].(*regexp.Regexp)
	neidb := args[4].(*tdb.TaxonomyDB)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		counts := make(map[int]*Count)
		traverseTree(tree, counts, tregex, nregex, uregex, neidb)
		nt := float64(counts[tree.Id].vt)
		nn := float64(counts[tree.Id].vn)
		nu := float64(counts[tree.Id].vu)
		for _, count := range counts {
			van := nn - float64(count.vn)
			vau := nu - float64(count.vu)
			vt := float64(count.vt)
			count.sv = (vt + van + math.Log(vau+1.0)) /
				(nt + nn + math.Log(nu+1.0)) * 100.0
		}
		cs := make([]*Count, 0)
		for _, count := range counts {
			cs = append(cs, count)
		}
		slices.SortFunc(cs, func(a, b *Count) int {
			if a.sv < b.sv {
				return 1
			} else if a.sv > b.sv {
				return -1
			}
			return strings.Compare(a.label, b.label)
		})
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprint(w, "#Clade\tTargets\tNeighbors\tUnknowns\t"+
			"Split (%)\tParent\tDist(Parent)\n")
		i := 0
		for ; i < len(cs) && cs[0].sv == cs[i].sv; i++ {
			pl := cs[i].parent
			if pl == "" {
				pl = "-"
			}
			fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%.2f\t%s\t%g\n",
				cs[i].label, cs[i].vt, cs[i].vn, cs[i].vu,
				cs[i].sv, pl, cs[i].dp)
		}
		if *optA {
			for ; i < len(cs); i++ {
				pl := cs[i].parent
				if pl == "" {
					pl = "-"
				}
				fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%.2f\t%s\t%g\n",
					cs[i].label, cs[i].vt, cs[i].vn, cs[i].vu,
					cs[i].sv, pl, cs[i].dp)
			}
		}
		w.Flush()
	}
}
func traverseTree(v *nwk.Node, counts map[int]*Count,
	tregex, nregex, uregex *regexp.Regexp,
	neidb *tdb.TaxonomyDB) {
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
	traverseTree(v.Child, counts, tregex, nregex, uregex,
		neidb)
	traverseTree(v.Sib, counts, tregex, nregex, uregex,
		neidb)
	if v.Child == nil {
		isTar := false
		isUnk := false
		isNei := false
		var err error
		isTar, err = hatch(tregex, v.Label, neidb)
		util.Check(err)
		if uregex != nil && uregex.MatchString(v.Label) {
			isUnk = true
		}
		if nregex == nil {
			if !isTar && !isUnk {
				isNei = true
			}
		} else {
			isNei, err = hatch(nregex, v.Label, neidb)
			util.Check(err)
		}
		dc := 0
		if isTar {
			dc++
		}
		if isUnk {
			dc++
		}
		if isNei {
			dc++
		}
		if dc == 0 {
			fmt.Fprintf(os.Stderr, "WARNING[fintac]: %q "+
				"is neither target, neighbor, nor "+
				"unknown\n",
				v.Label)
		}
		if dc > 1 {
			log.Fatalf("%q is ambiguous: t %t, n %t, u %t",
				v.Label, isTar, isNei, isUnk)
		}
		if isTar {
			counts[v.Id].vt = 1
		} else if isNei {
			counts[v.Id].vn = 1
		} else if isUnk {
			counts[v.Id].vu = 1
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
		counts[v.Parent.Id].vu += counts[v.Id].vu
	}
}
func hatch(regex *regexp.Regexp, label string,
	neidb *tdb.TaxonomyDB) (bool, error) {
	isMatch := regex.MatchString(label)
	if !isMatch && neidb != nil {
		taxid := -1
		str := strings.Split(label, "_")[0]
		taxid, err := strconv.Atoi(str)
		if err != nil {
			return false, err
		}
		for taxid != 1 {
			parent, err := neidb.Parent(taxid)
			if err != nil {
				return false, err
			}
			taxid = parent
			str := strconv.Itoa(taxid) + "_"
			isMatch = regex.MatchString(str)
			if isMatch {
				break
			}
		}
	}
	return isMatch, nil
}
func main() {
	util.SetName("fintac")
	u := "fintac [option]... [foo.nwk]..."
	p := "Find target clade for taxa identified by " +
		"regular expressions in Newick tree."
	e := "fintac -t \"^991910\" -u \"^562\" eco7k.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optA := flag.Bool("a", false, "all splits (default maximal)")
	optT := flag.String("t", "^t", "target")
	optU := flag.String("u", "", "unknown: either target "+
		"or neighbor")
	optN := flag.String("n", "", "neighbor "+
		"(default complement of -t and -u)")
	optNN := flag.String("N", "", "neighbors datbase to activate "+
		"hierarchical matching for targets and neighbors")
	flag.Parse()
	if *optV {
		util.PrintInfo("fintac")
	}
	tregex, err := regexp.Compile(*optT)
	util.Check(err)
	var nregex *regexp.Regexp
	if *optN != "" {
		nregex, err = regexp.Compile(*optN)
		util.Check(err)
	}
	var uregex *regexp.Regexp
	if *optU != "" {
		uregex, err = regexp.Compile(*optU)
		util.Check(err)
	}
	var neidb *tdb.TaxonomyDB
	if *optNN != "" {
		neidb, err = tdb.OpenTaxonomyDBcheck(*optNN)
		util.Check(err)
	}
	files := flag.Args()
	clio.ParseFiles(files, parse, optA, tregex, nregex, uregex,
		neidb)
}
