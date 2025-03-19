tangle: db.sh query.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
test: db.sh query.sh
	rm -rf targets neighbors neidb.tgz neidb au.db ndata.zip ndata tdata tdata.zip
	bash db.sh
	bash query.sh
clean:
	rm -rf 292.db/ 292.fasta 293.db/ 293.fasta 300_293.db/ 300_293.fasta 300.db 300.fasta 301_294.db 301_294.fasta 301.db acc.txt all/ db.sh nacc.txt neidb ndata.zip ndata/ neighbors/ o157.dist o157.nwk query.sh tacc.txt targets/ tdata.zip tdata/ tlen.dat
