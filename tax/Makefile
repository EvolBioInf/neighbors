all : tax

tax: tax.go
	go build tax.go
tax.go: tax.org
	awk -f ../scripts/preTangle.awk tax.org | bash ../scripts/org2nw | notangle -Rtax.go | gofmt > tax.go
test: tax.go tax_test.go
	go test -v
tax_test.go: tax.org
	awk -f ../scripts/preTangle.awk tax.org | bash ../scripts/org2nw | notangle -Rtax_test.go | gofmt > tax_test.go

clean:
	rm -f *.go

