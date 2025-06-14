SRC1 =   makeNeiDb.tex taxi.tex ants.tex climt.tex neighbors.tex dree.tex fintac.tex land.tex outliers.tex pickle.tex tdb.tex tutorial.tex util.tex
SRC2 =   header.tex introduction.tex neighborsDoc.tex labels.tex

date = $(shell git log | grep -m 1 Date | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

all: neighborsDoc.pdf
neighborsDoc.pdf: neighborsDoc.tex $(SRC1) $(SRC2) db.ps
	echo $(date) | tr '_' ' ' > date.txt
	echo $(version) | tr '-' ' ' | awk '{printf "%s", $$1; if ($$2) printf "-%s", $$2; printf "\n"}' > version.txt
	latex neighborsDoc.tex
	bibtex neighborsDoc
	latex neighborsDoc.tex
	latex neighborsDoc.tex
	dvips neighborsDoc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY neighborsDoc.ps
publish: neighborsDoc.pdf
	if [ -d /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/ ]; then \
		cp neighborsDoc.pdf /mnt/c/Users/haubold/ownCloud\ -\ haubold@evolbio.mpg.de@owncloud.gwdg.de/docs/; \
	fi

db.ps: db.dot
	neato -T eps db.dot > db.ps
makeNeiDb.tex: ../makeNeiDb/makeNeiDb.org 
	awk -f ../scripts/preWeave.awk ../makeNeiDb/makeNeiDb.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > makeNeiDb.tex
taxi.tex: ../taxi/taxi.org 
	awk -f ../scripts/preWeave.awk ../taxi/taxi.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > taxi.tex
ants.tex: ../ants/ants.org
	awk -f ../scripts/preWeave.awk ../ants/ants.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > ants.tex
climt.tex: ../climt/climt.org
	awk -f ../scripts/preWeave.awk ../climt/climt.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > climt.tex
neighbors.tex: ../neighbors/neighbors.org
	awk -f ../scripts/preWeave.awk ../neighbors/neighbors.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > neighbors.tex
dree.tex: ../dree/dree.org
	awk -f ../scripts/preWeave.awk ../dree/dree.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > dree.tex
fintac.tex: ../fintac/fintac.org
	awk -f ../scripts/preWeave.awk ../fintac/fintac.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > fintac.tex
land.tex: ../land/land.org
	awk -f ../scripts/preWeave.awk ../land/land.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > land.tex
pickle.tex: ../pickle/pickle.org
	awk -f ../scripts/preWeave.awk ../pickle/pickle.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > pickle.tex
outliers.tex: ../outliers/outliers.org
	awk -f ../scripts/preWeave.awk ../outliers/outliers.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > outliers.tex
tdb.tex: ../tdb/tdb.org
	awk -f ../scripts/preWeave.awk ../tdb/tdb.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > tdb.tex	
tutorial.tex: ../tutorial/tutorial.org
	awk -f ../scripts/preWeave.awk ../tutorial/tutorial.org | bash ../scripts/org2nw | noweave -n -x > tutorial.tex
util.tex: ../util/util.org
	awk -f ../scripts/preWeave.awk ../util/util.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > util.tex

clean:
	rm -f neighborsDoc.pdf neighborsDoc.ps neighborsDoc.dvi neighborsDoc.aux $(SRC1) db.ps
