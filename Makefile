test:
	go test -race ./...

build:
	@GOPROXY=https://mirrors.aliyun.com/goproxy/ go build cmd/cron/main.go
	@echo "Now can run server with ./cron"

docker_build:
	docker run --rm -it -v "${PWD}:/app" -w /app golang:1.14 make build
