all: setup app_gin_admin app_gin_api modtidy

setup:
	go mod vendor

app_gin_admin:
	go build -o bin/app_gin/red-packet-backend-admin ./app_gin/admin

app_gin_api:
	go build -o bin/app_gin/red-packet-backend-api ./app_gin/api

modtidy:
	go mod tidy
