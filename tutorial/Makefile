tangle: db.sh query.sh small.sh large.sh
db.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rdb.sh > db.sh
query.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rquery.sh > query.sh
small.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rsmall.sh > small.sh
large.sh: tutorial.org
	bash ../scripts/org2nw tutorial.org | notangle -Rlarge.sh > large.sh
test: db.sh query.sh small.sh large.sh
	make reset
	bash db.sh
	bash query.sh
	bash small.sh
	bash large.sh
reset:
	rm -rf large neidb small
clean:
	make reset
	rm db.sh query.sh small.sh large.sh
