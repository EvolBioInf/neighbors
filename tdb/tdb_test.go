package tdb

import (
	"testing"
)

func TestTdb(t *testing.T) {
	p := "../data/"
	no := p + "nodesTest.dmp"
	na := p + "namesTest.dmp"
	gb := p + "gbTest.txt"
	rs := p + "rsTest.txt"
	db := p + "taxSmall.db"
	NewTaxonomyDB(no, na, gb, rs, db)
	taxdb := OpenTaxonomyDB(db)
	subtree := taxdb.Subtree(207598)
	if len(subtree) != 26 {
		t.Errorf("get %d rows, want 26", len(subtree))
		for _, s := range subtree {
			println(s)
		}
	}
	tid := 9606
	want := "Homo sapiens"
	get := taxdb.Name(tid)
	if get != want {
		t.Errorf("get: %q; want: %q", get, want)
	}
	w := 9605
	g := taxdb.Parent(tid)
	if g != w {
		t.Errorf("get parent: %d; want: %d", g, w)
	}
	w = 2
	g = len(taxdb.Children(tid))
	if g != w {
		t.Errorf("get %d children; want %d", g, w)
	}
	tid = 207598
	w = 26
	taxa := taxdb.Subtree(tid)
	g = len(taxa)
	if g != w {
		t.Errorf("get %d nodes in subtree; want %d", g, w)
	}
	w = 4
	taxa = taxdb.Taxids("%homo sapiens%")
	g = len(taxa)
	if g != w {
		t.Errorf("get %d taxa for homo sapiens; want %d",
			g, w)
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
		get := taxdb.MRCA(target)
		want := res[i]
		if get != want {
			t.Errorf("get: %d\nwant: %d\n", get, want)
		}
	}
	want = "contig"
	acc := "GCA_049640585.1"
	get = taxdb.Level(acc)
	if want != get {
		t.Errorf("get: %s\nwant: %s\n", get, want)
	}
	accessions := []string{
		"GCA_000002115.2",
		"GCA_000004845.2",
		"GCA_000181135.1"}
	levels := make(map[string]bool)
	levels["complete"] = true
	filteredAcc := taxdb.FilterAccessions(accessions, levels)
	if len(filteredAcc) != 0 {
		t.Errorf("want 0 accessions, get %d\n", len(filteredAcc))
	}
	levels["chromosome"] = true
	filteredAcc = taxdb.FilterAccessions(accessions, levels)
	if len(filteredAcc) != 1 {
		t.Errorf("want 1 accession, get %d\n", len(filteredAcc))
	}
	if accessions[0] != filteredAcc[0] {
		t.Errorf("want:\n%s\nget:\n%s\n",
			accessions[0],
			filteredAcc[0])
	}
	levels["scaffold"] = true
	filteredAcc = taxdb.FilterAccessions(accessions, levels)
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
	filteredAcc = taxdb.FilterAccessions(accessions, levels)
	for i, accession := range accessions {
		if filteredAcc[i] != accession {
			t.Errorf("want:\n%s\nget:\n%s\n",
				accession,
				filteredAcc[i])
		}
	}
	w = 33
	g = taxdb.NumTaxa()
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}
	w = 1919
	g = taxdb.NumGenomes()
	if g != w {
		t.Errorf("get:\n%d\nwant:\n%d\n", g, w)
	}

}
