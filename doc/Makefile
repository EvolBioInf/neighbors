SRC1 = ants.tex climt.tex clusters.tex dree.tex fintac.tex listNeiDbs.tex land.tex makeNeiDb.tex mrca.tex \
neighbors.tex outliers.tex pickle.tex ranks.tex taxi.tex tdb.tex tutorial.tex util.tex
SRC2 =   header.tex introduction.tex neighborsDoc.tex labels.tex covTree.tex

date = $(shell git log | grep -m 1 Date | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)

all: $(SRC1) neighborsDoc.pdf 
%.tex:	../*/%.org neighborsDoc.tex
	bash ../scripts/org2nw $< | awk -f ../scripts/preWeave.awk | noweave -n -x | sed 's/_test/\\_test/g' > $@
neighborsDoc.pdf: neighborsDoc.tex $(SRC1) $(SRC2) db.ps
	echo $(date) | tr '_' ' ' > date.txt
	echo $(version) | tr '-' ' ' | awk '{printf "%s", $$1; if ($$2) printf "-%s", $$2; printf "\n"}' > version.txt
	latex neighborsDoc.tex
	bibtex neighborsDoc
	latex neighborsDoc.tex
	latex neighborsDoc.tex
	dvipdf -dALLOWPSTRANSPARENCY neighborsDoc
publish: neighborsDoc.pdf
	if [ -d ~/WinHome/ownCloud/ ]; then \
		cp neighborsDoc.pdf ~/WinHome/ownCloud/docs/; \
	fi

db.ps: db.dot
	neato -T eps db.dot > db.ps

clean:
	rm -f neighborsDoc.pdf neighborsDoc.dvi neighborsDoc.aux $(SRC1) db.ps
