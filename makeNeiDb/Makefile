exec = makeNeiDb
all : $(exec)

makeNeiDb: $(exec).go
	go build $(exec).go
makeNeiDb.go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec).go | gofmt > $(exec).go
test: $(exec)_test.go
	rm -f test.db
	@go test
$(exec)_test.go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec)_test.go | gofmt > $(exec)_test.go	
clean:
	rm -f *.go

