all: app_gin_admin

setup:
	go mod vendor

app_gin_admin:
	go build -o bin/app_gin/red-packet-backend-admin ./app_gin/admin