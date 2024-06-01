version = $(shell bash ../scripts/getVersion.sh)
date = $(shell bash ../scripts/getDate.sh)

exe = makeNeiDb
all : $(exe)

makeNeiDb:
	go build -ldflags "-X github.com/evolbioinf/neighbors/util.version=$(version) -X github.com/evolbioinf/neighbors/util.date=$(date)" $(exe).go
tangle: $(exe).go $(exe)_test.go
makeNeiDb.go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go | gofmt > $(exe).go
test: $(exe)_test.go $(exe)
	rm -f test.db
	@go test
$(exe)_test.go: $(exe).org
	bash ../scripts/org2nw $(exe).org | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go	
clean:
	rm -f *.go $(exe)
