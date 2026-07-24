package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"strings"
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
	var res []string
	json.Unmarshal([]byte(rawRes), &res)
	fmt.Println(strings.Join(res, ", "))
}
