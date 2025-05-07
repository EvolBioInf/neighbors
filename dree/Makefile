version = $(shell bash ../scripts/getVersion.sh)
date = $(shell bash ../scripts/getDate.sh)
nw = $(shell which noweb)
exe = dree

all : $(exe) $(exe).sh

tangle: $(exe).go $(exe)_test.go
$(exe): $(exe).go
	go build -ldflags "-X github.com/evolbioinf/neighbors/util.version=$(version) -X github.com/evolbioinf/neighbors/util.date=$(date)" $(exe).go
$(exe).go: $(exe).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go | gofmt > $(exe).go;\
	fi
$(exe).sh: $(exe).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(exe).org | notangle -R$(exe).sh > $(exe).sh;\
	fi
test: $(exe)_test.go $(exe)
	@go test
$(exe)_test.go: $(exe).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(exe).org | notangle -R$(exe)_test.go | gofmt > $(exe)_test.go;\
	fi
clean:
	rm -f *.go $(exe) $(exe).sh

