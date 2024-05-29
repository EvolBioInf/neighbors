all : tdb

tdb:
	go build tdb.go
tangle: tdb.go tdb_test.go
tdb.go: tdb.org
	awk -f ../scripts/preTangle.awk tdb.org | bash ../scripts/org2nw | notangle -Rtdb.go | gofmt > tdb.go
test: tdb.go tdb_test.go
	rm -f ../data/taxSmall.db
	go test -v
tdb_test.go: tdb.org
	awk -f ../scripts/preTangle.awk tdb.org | bash ../scripts/org2nw | notangle -Rtdb_test.go | gofmt > tdb_test.go
clean:
	rm -f *.go

