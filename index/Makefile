all : index

index: index.go
	go build index.go
index.go: index.org
	bash ../scripts/org2nw index.org | notangle -Rindex.go | gofmt > index.go
clean:
	rm -f *.go

