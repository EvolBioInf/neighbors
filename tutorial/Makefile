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
	rm -rf large neidb small
clean:
	make reset
	rm db.sh query.sh
