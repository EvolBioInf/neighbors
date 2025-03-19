package tdb

import (
	"testing"
)

func TestTdb(t *testing.T) {
	p := "../data/"
	no := p + "nodesTest.dmp"
	na := p + "namesTest.dmp"
	pr := p + "prokaryotes.txt"
	eu := p + "eukaryotes.txt"
	vi := p + "viruses.txt"
	d1 := p + "taxSmall.db"
	NewTaxonomyDB(no, na, pr, eu, vi, d1)
	taxdb := OpenTaxonomyDB(d1)
	subtree := taxdb.Subtree(1)
	if len(subtree) != 11 {
		t.Errorf("get %d rows, want 11", len(subtree))
		for _, s := range subtree {
			println(s)
		}
	}
	d2 := p + "neidb"
	taxdb = OpenTaxonomyDB(d2)
	tid := 866775
	reps := taxdb.Replicons(tid)
	get := reps[0]
	want := "chromosome:NC_015278.1/CP002512.1"
	if get != want {
		t.Errorf("get: %q; want: %q", get, want)
	}
	tid = 9606
	want = "Homo sapiens"
	get = taxdb.Name(tid)
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
	w = 11
	taxa = taxdb.Taxids("%homo sapiens%")
	g = len(taxa)
	if g != w {
		t.Errorf("get %d taxa for homo sapiens; want %d", g, w)
	}
}
