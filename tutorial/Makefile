all: db.sh query.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
test: db.sh query.sh
	bash db.sh
	bash query.sh
clean:
	rm -rf acc.txt au2.db au2.fasta au2.nwk au3.nwk au.nwk db.sh frag*.fasta iv.txt macs*.out nei.* neidb neidb.tgz neighbors query.sh tar.* targets tmp.svg
