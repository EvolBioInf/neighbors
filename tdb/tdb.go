// Package tdb constructs and queries the taxonomy database.
package tdb

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"github.com/evolbioinf/neighbors/util"
	_ "github.com/mattn/go-sqlite3"
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
	numChildren   int
	raw           map[string]int
	rec           map[string]int
}
type genome struct {
	taxid            int
	accession, level string
	size             float64
	written          bool
}

var assemblyLevels = []string{"complete",
	"chromosome",
	"scaffold",
	"contig"}

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

// The method Name takes as argument a taxon-ID and returns the  taxon's name and an error.
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

// The method Rank takes as argument a taxon-ID and returns the taxon's rank and an error.
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

// The method Parent takes as argument a taxon-ID and returns the  taxon-ID of its parent and an error.
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

// The method Children takes as argument a taxon-ID and returns its  children and an error.
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

// Taxids takes as arguments a taxon name, a limit on the number  of names returned, and an offset into the list of matching names. It  matches the taxon name, orders them by their score, imposes the  limit and offset, and returns the corresponding taxon-IDs and an  error.
func (t *TaxonomyDB) Taxids(name string,
	limit, offset int) ([]int, error) {
	var err error
	taxids := make([]int, 0)
	q := fmt.Sprintf(taxidsT, name, limit, offset)
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

// The method MRCA takes as input a slice of taxon-IDs and returns their most recent common ancestor and an error.
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

// The method NumGenomes takes as argument a taxon-ID and an assembly level and returns the raw number of genomes associated with this taxon assembled to that level and an error.  NB: This is not the number of genomes in the subtree rooted on that taxon, please use the method NumGenomesRecursive for that.
func (d *TaxonomyDB) NumGenomes(taxid int, level string) (int, error) {
	n := 0
	var err error
	q := "select raw from genome_count " +
		"where taxid=%d and " +
		"level like '%s'"
	q = fmt.Sprintf(q, taxid, level)
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

// The method NumGenomesRec takes as argument a taxon-ID and an assembly level. It returns the number of genomes assembled to that level contained in the subtree rooted on the taxon-ID and an error.
func (d *TaxonomyDB) NumGenomesRec(taxid int, level string) (int, error) {
	n := 0
	var err error
	q := "select recursive from genome_count " +
		"where taxid=%d and " +
		"level like '%s'"
	q = fmt.Sprintf(q, taxid, level)
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

// The function NewTaxonomyDB takes as parameters the names of the  five input files from which we construct the database, and the name of  the database. It opens these files, opens a new database, and  constructs the database.
func NewTaxonomyDB(nodes, names, merged,
	genbank, refseq, dbName string) {
	of := util.Open(nodes)
	defer of.Close()
	af := util.Open(names)
	defer af.Close()
	mf := util.Open(merged)
	defer mf.Close()
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
            taxid int primary key,
            parent int,
            name text,
            rank text,
            score float);
          create index taxon_parent_idx on taxon(parent);
          create index taxon_score_idx on taxon(score);`
	_, err = db.Exec(sqlStmt)
	util.Check(err)
	sqlStmt = `create table genome (
            taxid int,
            size real, 
            accession text primary key,
            level text,
            foreign key(taxid) references taxon(taxid));
          create index genome_taxid_idx on genome(taxid);
          create index genome_size_idx on genome(size);`
	_, err = db.Exec(sqlStmt)
	util.Check(err)
	sqlStmt = `create table genome_count (
            taxid int,
            level text,
            raw int,
            recursive int,
            primary key(taxid, level),
            foreign key(taxid) references taxon(taxid));
          create index genome_count_raw_idx on genome_count(raw);
          create index genome_count_recursive_idx on genome_count(recursive);`
	_, err = db.Exec(sqlStmt)
	util.Check(err)
	taxa := make(map[int]*taxon)
	scanner := bufio.NewScanner(of)
	for scanner.Scan() {
		row := scanner.Text()
		t := new(taxon)
		t.raw = make(map[string]int)
		t.rec = make(map[string]int)
		fields := strings.SplitN(row, "\t|\t", 4)
		t.taxid, err = strconv.Atoi(fields[0])
		util.Check(err)
		t.parent, err = strconv.Atoi(fields[1])
		util.Check(err)
		t.rank = fields[2]
		taxa[t.taxid] = t
	}
	scanner = bufio.NewScanner(af)
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Split(row, "\t|\t")
		id, err := strconv.Atoi(fields[0])
		util.Check(err)
		if fields[3][:3] == "sci" {
			taxa[id].name = fields[1]
		}
	}
	tx, err := db.Begin()
	util.Check(err)
	sqlStmt = "insert into taxon(taxid, parent, name, rank) " +
		"values(?, ?, ?, ?)"
	stmt, err := tx.Prepare(sqlStmt)
	util.Check(err)
	for _, v := range taxa {
		_, err = stmt.Exec(v.taxid, v.parent, v.name, v.rank)
		util.Check(err)
	}
	tx.Commit()
	stmt.Close()
	genomes := make(map[string]*genome)
	scanner = bufio.NewScanner(rf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		fields := strings.Split(row, "\t")
		k := coreAcc(fields[0])
		g := fields2genome(fields)
		genomes[k] = g
	}
	scanner = bufio.NewScanner(gf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		fields := strings.Split(row, "\t")
		k := coreAcc(fields[0])
		if genomes[k] == nil {
			g := fields2genome(fields)
			genomes[k] = g
		}
	}
	merge := make(map[int]int)
	scanner = bufio.NewScanner(mf)
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '#' {
			continue
		}
		f := strings.Split(row, "\t|\t")
		old, err := strconv.Atoi(f[0])
		util.Check(err)
		new, err := strconv.Atoi(f[1][:len(f[1])-2])
		util.Check(err)
		merge[old] = new
	}
	tx, err = db.Begin()
	util.Check(err)
	sqlStmt = "pragma foreign_keys=on"
	_, err = tx.Exec(sqlStmt)
	util.Check(err)
	sqlStmt = "insert into genome(accession, " +
		"taxid, level, size) " +
		"values(?, ?, ?, ?)"
	stmt, err = tx.Prepare(sqlStmt)
	util.Check(err)
	for _, genome := range genomes {
		x := merge[genome.taxid]
		if x != 0 {
			genome.taxid = x
		}
		_, err = stmt.Exec(genome.accession, genome.taxid,
			genome.level, genome.size)
		util.Check(err)
	}
	tx.Commit()
	stmt.Close()
	for _, genome := range genomes {
		t := taxa[genome.taxid]
		if t != nil {
			t.raw[genome.level]++
		} else {
			m := "WARNING[tdb]: no entry in taxonomy for " +
				"%d referred to by assembly %s; " +
				"could this be an unmerged taxon?\n"
			fmt.Fprintf(os.Stderr, m, genome.taxid,
				genome.accession)
		}
	}
	for _, taxon := range taxa {
		for level, count := range taxon.raw {
			taxon.rec[level] = count
		}
	}
	leaves := []*taxon{}
	for _, taxon := range taxa {
		parent := taxa[taxon.parent]
		if parent != nil && parent.taxid != taxon.taxid {
			parent.numChildren++
		}
	}
	for _, taxon := range taxa {
		if taxon.numChildren == 0 {
			leaves = append(leaves, taxon)
		}
	}
	for len(leaves) > 0 {
		i := 0
		for _, leaf := range leaves {
			var parent *taxon
			parent = taxa[leaf.parent]
			if parent == nil || leaf.taxid == parent.taxid {
				continue
			}
			for level, count := range leaf.rec {
				parent.rec[level] += count
			}
			parent.numChildren--
			if parent.numChildren == 0 {
				leaves[i] = parent
				i++
			}
		}
		leaves = leaves[:i]
	}
	tx, err = db.Begin()
	util.Check(err)
	sqlStmt = "pragma foreign_keys=on"
	_, err = tx.Exec(sqlStmt)
	util.Check(err)
	sqlStmt = `insert into genome_count(
          level, recursive, raw, taxid)
          values(?, ?, ?, ?)`
	stmt, err = tx.Prepare(sqlStmt)
	util.Check(err)
	for _, t := range taxa {
		for _, level := range assemblyLevels {
			recCount := t.rec[level]
			rawCount := t.raw[level]
			_, err = stmt.Exec(level,
				recCount, rawCount, t.taxid)
			util.Check(err)
		}
	}
	tx.Commit()
	stmt.Close()
	tx, err = db.Begin()
	util.Check(err)
	sqlStmt = `update taxon
            set score = ?
            where taxid = ?`
	stmt, err = tx.Prepare(sqlStmt)
	util.Check(err)
	for _, taxon := range taxa {
		sum := 0
		for _, count := range taxon.rec {
			sum += count
		}
		_, err := stmt.Exec(sum, taxon.taxid)
		util.Check(err)
	}
	tx.Commit()
	stmt.Close()
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
	util.Check(err)
	_, err = db.db.Exec("PRAGMA foreign_keys = ON;")
	util.Check(err)
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

// The function AssemblyLevels returns the slice of possible assembly levels.
func AssemblyLevels() []string {
	return assemblyLevels
}

var accessionT = "select accession " +
	"from genome " +
	"where taxid=%d"
var nameT = "select name from taxon where taxid=%d"
var rankT = "select rank from taxon where taxid=%d"
var parentT = "select parent from taxon where taxid=%d"
var childrenT = "select taxid from taxon where parent=%d"
var taxidsT = "select taxid from taxon " +
	"where name like '%s' " +
	"order by score desc " +
	"limit %d " +
	"offset %d"
var levelT = "select level from genome where accession='%s'"
