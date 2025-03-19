nw = $(shell which noweb)

all : tdb

tdb: tdb.go
	go build tdb.go
tangle: tdb.go tdb_test.go
tdb.go: tdb.org
	if [ "$(nw)" != "" ]; then\
		awk -f ../scripts/preTangle.awk tdb.org | bash ../scripts/org2nw | notangle -Rtdb.go | gofmt > tdb.go;\
	fi
test: tdb.go tdb_test.go
	rm -f ../data/taxSmall.db
	go test -v
tdb_test.go: tdb.org
	if [ "$(nw)" != "" ]; then\
		awk -f ../scripts/preTangle.awk tdb.org | bash ../scripts/org2nw | notangle -Rtdb_test.go | gofmt > tdb_test.go;\
	fi
clean:
	rm -f *.go

