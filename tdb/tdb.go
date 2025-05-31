// Package tdb constructs and queries the taxonomy database.
package tdb

import (
	"bufio"
	"database/sql"
	"errors"
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
	taxid            int
	accession, level string
	size             float64
	written          bool
}

// The method Close closes a taxonomy database.
func (t *TaxonomyDB) Close() {
	t.db.Close()
}

// The method Accessions takes as parameter a taxon-ID and returns a slice accessions of genome assemblies belonging to that taxon and an error.
func (t *TaxonomyDB) Accessions(taxon int) ([]string, error) {
	var err error
	accessions := make([]string, 0)
	q := fmt.Sprintf(accessionT, taxon)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	accession := ""
	for rows.Next() {
		err := rows.Scan(&accession)
		if err != nil {
			return nil, err
		}
		accessions = append(accessions, accession)
	}
	return accessions, err
}

// The method Name takes as argument a taxon ID and returns the  taxon's name and an error.
func (t *TaxonomyDB) Name(taxon int) (string, error) {
	var err error
	name := ""
	q := fmt.Sprintf(nameT, taxon)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, err
}

// The method Rank takes as argument a taxon ID and returns the taxon's rank and an error.
func (t *TaxonomyDB) Rank(taxon int) (string, error) {
	var err error
	rank := ""
	q := fmt.Sprintf(rankT, taxon)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	rows.Next()
	err = rows.Scan(&rank)
	if err != nil {
		return "", err
	}
	return rank, err
}

// The method Parent takes as argument a taxon ID and returns the  taxon ID of its parent and an error.
func (t *TaxonomyDB) Parent(c int) (int, error) {
	var err error
	parent := 0
	q := fmt.Sprintf(parentT, c)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&parent)
	if err != nil {
		return 0, err
	}
	return parent, err
}

// The method Children takes as argument a taxon ID and returns its  children and an error.
func (t *TaxonomyDB) Children(p int) ([]int, error) {
	var err error
	children := make([]int, 0)
	q := fmt.Sprintf(childrenT, p)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	child := 0
	for rows.Next() {
		err = rows.Scan(&child)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	return children, err
}

// The method Subtree returns all taxa in a subtree, including its  root, and an error.
func (t *TaxonomyDB) Subtree(r int) ([]int, error) {
	var err error
	taxa := make([]int, 0)
	taxa, err = traverseSubtree(t, r, taxa)
	if err != nil {
		return nil, err
	}
	return taxa, err
}

// Taxids matches the name of a taxon and returns the  corresponding taxon-IDs and an error.
func (t *TaxonomyDB) Taxids(name string) ([]int, error) {
	var err error
	taxids := make([]int, 0)
	q := fmt.Sprintf(taxidsT, name)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	taxid := 0
	for rows.Next() {
		err = rows.Scan(&taxid)
		if err != nil {
			return nil, err
		}
		taxids = append(taxids, taxid)
	}
	return taxids, err
}

// The method MRCA takes as input a slice of taxon IDs and returns their most recent common ancestor and an error.
func (t *TaxonomyDB) MRCA(ids []int) (int, error) {
	var err error
	mrca := -1
	if len(ids) == 0 {
		m := "Empty ID list in tdb.MRCA"
		return 0, errors.New(m)
	} else if len(ids) == 1 {
		return ids[0], nil
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
			parent, err := t.Parent(child)
			if err != nil {
				return 0, err
			}
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
	return mrca, err
}

// The method Level takes as argument a genome accession and  returns the assembly level and an eror.
func (t *TaxonomyDB) Level(acc string) (string, error) {
	var err error
	level := ""
	q := fmt.Sprintf(levelT, acc)
	rows, err := t.db.Query(q)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	rows.Next()
	err = rows.Scan(&level)
	if err != nil {
		return "", err
	}
	return level, err
}

// FilterAccessions takes as input a slice of genome accessions  and a list of desired assembly levels. It then removes any accession  that doesn't conform to one of the levels supplied and returns the  adjusted slice of genome accessions and an error. The input  accessions remain unchanged.
func (d *TaxonomyDB) FilterAccessions(acc []string,
	levels map[string]bool) ([]string, error) {
	newAcc := []string{}
	var err error
	for _, a := range acc {
		level, err := d.Level(a)
		if err != nil {
			return nil, err
		}
		if levels[level] {
			newAcc = append(newAcc, a)
		}
	}
	return newAcc, err
}

// The method NumTaxa returns the number of taxa in the database  and an error.
func (d *TaxonomyDB) NumTaxa() (int, error) {
	n := 0
	var err error
	q := "select count(*) from taxon"
	row, err := d.db.Query(q)
	defer row.Close()
	if err != nil {
		return 0, err
	}
	row.Next()
	err = row.Scan(&n)
	if err != nil {
		return 0, err
	}
	return n, err
}

// The method NumGenomes returns the number of genomes in the  database and an error.
func (d *TaxonomyDB) NumGenomes() (int, error) {
	n := 0
	var err error
	q := "select count(*) from genome"
	row, err := d.db.Query(q)
	defer row.Close()
	if err != nil {
		return 0, err
	}
	row.Next()
	err = row.Scan(&n)
	if err != nil {
		return 0, err
	}
	return n, err
}

// The function NewTaxonomyDB takes as parameters the names of the  four input files from which we construct the database, and the name of  the database. It opens these files, opens a new database, and  constructs the database.
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
	util.Check(err)
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
                   accession text, level text,
          primary key(accession),
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
		"taxid, level, size) " +
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
			g.level, g.size)
		util.Check(err)
	}
	for _, g := range rsGenomes {
		if g.written {
			continue
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.level, g.size)
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
			g.level, g.size)
		util.Check(err)
	}
	for _, g := range rsGenomes {
		if g.written {
			continue
		}
		_, err = stmt.Exec(g.accession, g.taxid,
			g.level, g.size)
	}
}
func coreAcc(acc string) string {
	s := strings.Index(acc, "_") + 1
	e := strings.Index(acc, ".")
	core := acc[s:e]
	return core
}
func fields2genome(fields []string) *genome {
	g := new(genome)
	g.accession = fields[0]
	id, err := strconv.Atoi(fields[5])
	util.Check(err)
	g.taxid = id
	g.level = fields[11]
	g.level = strings.Fields(g.level)[0]
	g.level = strings.ToLower(g.level)
	si, err := strconv.Atoi(fields[25])
	util.Check(err)
	g.size = float64(si) / 1000000.0
	return g
}

// The function OpenTaxonomyDB opens an existing taxonomy database  and returns a pointer to it.
func OpenTaxonomyDB(name string) *TaxonomyDB {
	db := new(TaxonomyDB)
	var err error
	db.db, err = sql.Open("sqlite3", name)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func traverseSubtree(t *TaxonomyDB, v int, taxa []int) ([]int, error) {
	taxa = append(taxa, v)
	children, err := t.Children(v)
	if err != nil {
		return nil, err
	}
	for _, child := range children {
		if child != v {
			taxa, err = traverseSubtree(t, child, taxa)
			if err != nil {
				return nil, err
			}
		}
	}
	return taxa, err
}

var accessionT = "select accession " +
	"from genome " +
	"where taxid=%d"
var nameT = "select name from taxon where taxid=%d"
var rankT = "select rank from taxon where taxid=%d"
var parentT = "select parent from taxon where taxid=%d"
var childrenT = "select taxid from taxon where parent=%d"
var taxidsT = "select taxid from taxon where name like '%s'"
var levelT = "select level from genome where accession='%s'"
