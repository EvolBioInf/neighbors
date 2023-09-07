all: db.sh query.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
test: db.sh query.sh
	rm -rf targets neighbors neidb.tgz neidb au.db ndata.zip ndata tdata tdata.zip
	bash db.sh
	bash query.sh
clean:
	rm -rf all.txt au.db au.dist db.sh markers*.fasta nacc.txt ndata.zip ndata neidb neighbors query.sh tacc.txt tar.txt targets tdata.zip tdata tmp.svg
