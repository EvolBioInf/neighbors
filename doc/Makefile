SRC1 =   makeNeiDb.tex taxi.tex ants.tex neighbors.tex dree.tex land.tex pickle.tex tax.tex tdb.tex tutorial.tex util.tex
SRC2 =   header.tex introduction.tex neighborsDoc.tex

all: neighborsDoc.pdf
neighborsDoc.pdf: neighborsDoc.tex $(SRC1) $(SRC2) 
	latex neighborsDoc.tex
	bibtex neighborsDoc
	latex neighborsDoc.tex
	latex neighborsDoc.tex
	dvips neighborsDoc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY neighborsDoc.ps

makeNeiDb.tex: ../makeNeiDb/makeNeiDb.org 
	awk -f ../scripts/preWeave.awk ../makeNeiDb/makeNeiDb.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > makeNeiDb.tex
taxi.tex: ../taxi/taxi.org 
	awk -f ../scripts/preWeave.awk ../taxi/taxi.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > taxi.tex
ants.tex: ../ants/ants.org
	awk -f ../scripts/preWeave.awk ../ants/ants.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > ants.tex
neighbors.tex: ../neighbors/neighbors.org
	awk -f ../scripts/preWeave.awk ../neighbors/neighbors.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > neighbors.tex
dree.tex: ../dree/dree.org
	awk -f ../scripts/preWeave.awk ../dree/dree.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > dree.tex
land.tex: ../land/land.org
	awk -f ../scripts/preWeave.awk ../land/land.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > land.tex
pickle.tex: ../pickle/pickle.org
	awk -f ../scripts/preWeave.awk ../pickle/pickle.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > pickle.tex
tax.tex: ../tax/tax.org
	awk -f ../scripts/preWeave.awk ../tax/tax.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > tax.tex
tdb.tex: ../tdb/tdb.org
	awk -f ../scripts/preWeave.awk ../tdb/tdb.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > tdb.tex	
tutorial.tex: ../tutorial/tutorial.org
	awk -f ../scripts/preWeave.awk ../tutorial/tutorial.org | bash ../scripts/org2nw | noweave -n -x > tutorial.tex
util.tex: ../util/util.org
	awk -f ../scripts/preWeave.awk ../util/util.org | bash ../scripts/org2nw | noweave -n -x | sed 's/_test/\\_test/g' > util.tex

clean:
	rm -f neighborsDoc.pdf neighborsDoc.ps neighborsDoc.dvi neighborsDoc.aux $(SRC1) 
