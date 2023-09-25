progs = ants dree land makeNeiDb neighbors outliers pickle taxi
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
.PHONY: doc test docker
doc:
	make -C doc
clean:
	for prog in $(progs) $(packs) doc; do \
		make clean -C $$prog; \
	done
	make clean -C doc
test:
	echo test
	for prog in $(progs) $(packs); do \
		make test -C $$prog; \
	done
docker:
	git clone https://github.com/IvanTsers/neighbors-docker; \
	cd neighbors-docker; \
	sudo docker build -t neighbors .; \
	cd ..
