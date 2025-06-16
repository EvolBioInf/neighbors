package main

import (
	"flag"
	"github.com/evolbioinf/clio"
	"github.com/evolbioinf/neighbors/tdb"
	"github.com/evolbioinf/neighbors/util"
)

func main() {
	var optG = flag.String("g", "assembly_summary_genbank.txt",
		"genbank assemblies")
	var optR = flag.String("r", "assembly_summary_refseq.txt",
		"refseq assemblies")
	var optA = flag.String("a", "names.dmp", "taxonomic names")
	var optO = flag.String("o", "nodes.dmp", "node information")
	var optM = flag.String("m", "merged.dmp", "merged taxa")
	var optD = flag.String("d", "neidb", "database name")
	var optV = flag.Bool("v", false, "print version & "+
		"program information")
	u := "makeNeiDb [option]..."
	p := "Construct a taxonomy database for discovering " +
		"neighbor genomes.\n\tGenomes:  " +
		"<ftp>/genomes/ASSEMBLY_REPORTS/assembly_summary_" +
		"(genbank|refseq).txt" +
		"\n\tTaxonomy: <ftp>/pub/taxonomy/taxdump.tar.gz" +
		"\n\t<ftp>=ftp.ncbi.nlm.nih.gov"
	e := "makeNeiDb -d myNew.db"
	clio.Usage(u, p, e)
	flag.Parse()
	if *optV {
		util.PrintInfo("makeNeiDb")
	}
	tdb.NewTaxonomyDB(*optO, *optA, *optM, *optG, *optR, *optD)
}
