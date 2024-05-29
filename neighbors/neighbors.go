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

func parse(r io.Reader, args ...interface{}) {
	taxdb := args[0].(*tdb.TaxonomyDB)
	optL := args[1].(bool)
	optG := args[2].(bool)
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
		g := taxdb.Accessions(target)
		if len(g) > 0 && g[0] != "-" {
			genomes[target] = g
		}
	}
	for _, neighbor := range neighbors {
		g := taxdb.Accessions(neighbor)
		if len(g) > 0 && g[0] != "-" {
			genomes[neighbor] = g
		}
	}
	if optL {
		w := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', 0)
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
		w.Flush()
	} else {
		mrcaTname := taxdb.Name(mrcaT)
		mrcaAname := taxdb.Name(mrcaA)
		fmt.Printf("# MRCA(targets): %d, %s\n", mrcaT, mrcaTname)
		fmt.Printf("# MRCA(targets+neighbors): %d, %s\n", mrcaA, mrcaAname)
		w := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', 0)
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
			if optG && g == "-" {
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
			if optG && g == "-" {
				continue
			}
			n := taxdb.Name(neighbor)
			fmt.Fprintf(w, "n\t%d\t%s\t%s\n", neighbor, n,
				strings.TrimPrefix(g, " "))
		}
		w.Flush()
	}
}
func main() {
	u := "neighbors [-h] [option]... <db> [targets.txt]..."
	p := "Given a taxonomy database computed with makeNeiDb and " +
		"a set of target taxon-IDs, find their closest " +
		"taxonomic neighbors."
	e := "neighbors neidb targetIds.txt"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optG := flag.Bool("g", false, "genome sequences only")
	optL := flag.Bool("l", false, "list genomes")
	flag.Parse()
	if *optV {
		util.PrintInfo("neighbors")
	}
	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr,
			"please provide a database name\n")
		os.Exit(0)
	}
	taxdb := tdb.OpenTaxonomyDB(files[0])
	files = files[1:]
	clio.ParseFiles(files, parse, taxdb, *optL, *optG)
}
