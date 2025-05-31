package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

func calcTarNei(taxa []int, taxdb *tdb.TaxonomyDB,
	list, onlyG, tab bool, levels map[string]bool) {
	mrcaT, err := taxdb.MRCA(taxa)
	util.Check(err)
	parent, err := taxdb.Parent(mrcaT)
	util.Check(err)
	if parent == mrcaT {
		m := "no neighbors as %d is the most " +
			"recenct common ancestor of " +
			"the targets and root"
		log.Fatalf(m, mrcaT)
	}
	targets, err := taxdb.Subtree(mrcaT)
	util.Check(err)
	newTargets := make(map[int]bool)
	sort.Ints(taxa)
	l := len(taxa)
	for _, t := range targets {
		i := sort.SearchInts(taxa, t)
		if !(i < l && taxa[i] == t) {
			newTargets[t] = true
		}
	}
	var neighbors []int
	mrcaA := mrcaT
	for len(neighbors) == 0 {
		mrcaA, err = taxdb.Parent(mrcaA)
		util.Check(err)
		nodes, err := taxdb.Subtree(mrcaA)
		util.Check(err)
		sort.Ints(targets)
		l = len(targets)
		for _, node := range nodes {
			i := sort.SearchInts(targets, node)
			if !(i < l && node == targets[i]) {
				if node != mrcaA {
					neighbors = append(neighbors, node)
				}
			}
		}
	}
	genomes := make(map[int][]string)
	for _, target := range targets {
		accessions, err := taxdb.Accessions(target)
		util.Check(err)
		accessions, err = taxdb.FilterAccessions(accessions, levels)
		util.Check(err)
		if len(accessions) > 0 {
			genomes[target] = accessions
		}
	}
	for _, neighbor := range neighbors {
		accessions, err := taxdb.Accessions(neighbor)
		util.Check(err)
		accessions, err = taxdb.FilterAccessions(accessions,
			levels)
		util.Check(err)
		if len(accessions) > 0 {
			genomes[neighbor] = accessions
		}
	}
	var w io.Writer
	if tab {
		w = os.Stdout
	} else {
		w = tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', 0)
	}
	if list {
		fmt.Fprintf(w, "# Sample\tAccession\n")
		var acc []string
		for _, target := range targets {
			accessions := genomes[target]
			for _, accession := range accessions {
				acc = append(acc, accession)
			}
		}
		sort.Strings(acc)
		sample := "t"
		for _, a := range acc {
			fmt.Fprintf(w, "%s\t%s\n", sample, a)
		}
		acc = acc[:0]
		for _, neighbor := range neighbors {
			accessions := genomes[neighbor]
			for _, accession := range accessions {
				acc = append(acc, accession)
			}
		}
		sort.Strings(acc)
		sample = "n"
		for _, a := range acc {
			fmt.Fprintf(w, "%s\t%s\n", sample, a)
		}
	} else {
		mrcaTname, err := taxdb.Name(mrcaT)
		util.Check(err)
		mrcaAname, err := taxdb.Name(mrcaA)
		util.Check(err)
		fmt.Printf("# MRCA(targets): %d, %s\n", mrcaT, mrcaTname)
		fmt.Printf("# MRCA(targets+neighbors): %d, %s\n", mrcaA,
			mrcaAname)
		fmt.Fprint(w, "# Type\tTaxon-ID\tName\tGenomes\n")
		for _, target := range targets {
			t := "t"
			if newTargets[target] {
				t = "tt"
			}
			g := "-"
			if len(genomes[target]) > 0 {
				g = strings.Join(genomes[target], "|")
			}
			if onlyG && g == "-" {
				continue
			}
			name, err := taxdb.Name(target)
			util.Check(err)
			g = strings.TrimPrefix(g, " ")
			fmt.Fprintf(w, "%s\t%d\t%s\t%s\n", t, target, name, g)
		}
		sort.Ints(neighbors)
		for _, neighbor := range neighbors {
			g := "-"
			if len(genomes[neighbor]) > 0 {
				g = strings.Join(genomes[neighbor], "|")
			}
			if onlyG && g == "-" {
				continue
			}
			n, err := taxdb.Name(neighbor)
			util.Check(err)
			fmt.Fprintf(w, "n\t%d\t%s\t%s\n", neighbor, n,
				strings.TrimPrefix(g, " "))
		}
	}
	if !tab {
		tw := w.(*tabwriter.Writer)
		tw.Flush()
	}
}
func parse(r io.Reader, args ...interface{}) {
	taxdb := args[0].(*tdb.TaxonomyDB)
	optL := args[1].(bool)
	optG := args[2].(bool)
	optTT := args[3].(bool)
	levels := args[4].(map[string]bool)
	var taxa []int
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		s := sc.Text()
		if s == "" || s[0] == '#' {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("couldn't convert %q", s)
		}
		taxa = append(taxa, i)
	}
	calcTarNei(taxa, taxdb, optL, optG, optTT, levels)
}
func main() {
	u := "neighbors [-h] [option]... <db> [targets.txt]..."
	p := "Given a taxonomy database computed with makeNeiDb and " +
		"a set of target taxon IDs, find their closest " +
		"taxonomic neighbors."
	e := "neighbors -t 9606 neidb"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optG := flag.Bool("g", false, "genome sequences only")
	optL := flag.Bool("l", false, "list genomes")
	optT := flag.String("t", "", "comma-delimited targets")
	optTT := flag.Bool("T", false, "tab-delimited output "+
		"(default pretty-printing)")
	optLL := flag.String("L", "", util.LevelMsg())
	flag.Parse()
	if *optV {
		util.PrintInfo("neighbors")
	}
	var targets []int
	if *optT != "" {
		if len(flag.Args()) > 1 {
			m := "please use either -t or input files, " +
				"not both"
			fmt.Fprintln(os.Stderr, m)
			os.Exit(1)
		}
		ts := strings.Split(*optT, ",")
		for _, t := range ts {
			target, e := strconv.Atoi(t)
			util.Check(e)
			targets = append(targets, target)
		}
	}
	knowns := util.AssemblyLevels()
	levels := make(map[string]bool)
	var requests []string
	if *optLL != "" {
		requests = strings.Split(*optLL, ",")
	}
	if len(requests) > 0 {
		for _, request := range requests {
			if knowns[request] {
				levels[request] = true
			} else {
				log.Fatalf("unknown level %q", request)
			}
		}
	} else {
		levels = knowns
	}
	files := flag.Args()
	if len(files) < 1 {
		fmt.Fprintf(os.Stderr, "please enter a datbase\n")
		os.Exit(1)
	}
	taxdb := tdb.OpenTaxonomyDB(files[0])
	files = files[1:]
	if len(targets) > 0 {
		calcTarNei(targets, taxdb, *optL, *optG,
			*optTT, levels)
	} else {
		clio.ParseFiles(files, parse, taxdb, *optL,
			*optG, *optTT, levels)
	}
}
