package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/util"
	"github.com/evolbioinf/nwk"
	"io"
	"log"
	"os"
	"strings"
)

func parse(r io.Reader, args ...interface{}) {
	names := args[0].(map[string]string)
	sc := nwk.NewScanner(r)
	for sc.Scan() {
		tree := sc.Tree()
		rename(tree, names)
		fmt.Println(tree)
	}
}
func rename(v *nwk.Node, names map[string]string) {
	if v == nil {
		return
	}
	newLabel := names[v.Label]
	if newLabel != "" {
		v.Label = newLabel
	}
	rename(v.Child, names)
	rename(v.Sib, names)
}
func main() {
	util.SetName("reno")
	u := "reno [-v] fon.txt [phy.nwk]..."
	p := "Rename the nodes of phylogenies."
	e := "reno simpleNames.txt largeTree.nwk"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	flag.Parse()
	if *optV {
		util.PrintInfo("reno")
	}
	names := make(map[string]string)
	var fon *os.File
	files := flag.Args()
	if len(files) == 0 {
		log.Fatal("please provide a file of node names")
	}
	fon, err := os.Open(files[0])
	util.Check(err)
	sc := bufio.NewScanner(fon)
	for sc.Scan() {
		t := sc.Text()
		if t[0] != '#' {
			fields := strings.Fields(t)
			if len(fields) != 2 {
				log.Fatalf("%q seems mal-formed\n", t)
			}
			names[fields[0]] = fields[1]
		}
	}
	fon.Close()
	clio.ParseFiles(files[1:], parse, names)
}
