package taxi

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func Run() {
	util.SetName("taxi")
	u := "taxi [option] <scientific-name|taxid> <db>"
	p := "Take user from scientific name to taxon-ID " +
		"or vice versa."
	e := "taxi \"homo sapiens\" neidb"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optE = flag.Bool("e", false, "exact match")
	var optL = flag.Int("l", -1, "limit output to <= l taxids")
	var optO = flag.Int("o", 0, "offset into taxid list")
	var optT = flag.Bool("t", false, "taxid instead of name")
	var optR = flag.Bool("r", false, "remote execution (implies db)")
	var optDD = flag.String("D", "", "name of remote database (implies remote execution)")
	flag.Parse()
	if *optV {
		util.PrintInfo("taxi")
	}
	if *optR || *optDD != "" {
		misc := map[string]string{}
		if *optDD != "" {
			misc["db"] = *optDD
		}
		resp := util.SendGetRequest(
			"api/v2/programs/taxi",
			util.SanitizeArguments(
				os.Args[1:],
				[]util.Option{
					{Name: "r", WithValue: false},
					{Name: "D", WithValue: true}}),
			[]string{},
			misc,
		)
		fmt.Print(resp)
		return
	}
	args := flag.Args()
	m := "please provide a taxon and a database"
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(1)
	}
	label := args[0]
	db := args[1]
	name := ""
	taxid := 0
	if *optT {
		x, err := strconv.ParseInt(label, 0, 0)
		util.Check(err)
		taxid = int(x)
	} else {
		name = label
		if !*optE {
			na := strings.Fields(name)
			name = strings.Join(na, "% %")
			name = "%" + name + "%"
		}
	}
	taxdb, err := tdb.OpenTaxonomyDBcheck(db)
	util.Check(err)
	taxa := []int{taxid}
	if !*optT {
		taxa, err = taxdb.Taxids(name, *optL, *optO)
		util.Check(err)
	}
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
