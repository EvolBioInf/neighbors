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
	list, onlyG, tab bool) {
	mrcaT := taxdb.MRCA(taxa)
	targets := taxdb.Subtree(mrcaT)
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
		mrcaA = taxdb.Parent(mrcaA)
		nodes := taxdb.Subtree(mrcaA)
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
		accessions := taxdb.Accessions(target)
		i := 0
		for _, accession := range accessions {
			if accession != "-" {
				accessions[i] = accession
				i++
			}
		}
		accessions = accessions[:i]
		for i := 0; i < len(accessions); i++ {
			arr := strings.Split(accessions[i], "/")
			accessions[i] = arr[0]
		}
		for i := 0; i < len(accessions); i++ {
			arr := strings.Split(accessions[i], ":")
			if len(arr) > 1 {
				accessions[i] = arr[1]
			}
		}
		if len(accessions) > 0 {
			genomes[target] = accessions
		}
	}
	for _, neighbor := range neighbors {
		accessions := taxdb.Accessions(neighbor)
		i := 0
		for _, accession := range accessions {
			if accession != "-" {
				accessions[i] = accession
				i++
			}
		}
		accessions = accessions[:i]
		for i := 0; i < len(accessions); i++ {
			arr := strings.Split(accessions[i], "/")
			accessions[i] = arr[0]
		}
		for i := 0; i < len(accessions); i++ {
			arr := strings.Split(accessions[i], ":")
			if len(arr) > 1 {
				accessions[i] = arr[1]
			}
		}
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
		sample := "t"
		for _, target := range targets {
			accessions := genomes[target]
			for _, accession := range accessions {
				fmt.Fprintf(w, "%s\t%s\n", sample, accession)
			}
		}
		sample = "n"
		for _, neighbor := range neighbors {
			accessions := genomes[neighbor]
			for _, accession := range accessions {
				fmt.Fprintf(w, "%s\t%s\n", sample, accession)
			}
		}
	} else {
		mrcaTname := taxdb.Name(mrcaT)
		mrcaAname := taxdb.Name(mrcaA)
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
			fmt.Fprintf(w, "%s\t%d\t%s\t%s\n", t, target,
				taxdb.Name(target),
				strings.TrimPrefix(g, " "))
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
			n := taxdb.Name(neighbor)
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
	calcTarNei(taxa, taxdb, optL, optG, optTT)
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
	flag.Parse()
	if *optV {
		util.PrintInfo("neighbors")
	}
	var targets []int
	if *optT != "" {
		ts := strings.Split(*optT, ",")
		for _, t := range ts {
			target, e := strconv.Atoi(t)
			util.Check(e)
			targets = append(targets, target)
		}
	}
	files := flag.Args()
	if len(files) < 1 {
		fmt.Fprintf(os.Stderr, "please enter a datbase")
		os.Exit(1)
	}
	taxdb := tdb.OpenTaxonomyDB(files[0])
	files = files[1:]
	if len(targets) > 0 {
		calcTarNei(targets, taxdb, *optL, *optG, *optTT)
	} else {
		clio.ParseFiles(files, parse, taxdb, *optL,
			*optG, *optTT)
	}
}
