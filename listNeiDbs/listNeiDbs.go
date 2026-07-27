package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"os"
	"text/tabwriter"
)

func main() {
	util.SetName("listNeiDbs")
	u := "listNeiDbs [option]"
	p := "Print all of never's available databases."
	e := "listNeiDbs"
	clio.Usage(u, p, e)
	var optV = flag.Bool("v", false, "version")
	flag.Parse()
	if *optV {
		util.PrintInfo("listNeiDbs")
	}
	a := make(map[string]string)
	a["plain_data"] = "true"
	rawRes := util.SendGetRequest("api/v2/databases", nil, nil, a)
	var databases []string
	json.Unmarshal([]byte(rawRes), &databases)
	w := tabwriter.NewWriter(os.Stdout, 2, 1, 2, ' ', 0)
	fmt.Fprintf(w, "#Nr\tName\n")
	for i, database := range databases {
		fmt.Fprintf(w, "%d\t%s\n", i+1, database)
	}
	w.Flush()
}
