exec = dree
all : $(exec)

$(exec): 
	go build $(exec).go
tangle: $(exec).go $(exec)_test.go
$(exec).go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec).go | gofmt > $(exec).go
test: $(exec)_test.go $(exec)
	@go test
$(exec)_test.go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec)_test.go | gofmt > $(exec)_test.go	
clean:
	rm -f *.go $(exec)
