exec = ants
all : $(exec)

$(exec): $(exec).go
	go build $(exec).go
$(exec).go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec).go | gofmt > $(exec).go
test: $(exec)_test.go
	@go test
$(exec)_test.go: $(exec).org
	bash ../scripts/org2nw $(exec).org | notangle -R$(exec)_test.go | gofmt > $(exec)_test.go	
clean:
	rm -f *.go

