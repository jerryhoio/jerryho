.PHONY: test
test:
	go test -v ./...
.PHONY: lint
lint:
	golangci-lint run
.PHONY: licences-check
licences-check:
	licensei check
.PHONY: check
check: test lint licences-check
