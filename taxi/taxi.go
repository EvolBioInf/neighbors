package main

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	u := "taxi [option] <scientific-name> <db>"
	p := "Take user from scientific name to taxon-ID."
	e := "taxi \"homo sapiens\" neidb"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optE = flag.Bool("e", false, "exact match")
	var optL = flag.Int("l", -1, "limit output to <= l taxids")
	var optO = flag.Int("o", 0, "offset into taxid list")
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
	if !*optE {
		na := strings.Fields(name)
		name = strings.Join(na, "% %")
		name = "%" + name + "%"
	}
	taxdb := tdb.OpenTaxonomyDB(db)
	taxa, err := taxdb.Taxids(name, *optL, *optO)
	util.Check(err)
	if len(taxa) == 0 {
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "# ID\tParent\tName\n")
	for _, taxon := range taxa {
		name, err := taxdb.Name(taxon)
		util.Check(err)
		p, err := taxdb.Parent(taxon)
		util.Check(err)
		fmt.Fprintf(w, "  %d\t%d\t%s\n", taxon, p, name)
	}
}
