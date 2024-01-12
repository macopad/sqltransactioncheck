BIN := bin

PHONY: build install test

$(BIN):
	mkdir -p $@

build: $(BIN)
	go build -o $(BIN)/sqltransactioncheck .

install:
	go install

test: build
	go test ./...
	# Due to an issue with importing in a anaylsistest's test data some hoop jumping is required
	# I call twice to avoid collecting package downloads in output
	-go vet -vettool=$(BIN)/sqltransactioncheck ./testdata/trans_examples1
	-go vet -vettool=$(BIN)/sqltransactioncheck ./testdata/trans_examples1 2> trans_examples_results.txt
	diff -a trans_examples_results.txt ./testdata/trans_examples1/expected_results.txt

lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.54.1
	./bin/golangci-lint run
