package main

import (
	"flag"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
)

func main() {
	var optP = flag.String("p", "prokaryotes.txt", "prokaryote genomes")
	var optE = flag.String("e", "eukaryotes.txt", "eukaryote genomes")
	var optI = flag.String("i", "viruses.txt", "virus genomes")
	var optA = flag.String("a", "names.dmp", "taxonomic names")
	var optO = flag.String("o", "nodes.dmp", "node information")
	var optD = flag.String("d", "neidb", "database name")
	var optV = flag.Bool("v", false, "print version & "+
		"program information")
	u := "makeNeiDb [option]..."
	p := "Construct a taxonomy database for discovering neighbor genomes." +
		"\n\tGenomes:  <ftp>/genomes/GENOME_REPORTS/" +
		"((pro|eu)karyotes|viruses).txt" +
		"\n\tTaxonomy: <ftp>/pub/taxonomy/taxdump.tar.gz" +
		"\n\t<ftp>=ftp.ncbi.nlm.nih.gov"
	e := "makeNeiDb -d myNew.db"
	clio.Usage(u, p, e)
	flag.Parse()
	if *optV {
		util.PrintInfo("makeNeiDb")
	}
	tdb.NewTaxonomyDB(*optO, *optA, *optP, *optE, *optI, *optD)
}
