all : makeNeiDb

makeNeiDb: makeNeiDb.go
	go build makeNeiDb.go
makeNeiDb.go: makeNeiDb.org
	bash ../scripts/org2nw makeNeiDb.org | notangle -RmakeNeiDb.go | gofmt > makeNeiDb.go
clean:
	rm -f *.go

