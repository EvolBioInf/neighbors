package ants

import (
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

func Run() {
	util.SetName("ants")
	u := "ants [option] <taxon-ID> <db>"
	p := "Get a taxon's ancestors."
	e := "ants 9606 neidb"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	var optR = flag.String("r", "", "name of remote database (implies remote execution)")
	flag.Parse()
	if *optV {
		util.PrintInfo("ants")
	}
	args := flag.Args()
	if *optR != "" {
		resp := util.SendGetRequest(
			"api/v2/programs/ants",
			os.Args[1:],
			[]string{},
			map[string]string{"db": *optR},
		)
		fmt.Print(resp)
		return
	}
	if len(args) != 2 {
		m := "please provide a taxon and a database"
		fmt.Fprintf(os.Stderr, "%s\n", m)
		os.Exit(-1)
	}
	tid, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	db := args[1]
	taxdb := tdb.OpenTaxonomyDB(db)
	var ants []int
	ants = append(ants, tid)
	a, err := taxdb.Parent(tid)
	util.Check(err)
	for tid != a {
		ants = append(ants, a)
		tid = a
		a, err = taxdb.Parent(tid)
		util.Check(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer w.Flush()
	fmt.Fprintf(w, "# Back\tID\tName\tRank\n")
	for i := len(ants) - 1; i >= 0; i-- {
		a := ants[i]
		n, err := taxdb.Name(a)
		util.Check(err)
		r, err := taxdb.Rank(a)
		util.Check(err)
		fmt.Fprintf(w, "  %d\t%d\t%s\t%s\n", i, a, n, r)
	}
}
