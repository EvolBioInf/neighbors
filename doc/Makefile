SRC1 =   makeNeiDb.tex neighbors.tex dree.tex tax.tex tdb.tex tutorial.tex util.tex
SRC2 =   header.tex introduction.tex neighborsDoc.tex

all: neighborsDoc.pdf
neighborsDoc.pdf: neighborsDoc.tex $(SRC1) $(SRC2)
	latex neighborsDoc.tex
	bibtex neighborsDoc
	latex neighborsDoc.tex
	latex neighborsDoc.tex
	dvips neighborsDoc -o -q
	ps2pdf neighborsDoc.ps

makeNeiDb.tex: ../makeNeiDb/makeNeiDb.org 
	bash ../scripts/org2nw ../makeNeiDb/makeNeiDb.org | noweave -n -x > makeNeiDb.tex
neighbors.tex: ../neighbors/neighbors.org
	bash ../scripts/org2nw ../neighbors/neighbors.org | noweave -n -x  | sed 's/_test/\\_test/g' > neighbors.tex
dree.tex: ../dree/dree.org
	bash ../scripts/org2nw ../dree/dree.org | noweave -n -x | sed 's/_test/\\_test/g' > dree.tex
tax.tex: ../tax/tax.org
	bash ../scripts/org2nw ../tax/tax.org | noweave -n -x | sed 's/_test/\\_test/g' > tax.tex
tdb.tex: ../tdb/tdb.org
	bash ../scripts/org2nw ../tdb/tdb.org | noweave -n -x | sed 's/_test/\\_test/g' > tdb.tex	
tutorial.tex: ../tutorial/tutorial.org
	bash ../scripts/org2nw ../tutorial/tutorial.org | noweave -n -x > tutorial.tex
util.tex: ../util/util.org
	bash ../scripts/org2nw ../util/util.org | noweave -n -x > util.tex

clean:
	rm -f neighborsDoc.pdf neighborsDoc.ps neighborsDoc.dvi neighborsDoc.aux $(SRC1) 
