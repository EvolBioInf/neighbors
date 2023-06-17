SRC1 =   makeNeiDb.tex taxi.tex ants.tex neighbors.tex dree.tex findMacs.tex land.tex pickle.tex tax.tex tdb.tex tutorial.tex util.tex
SRC2 =   header.tex introduction.tex neighborsDoc.tex

all: neighborsDoc.pdf
neighborsDoc.pdf: neighborsDoc.tex $(SRC1) $(SRC2) ../findMacs/int.tex
	latex neighborsDoc.tex
	bibtex neighborsDoc
	latex neighborsDoc.tex
	latex neighborsDoc.tex
	dvips neighborsDoc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY neighborsDoc.ps

makeNeiDb.tex: ../makeNeiDb/makeNeiDb.org 
	bash ../scripts/org2nw ../makeNeiDb/makeNeiDb.org | noweave -n -x | sed 's/_test/\\_test/g' > makeNeiDb.tex
taxi.tex: ../taxi/taxi.org 
	bash ../scripts/org2nw ../taxi/taxi.org | noweave -n -x | sed 's/_test/\\_test/g' > taxi.tex
ants.tex: ../ants/ants.org
	bash ../scripts/org2nw ../ants/ants.org | noweave -n -x | sed 's/_test/\\_test/g' > ants.tex
neighbors.tex: ../neighbors/neighbors.org
	bash ../scripts/org2nw ../neighbors/neighbors.org | noweave -n -x  | sed 's/_test/\\_test/g' > neighbors.tex
dree.tex: ../dree/dree.org
	bash ../scripts/org2nw ../dree/dree.org | noweave -n -x | sed 's/_test/\\_test/g' > dree.tex
land.tex: ../land/land.org
	bash ../scripts/org2nw ../land/land.org | noweave -n -x | sed 's/_test/\\_test/g' > land.tex
findMacs.tex: ../findMacs/findMacs.org
	bash ../scripts/org2nw ../findMacs/findMacs.org | noweave -n -x > findMacs.tex
pickle.tex: ../pickle/pickle.org
	bash ../scripts/org2nw ../pickle/pickle.org | noweave -n -x | sed 's/_test/\\_test/g' > pickle.tex
tax.tex: ../tax/tax.org
	bash ../scripts/org2nw ../tax/tax.org | noweave -n -x | sed 's/_test/\\_test/g' > tax.tex
tdb.tex: ../tdb/tdb.org
	bash ../scripts/org2nw ../tdb/tdb.org | noweave -n -x | sed 's/_test/\\_test/g' > tdb.tex	
tutorial.tex: ../tutorial/tutorial.org
	bash ../scripts/org2nw ../tutorial/tutorial.org | noweave -n -x > tutorial.tex
util.tex: ../util/util.org
	bash ../scripts/org2nw ../util/util.org | noweave -n -x | sed 's/_test/\\_test/g' > util.tex

clean:
	rm -f neighborsDoc.pdf neighborsDoc.ps neighborsDoc.dvi neighborsDoc.aux $(SRC1) 
