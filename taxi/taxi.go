package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"os"
	"sort"
	"text/tabwriter"
)

func main() {
	u := "taxi [option] <scientific-name> <db>"
	p := "Take user from scientific name to taxon-ID."
	e := "taxi \"homo sapiens\" neidb"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optS = flag.Bool("s", false, "substring match")
	flag.Parse()
	if *optV {
		util.PrintInfo("taxi")
	}
	args := flag.Args()
	m := "please provide a taxon and a database"
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(-1)
	}
	name := args[0]
	db := args[1]
	if *optS {
		name = fmt.Sprintf("%%%s%%", name)
	}
	taxdb := tdb.OpenTaxonomyDB(db)
	taxa := taxdb.Taxids(name)
	if len(taxa) == 0 {
		return
	}
	sort.Ints(taxa)
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "# ID\tParent\tName\n")
	for _, taxon := range taxa {
		name := taxdb.Name(taxon)
		p := taxdb.Parent(taxon)
		fmt.Fprintf(w, "  %d\t%d\t%s\n", taxon, p, name)
	}
}
