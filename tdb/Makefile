all : tdb

tdb: tdb.go
	go build tdb.go
tdb.go: tdb.org
	bash ../scripts/org2nw tdb.org | notangle -Rtdb.go | gofmt > tdb.go
clean:
	rm -f *.go

