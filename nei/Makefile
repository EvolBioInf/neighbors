all : nei

nei: nei.go
	go build nei.go
nei.go: nei.org
	bash ../scripts/org2nw nei.org | notangle -Rnei.go | gofmt > nei.go
clean:
	rm -f *.go

