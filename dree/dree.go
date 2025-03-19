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
	flag.Parse()
	if *optV {
		util.PrintInfo("dree")
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
	taxdb := tdb.OpenTaxonomyDB(dbname)
	subtree := taxdb.Subtree(tid)
	hasGenome := make(map[int]bool)
	hasGsub := make(map[int]bool)
	for _, v := range subtree {
		if len(taxdb.Accessions(v)) > 0 {
			hasGenome[v] = true
			hasGsub[v] = true
		}
	}
	for _, v := range subtree {
		if hasGsub[v] {
			u := v
			p := taxdb.Parent(u)
			for u != tid {
				hasGsub[p] = true
				u = p
				p = taxdb.Parent(u)
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
			n := len(taxdb.Accessions(v))
			if !*optG || n > 0 {
				r := taxdb.Rank(v)
				fmt.Fprintf(w, "%d\t%s\t%d", v, r, n)
				if *optN {
					a := taxdb.Name(v)
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
					fmt.Printf(t2, v, taxdb.Name(v))
				}
				if v != tid {
					p := taxdb.Parent(v)
					if p != v {
						fmt.Printf("\t%d -> %d\n", p, v)
					}
				}
			}
		}
		fmt.Printf("}\n")
	}
}
