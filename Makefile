progs = ants climt dree land fintac makeNeiDb neighbors outliers pickle taxi
packs = util tdb # Order matters

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
.PHONY: doc test docker
doc:
	make -C doc
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
	make clean -C doc
	make clean -C tutorial
	rm -f bin/*
test: data
	echo test
	for prog in $(packs) $(progs); do \
		make test -C $$prog; \
	done
data:
	wget https://owncloud.gwdg.de/index.php/s/V7DMhBIziwUNqkC/download
	tar -xvzf download
	rm download
docker:
	cd docker; \
	sudo docker build -t neighbors .; \
	cd ..
