appName=shopmall
VERSION=$(shell git describe --tags --always 2>/dev/null || echo "v0.0.1")
BUILDTIME=$(shell date -u '+%Y-%m-%d %H:%M:%S')
LDFLAGS=-X "main.Version=$(VERSION)" -X "main.BuildTime=$(BUILDTIME)"

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags '$(LDFLAGS)' -o ./bin/$(appName) ./cmd/server

.PHONY: run
# run
run: build
	./bin/$(appName)

.PHONY: clean
# clean
clean:
	rm -rf bin/
	rm -rf coverage.out

.PHONY: fmt
# format go code
fmt:
	gofmt -s -w .
	go mod tidy

.PHONY: lint
# lint go code
lint:
	golangci-lint run ./...

.PHONY: test
# run unit test
test:
	go test -v ./...

.PHONY: cover
# run unit test with coverage
cover:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: docker
# build docker image
docker:
	docker build -t shopmall:$(VERSION) .

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help