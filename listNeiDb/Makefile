version = $(shell bash ../scripts/getVersion.sh)
date = $(shell bash ../scripts/getDate.sh)
nw = $(shell which noweb)
exe = listNeiDb

all : $(exe)

tangle: $(exe).go 
$(exe): $(exe).go
	go build -ldflags "-X github.com/evolbioinf/neighbors/util.version=$(version) -X github.com/evolbioinf/neighbors/util.date=$(date)" $(exe).go
$(exe).go: $(exe).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(exe).org | notangle -R$(exe).go | gofmt > $(exe).go;\
	fi

test:

clean:
	rm -f *.go $(exe)