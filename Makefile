progs = makeNeiDb ants climt dree land fintac neighbors outliers pickle taxi
packs = util tdb # Order matters
ftp = ftp.ncbi.nlm.nih.gov

all:
	test -d bin || mkdir bin
	for pack in $(packs); do \
		make -C $$pack; \
	done
	for prog in $(progs); do \
		make -C $$prog; \
		cp $$prog/$$prog bin; \
	done
tangle:
	for pack in $(packs); do \
		make tangle -C $$pack; \
	done
	for prog in $(progs); do \
		make tangle -C $$prog; \
	done
	make tangle -C tutorial
.PHONY: data doc test docker
doc:
	make -C doc
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
	make clean -C doc
	make clean -C tutorial
	rm -f bin/*
test:
	echo test
	for prog in $(packs) $(progs); do \
		make test -C $$prog; \
	done
data:
	mkdir -p data
	cd data; wget $(ftp)/pub/taxonomy/taxdump.tar.gz; \
		tar -xvzf taxdump.tar.gz; \
		wget $(ftp)/genomes/ASSEMBLY_REPORTS/assembly_summary_genbank.txt; \
		wget $(ftp)/genomes/ASSEMBLY_REPORTS/assembly_summary_refseq.txt
docker:
	cd docker; \
	sudo docker build -t neighbors .; \
	cd ..
