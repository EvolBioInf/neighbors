package tdb

import (
	"testing"
)

func TestTdb(t *testing.T) {
	p := "../data/"
	no := p + "nodesTest.dmp"
	na := p + "namesTest.dmp"
	me := p + "mergedTest.dmp"
	gb := p + "gbTest.txt"
	rs := p + "rsTest.txt"
	db := p + "taxSmall.db"
	NewTaxonomyDB(no, na, me, gb, rs, db)
	taxdb := OpenTaxonomyDB(db)
	subtree, _ := taxdb.Subtree(207598)
	if len(subtree) != 26 {
		t.Errorf("get %d rows, want 26", len(subtree))
		for _, s := range subtree {
			println(s)
		}
	}
	nl := len(AssemblyLevels())
	if nl != 4 {
		t.Errorf("get %d assembly levels, want 4", nl)
	}
	tid := 9606
	get, _ := taxdb.Name(tid)
	want := "Homo sapiens"
	if get != want {
		t.Errorf("get: %q; want: %q", get, want)
	}
	get, _ = taxdb.CommonName(tid)
	want = "human"
	if get != want {
		t.Errorf("get common_name: %s; want: %s\n",
			get, want)
	}
	get, _ = taxdb.Rank(tid)
	want = "species"
	if get != want {
		t.Errorf("get rank: %s; want: %s\n", get, want)
	}
	g, _ := taxdb.Parent(tid)
	w := 9605
	if g != w {
		t.Errorf("get parent: %d; want: %d", g, w)
	}
	children, _ := taxdb.Children(tid)
	g = len(children)
	w = 2
	if g != w {
		t.Errorf("get %d children; want %d", g, w)
	}
	tid = 207598
	taxa, _ := taxdb.Subtree(tid)
	g = len(taxa)
	w = 26
	if g != w {
		t.Errorf("get %d nodes in subtree; want %d", g, w)
	}
	targets := make([][]int, 0)
	var res []int
	targets = append(targets, []int{46359})
	res = append(res, 46359)
	targets = append(targets, []int{46359, 1159185})
	res = append(res, 499232)
	targets = append(targets, []int{46359, 406788})
	res = append(res, 9592)
	targets = append(targets, []int{37011, 9597})
	res = append(res, 9596)
	targets = append(targets, []int{37011, 9597, 46359})
	res = append(res, 207598)
	for i, target := range targets {
		get, _ := taxdb.MRCA(target)
		want := res[i]
		if get != want {
			t.Errorf("get: %d\nwant: %d\n", get, want)
		}
	}
	g, _ = taxdb.NumTaxa()
	w = 33
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
	taxa, _ = taxdb.Taxids("%homo sapiens%", 2, 2)
	g = len(taxa)
	w = 2
	if g != w {
		t.Errorf("get %d taxa for homo sapiens; want %d",
			g, w)
	}
	taxa, _ = taxdb.CommonTaxids("%man%", -1, 0)
	g = len(taxa)
	w = 2
	if g != w {
		t.Errorf("get %d taxa for man; want %d",
			g, w)
	}
	getBool, _ := taxdb.IsLeaf(9606)
	wantBool := false
	if getBool != wantBool {
		t.Errorf("get isLeaf %v; want %v",
			getBool, wantBool)
	}
	tid = 9606
	arr, _ := taxdb.Accessions(tid)
	g = len(arr)
	w = 1851
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
	acc := "GCA_049640585.1"
	get, _ = taxdb.Level(acc)
	want = "contig"
	if want != get {
		t.Errorf("get: %s\nwant: %s\n", get, want)
	}
	accessions := []string{
		"GCA_000002115.2",
		"GCA_000004845.2",
		"GCA_000181135.1"}
	levels := make(map[string]bool)
	levels["complete"] = true
	filteredAcc, _ := taxdb.FilterAccessions(accessions, levels)
	if len(filteredAcc) != 0 {
		t.Errorf("want 0 accessions, get %d\n", len(filteredAcc))
	}
	levels["chromosome"] = true
	filteredAcc, _ = taxdb.FilterAccessions(accessions, levels)
	if len(filteredAcc) != 1 {
		t.Errorf("want 1 accession, get %d\n", len(filteredAcc))
	}
	if accessions[0] != filteredAcc[0] {
		t.Errorf("want:\n%s\nget:\n%s\n",
			accessions[0],
			filteredAcc[0])
	}
	levels["scaffold"] = true
	filteredAcc, _ = taxdb.FilterAccessions(accessions, levels)
	if len(filteredAcc) != 2 {
		t.Errorf("want 2 accessions, get %d\n", len(filteredAcc))
	}
	for i, a := range filteredAcc {
		if accessions[i] != a {
			t.Errorf("want:\n%s\nget:\n%s\n",
				accessions[i],
				a)
		}
	}
	levels["contig"] = true
	filteredAcc, _ = taxdb.FilterAccessions(accessions, levels)
	for i, accession := range accessions {
		if filteredAcc[i] != accession {
			t.Errorf("want:\n%s\nget:\n%s\n",
				accession,
				filteredAcc[i])
		}
	}
	tid = 9606
	g, _ = taxdb.NumGenomes(tid, "complete")
	w = 58
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
	tid = 207598
	g, _ = taxdb.NumGenomesRec(207598, "complete")
	w = 58
	if g != 58 {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
	g, _ = taxdb.NumGenomesRec(tid, "complete")
	x, _ := taxdb.NumGenomesRec(tid, "chromosome")
	g += x
	x, _ = taxdb.NumGenomesRec(tid, "scaffold")
	g += x
	x, _ = taxdb.NumGenomesRec(tid, "contig")
	g += x
	w = 1888
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
}
