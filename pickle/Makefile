version = $(shell bash ../scripts/getVersion.sh)
date = $(shell bash ../scripts/getDate.sh)
nw = $(shell which noweb)
pro = pickle
exe = cmd/$(pro)

all : $(exe)

tangle: $(pro).go $(pro)_test.go
$(exe): $(pro).go
	go build -ldflags "-X github.com/evolbioinf/neighbors/util.version=$(version) -X github.com/evolbioinf/neighbors/util.date=$(date)" $(pro).go
	cd cmd && go build -o $(pro) main.go
$(pro).go: $(pro).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(pro).org | notangle -R$(pro).go | gofmt > $(pro).go;\
	fi
test: $(pro)_test.go $(exe)
	@go test
$(pro)_test.go: $(pro).org
	if [ "$(nw)" != "" ]; then\
		bash ../scripts/org2nw $(pro).org | notangle -R$(pro)_test.go | gofmt > $(pro)_test.go;\
	fi
clean:
	rm -f *.go $(exe)

