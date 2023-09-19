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
	rm -rf 256.db/ 256.fasta 264.db/ 264_257.db/ 264_257.fasta acc.txt all/ db.sh nacc.txt neidb ndata.zip ndata/ neighbors/ o157.dist o157.nwk query.sh tacc.txt targets/ tdata.zip tdata/
