all: setup app_admin app_api

setup:
	mkdir -p log
	go mod vendor
	go mod tidy

app_admin:
	go build -o bin/admin ./app/admin

app_api:
	go build -o bin/api ./app/api

unittests = $(shell go list ./... | grep -Ev "tests")
unittest:
	@go test -v -short -mod=readonly $(unittests)

# Dockerfile
docker_admin:
    docker build -t red_packer_backend_admin:$version -f Dockerfile-admin .

docker_api:
    docker build -t red_packer_backend_api:$version -f Dockerfile-api .
