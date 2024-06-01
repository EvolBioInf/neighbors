version = $(shell bash ../scripts/getVersion.sh)
date = $(shell bash ../scripts/getDate.sh)

exe = land
all : $(exe)

tangle: $(exe).go $(exe)_test.go
$(exe):
	go build -ldflags "-X github.com/evolbioinf/neighbors/util.version=$(version) -X github.com/evolbioinf/neighbors/util.date=$(date)" $(exe).go
$(exe).go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go | gofmt > $(exe).go
test: $(exe)_test.go $(exe)
	@go test
$(exe)_test.go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go	
clean:
	rm -f *.go $(exe)
