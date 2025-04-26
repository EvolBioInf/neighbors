// Package tdb constructs and queries the taxonomy database.
package tdb

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/evolbioinf/neighbors/util"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
)

type TaxonomyDB struct {
	db *sql.DB
}
type taxon struct {
	taxid, parent int
	name, rank    string
}
type genome struct {
	taxid             int
	accession, status string
	size              float64
	written           bool
}

// Close closes the taxonomy database.
func (t *TaxonomyDB) Close() {
	t.db.Close()
}

// The method Accessions takes as parameter a taxon-ID and returns a slice of assembly accessions.
func (t *TaxonomyDB) Accessions(tid int) []string {
	var accessions []string
	tmpl := "select accession from genome where taxid=%d"
	q := fmt.Sprintf(tmpl, tid)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	accession := ""
	for rows.Next() {
		err := rows.Scan(&accession)
		if err != nil {
			log.Fatal(err)
		}
		accessions = append(accessions, accession)
	}
	return accessions
}

// Name returns a taxon's name.
func (t *TaxonomyDB) Name(taxon int) string {
	n := ""
	tmpl := "select name from taxon where taxid=%d"
	q := fmt.Sprintf(tmpl, taxon)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// The method Rank takes as argument a taxon ID and returns the taxon's name. We construct the query, execute it, and extract the name.
func (t *TaxonomyDB) Rank(taxon int) string {
	rank := ""
	tmpl := "select rank from taxon where taxid=%d"
	q := fmt.Sprintf(tmpl, taxon)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&rank)
	if err != nil {
		log.Fatal(err)
	}
	return rank
}

// Parent returns a taxon's parent.
func (t *TaxonomyDB) Parent(c int) int {
	p := 0
	tmpl := "select parent from taxon where taxid=%d"
	q := fmt.Sprintf(tmpl, c)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// Children returns a taxon's children.
func (t *TaxonomyDB) Children(p int) []int {
	c := make([]int, 0)
	tmpl := "select taxid from taxon where parent=%d"
	q := fmt.Sprintf(tmpl, p)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	x := 0
	for rows.Next() {
		err = rows.Scan(&x)
		if err != nil {
			log.Fatal(err)
		}
		c = append(c, x)
	}
	return c
}

// Subtree returns the taxa in the subtree rooted on the given taxon.
func (t *TaxonomyDB) Subtree(r int) []int {
	taxa := make([]int, 0)
	taxa = traverseSubtree(t, r, taxa)
	return taxa
}

// Taxids matches the name of a taxon and returns the corresponding
// taxon-IDs.
func (t *TaxonomyDB) Taxids(name string) []int {
	taxids := make([]int, 0)
	q := "select taxid from taxon where name like '%s'"
	q = fmt.Sprintf(q, name)
	rows, err := t.db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	taxid := 0
	for rows.Next() {
		err = rows.Scan(&taxid)
		if err != nil {
			log.Fatal(err)
		}
		taxids = append(taxids, taxid)
	}
	return taxids
}

// The method MRCA takes as input a slice of taxon IDs and returns their most recent common ancestor.
func (t *TaxonomyDB) MRCA(ids []int) int {
	mrca := -1
	if len(ids) == 0 {
		log.Fatal("Empty ID list in tdb.MRCA")
	} else if len(ids) == 1 {
		return ids[0]
	}
	desc := make(map[int]int)
	for _, id := range ids {
		desc[id] = 1
	}
	parents := make([]int, 0)
	children := make([]int, 0)
	for _, id := range ids {
		children = append(children, id)
	}
	for len(children) > 1 {
		for _, child := range children {
			parent := t.Parent(child)
			desc[parent] += desc[child]
			if desc[parent] >= len(ids) {
				mrca = parent
				break
			}
			parents = append(parents, parent)
		}
		if mrca == -1 {
			children = children[:0]
			for _, parent := range parents {
				children = append(children, parent)
			}
			parents = parents[:0]
		} else {
			break
		}
	}
	return mrca
}

// NewTaxonomyDB takes as parameters the names
// of the four data files and the database name,
// and constructs the database from them.
func NewTaxonomyDB(nodes, names, genbank,
	refseq, dbName string) {
	of := util.Open(nodes)
	defer of.Close()
	af := util.Open(names)
	defer af.Close()
	gf := util.Open(genbank)
	defer gf.Close()
	rf := util.Open(refseq)
	defer rf.Close()
	_, err := os.Stat(dbName)
	if err == nil {
		fmt.Fprintf(os.Stderr, "database %s already exists\n",
			dbName)
		os.Exit(1)
	}
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `create table taxon (
          taxid int, parent int, name text, rank text,
          primary key(taxid));
          create index taxon_parent_idx on taxon(parent);`
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatal(err)
	}
	sqlStmt = `create table genome (
          taxid int, size real, 
                   accession text, status text,
          foreign key(taxid) references taxon(taxid));
          create index genome_taxid_idx on genome(taxid);
          create index genome_size_idx on genome(size);`
	if _, err := db.Exec(sqlStmt); err != nil {
		log.Fatal(err)
	}
	taxa := make(map[int]*taxon)
	scanner := bufio.NewScanner(of)
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.SplitN(row, "\t|\t", 4)
		t := new(taxon)
		t.taxid, err = strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		t.parent, err = strconv.Atoi(fields[1])
		if err != nil {
			log.Fatal(err)
		}
		t.rank = fields[2]
		taxa[t.taxid] = t
	}
	scanner = bufio.NewScanner(af)
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Split(row, "\t|\t")
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatal(err)
		}
		if fields[3][:3] == "sci" {
			taxa[id].name = fields[1]
		}
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt = "insert into taxon(taxid, parent, name, rank) " +
		"values(?, ?, ?, ?)"
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range taxa {
		_, err = stmt.Exec(v.taxid, v.parent, v.name, v.rank)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
	stmt.Close()
	rsGenomes := make(map[string]*genome)
	scanner = bufio.NewScanner(rf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		fields := strings.Split(row, "\t")
		k := coreAcc(fields[0])
		g := fields2genome(fields)
		rsGenomes[k] = g
	}
	tx, err = db.Begin()
	util.Check(err)
	sqlStmt = "insert into genome(accession, " +
		"taxid, status, size) " +
		"values(?, ?, ?, ?)"
	stmt, err = tx.Prepare(sqlStmt)
	util.Check(err)
	defer tx.Commit()
	defer stmt.Close()
	scanner = bufio.NewScanner(gf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		fields := strings.Split(row, "\t")
		ca := coreAcc(fields[0])
		var g *genome
		if _, ok := rsGenomes[ca]; ok {
			g = rsGenomes[ca]
			rsGenomes[ca].written = true
		} else {
			g = fields2genome(fields)
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.status, g.size)
		util.Check(err)
	}
	for _, g := range rsGenomes {
		if g.written {
			continue
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.status, g.size)
	}
	scanner = bufio.NewScanner(gf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		fields := strings.Split(row, "\t")
		ca := coreAcc(fields[0])
		var g *genome
		if _, ok := rsGenomes[ca]; ok {
			g = rsGenomes[ca]
			rsGenomes[ca].written = true
		} else {
			g = fields2genome(fields)
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.status, g.size)
		util.Check(err)
	}
	for _, g := range rsGenomes {
		if g.written {
			continue
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.status, g.size)
	}
}
func coreAcc(acc string) string {
	s := strings.Index(acc, "_") + 1
	e := strings.Index(acc, ".")
	core := acc[s:e]
	return core
}
func fields2genome(fields []string) *genome {
	acc := fields[0]
	id, err := strconv.Atoi(fields[5])
	util.Check(err)
	stat := fields[11]
	si, err := strconv.Atoi(fields[25])
	util.Check(err)
	sf := float64(si) / 1000000.0
	g := genome{accession: acc, taxid: id,
		status: stat, size: sf}
	return &g
}

// OpenTaxonomyDB opens an existing taxonomy database and returns a
// pointer to it.
func OpenTaxonomyDB(name string) *TaxonomyDB {
	db := new(TaxonomyDB)
	var err error
	db.db, err = sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func traverseSubtree(t *TaxonomyDB, r int, taxa []int) []int {
	taxa = append(taxa, r)
	ch := t.Children(r)
	for _, c := range ch {
		if c != r {
			taxa = traverseSubtree(t, c, taxa)
		}
	}
	return taxa
}
