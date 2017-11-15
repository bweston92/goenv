test:
	go test -race

coverage:
	go test -race -coverprofile=coverage.txt -covermode=atomic

deps:
	go get -t -v ./...

.PHONY: test deps coverage
