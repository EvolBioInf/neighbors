package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
)

func main() {
	u := "dree [-h] [option]... <taxon-ID> <db>"
	p := "Get the taxonomy rooted on a specific taxon."
	e := "dree -n -g 207598 neidb | dot -T x11"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optN := flag.Bool("n", false,
		"print names instead of taxon-IDs")
	optG := flag.Bool("g", false,
		"only taxa with genome sequences")
	optL := flag.Bool("l", false, "list taxa")
	optLL := flag.String("L", "", util.LevelMsg())
	flag.Parse()
	if *optV {
		util.PrintInfo("dree")
	}
	levels := make(map[string]bool)
	knowns := make(map[string]bool)
	for _, level := range tdb.AssemblyLevels() {
		knowns[level] = true
	}
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
	tokens := flag.Args()
	if len(tokens) != 2 {
		fmt.Fprintf(os.Stderr,
			"please provide a taxon-ID and a database\n")
		os.Exit(0)
	}
	tid, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatalf("couldn't convert %q", tokens[0])
	}
	dbname := tokens[1]
	neidb := tdb.OpenTaxonomyDB(dbname)
	subtree, err := neidb.Subtree(tid)
	util.Check(err)
	hasGenome := make(map[int]bool)
	hasGsub := make(map[int]bool)
	for _, v := range subtree {
		acc, err := neidb.Accessions(v)
		util.Check(err)
		acc, err = neidb.FilterAccessions(acc, levels)
		util.Check(err)
		if len(acc) > 0 {
			hasGenome[v] = true
			hasGsub[v] = true
		}
	}
	for _, v := range subtree {
		if hasGsub[v] {
			u := v
			p, err := neidb.Parent(u)
			util.Check(err)
			for u != tid {
				hasGsub[p] = true
				u = p
				p, err = neidb.Parent(u)
				util.Check(err)
			}
		}
	}
	sort.Ints(subtree)
	if *optL {
		w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
		fmt.Fprint(w, "# Taxid\tRank\tGenomes")
		if *optN {
			fmt.Fprint(w, "\tName")
		}
		fmt.Fprint(w, "\n")
		for _, v := range subtree {
			numAcc := 0
			acc, err := neidb.Accessions(v)
			util.Check(err)
			acc, err = neidb.FilterAccessions(acc, levels)
			util.Check(err)
			numAcc = len(acc)
			if !*optG || numAcc > 0 {
				r, err := neidb.Rank(v)
				util.Check(err)
				fmt.Fprintf(w, "%d\t%s\t%d", v, r, numAcc)
				if *optN {
					a, err := neidb.Name(v)
					util.Check(err)
					fmt.Fprintf(w, "\t%s", a)
				}
				fmt.Fprintf(w, "\n")
			}
		}
		w.Flush()
	} else {
		t1 := "\t%d [color=\"lightsalmon\",style=filled]\n"
		t2 := "\t%d [label=\"%s\"]\n"
		fmt.Printf("digraph g {\n\trankdir=LR\n")
		for _, v := range subtree {
			if !*optG || (*optG && hasGsub[v]) {
				if hasGenome[v] {
					fmt.Printf(t1, v)
				}
				if *optN {
					name, err := neidb.Name(v)
					util.Check(err)
					fmt.Printf(t2, v, name)
				}
				if v != tid {
					p, err := neidb.Parent(v)
					util.Check(err)
					if p != v {
						fmt.Printf("\t%d -> %d\n", p, v)
					}
				}
			}
		}
		fmt.Printf("}\n")
	}
}
