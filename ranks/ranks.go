package ranks

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
)

var Id int = 0

type Node struct {
	Id       int
	N        int
	Percent  float64
	Label    string
	Parent   *Node
	Children map[string]*Node
}

func newNode(label string, count, total int,
	parent *Node) *Node {
	v := new(Node)
	v.Label = label
	v.N = count
	v.Percent = float64(count) / float64(total) * 100.0
	v.Parent = parent
	v.Children = make(map[string]*Node)
	v.Id = Id
	Id++
	return v
}
func printNodes(v *Node) {
	fmt.Printf("\t%d [label=\"%s\\n%d (",
		v.Id, v.Label, v.N)
	if v.Percent >= 0.1 {
		fmt.Printf("%.1f%%", v.Percent)
	} else {
		fmt.Printf("< 0.1%%")
	}
	fmt.Printf(")\"]\n")
	keys := []string{}
	for key, _ := range v.Children {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, key := range keys {
		printNodes(v.Children[key])
	}
}
func printEdges(v *Node) {
	if v.Parent != nil {
		fmt.Printf("\t%d -> %d\n", v.Parent.Id, v.Id)
	}
	keys := []string{}
	for key, _ := range v.Children {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	for _, key := range keys {
		printEdges(v.Children[key])
	}
}
func Run() {
	genomes := []string{}
	util.SetName("ranks")
	u := "ranks [option]... <taxon-ID> <db>"
	p := "Calculate distribution of genomes " +
		"among taxonomic ranks."
	e := "ranks -g myGenomeList.txt 562 neidb"
	clio.Usage(u, p, e)
	optV := flag.Bool("v", false, "version")
	optL := flag.Bool("l", false, "list genomes (implies -t)")
	optLL := flag.String("L", "", util.LevelMsg())
	optT := flag.Bool("t", false, "tabular output (default tree)")
	optG := flag.String("g", "", "read genome accessions "+
		"from file")
	optR := flag.Bool("r", false, "remote execution (implies db)")
	optDD := flag.String("D", "", "name of remote database"+
		"(implies remote execution)")
	flag.Parse()
	if *optV {
		util.PrintInfo("ranks")
	}
	if *optR || *optDD != "" {
		var resp string
		options := []util.Option{
			{Name: "r", WithValue: false},
			{Name: "D", WithValue: true}}
		misc := map[string]string{}
		if *optDD != "" {
			misc["db"] = *optDD
		}
		if *optG != "" {
			resp = util.SendPostRequest(
				"api/v2/programs/ranks",
				util.SanitizeArguments(
					os.Args[1:],
					options),
				[]string{},
				misc,
				[]*os.File{util.Open(*optG)},
				nil,
			)
		} else {
			resp = util.SendGetRequest(
				"api/v2/programs/ranks",
				util.SanitizeArguments(
					os.Args[1:],
					options),
				[]string{},
				misc,
			)
		}
		fmt.Print(resp)
		return
	}
	args := flag.Args()
	if len(args) < 2 {
		m := "please enter the root taxon-ID " +
			"and the Neighbors database"
		log.Fatal(m)
	}
	root, err := strconv.Atoi(args[0])
	util.Check(err)
	neidb, err := tdb.OpenTaxonomyDBcheck(args[1])
	util.Check(err)
	levels := make(map[string]bool)
	knowns := tdb.AssemblyLevels()
	if *optLL != "" {
		requests := strings.Split(*optLL, ",")
		slices.Sort(knowns)
		for _, request := range requests {
			_, found := slices.BinarySearch(knowns, request)
			if !found {
				log.Fatalf("unknown level %q", request)
			}
			levels[request] = true
		}
	} else {
		for _, known := range knowns {
			levels[known] = true
		}
	}
	if *optG != "" {
		f, err := os.Open(*optG)
		util.Check(err)
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			genomes = append(genomes, sc.Text())
		}
		f.Close()
	} else {
		taxa, err := neidb.Subtree(root)
		util.Check(err)
		for _, taxon := range taxa {
			accessions, err := neidb.Accessions(taxon)
			util.Check(err)
			for _, accession := range accessions {
				genomes = append(genomes, accession)
			}
		}
	}
	genomes, err = neidb.FilterAccessions(genomes, levels)
	util.Check(err)
	tree := make(map[string][]string)
	ranks := []string{}
	for _, genome := range genomes {
		path := ""
		ranks = ranks[:0]
		currTaxid, err := neidb.AccessionTaxid(genome)
		if err != nil {
			util.Check(err)
		}
		rank, err := neidb.Rank(currTaxid)
		util.Check(err)
		ranks = append(ranks, rank)
		for currTaxid != root {
			currTaxid, err = neidb.Parent(currTaxid)
			util.Check(err)
			rank, err := neidb.Rank(currTaxid)
			util.Check(err)
			ranks = append(ranks, rank)
		}
		slices.Reverse(ranks)
		path = strings.Join(ranks, ":")
		tree[path] = append(tree[path], genome)
	}
	paths := []string{}
	for path, _ := range tree {
		paths = append(paths, path)
	}
	slices.Sort(paths)
	if *optT || *optL {
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
		fmt.Fprintf(w, "#%s\t%s\t%s",
			"Path",
			"Count",
			"%")
		if *optL {
			fmt.Fprintf(w, "\tGenomes")
		}
		fmt.Fprintf(w, "\n")
		n := float64(len(genomes))
		for _, path := range paths {
			c := float64(len(tree[path]))
			p := c / n * 100
			ps := strings.Replace(path, " ", "_", -1)
			fmt.Fprintf(w, "%s\t%.0f\t%.4f", ps, c, p)
			if *optL {
				slices.Sort(tree[path])
				gs := strings.Join(tree[path], " ")
				fmt.Fprintf(w, "\t%s", gs)
			}
			fmt.Fprintf(w, "\n")
		}
		w.Flush()
	} else {
		start := newNode("start", 0, 1, nil)
		for _, path := range paths {
			v := start
			n := len(tree[path])
			pathElements := strings.Split(path, ":")
			for _, pathElement := range pathElements {
				parent := v
				child := v.Children[pathElement]
				if child == nil {
					child = newNode(pathElement, 0, 1, parent)
					v.Children[pathElement] = child
				}
				v = child
			}
			v.N = n
			v.Percent = float64(n) / float64(len(genomes)) * 100.0
		}
		children := []*Node{}
		for _, k := range start.Children {
			children = append(children, k)
		}
		if len(children) != 1 {
			log.Fatalf("start has %d children",
				len(children))
		}
		root := children[0]
		root.Parent = nil
		fmt.Println("digraph g {")
		fmt.Println("\trankdir=LR")
		printNodes(root)
		printEdges(root)
		fmt.Println("}")
	}
}
