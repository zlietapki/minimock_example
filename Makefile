.PHONY: mock
mock:
	go generate ./...

.PHONY: test
test:
	go test -v -count 1 ./...
