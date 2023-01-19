export GO111MODULE=on
export GOPROXY=direct
export GOSUMDB=off
export COMPOSE_PROJECT_NAME=sapper

.PHONY: unit_test
unit_test:
	go test -v -cover `go list ./... | grep -v system_test`

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run:
	go run ./
