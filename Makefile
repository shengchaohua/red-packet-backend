all: setup app_gin_admin app_gin_api

setup:
	go mod vendor
	go mod tidy

app_gin_admin:
	go build -o bin/red-packet-backend-admin ./app/admin

app_gin_api:
	go build -o bin/red-packet-backend-api ./app/api

unittests = $(shell go list ./... | grep -Ev "tests")
unittest:
	@go test -v -short -mod=readonly $(unittests)
