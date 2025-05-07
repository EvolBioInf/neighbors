tangle: db.sh query.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
test: db.sh query.sh
	make reset
	bash db.sh
	bash query.sh
reset:
	rm -rf acc.txt lpn.dist lpn.nwk ndata neidb targets tdata.zip all lpn.db lpn.fasta nacc.txt ndata.zip neighbors tacc.txt tdata
clean:
	make reset
	rm db.sh query.sh
