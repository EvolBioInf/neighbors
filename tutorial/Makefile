all: db.sh query.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
test: db.sh query.sh
	bash db.sh
	bash query.sh
clean:
	rm -rf acc.txt au2.db au2.fasta au2.nwk au3.db au3.fasta au3.nwk au.db au.fasta au.nwk db.sh neidb neidb.tgz neighbors query.sh targets tmp.svg
