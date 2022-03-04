all : util

util: util.go
	go build util.go
util.go: util.org
	bash ../scripts/org2nw util.org | notangle -Rutil.go | gofmt > util.go
clean:
	rm -f *.go

