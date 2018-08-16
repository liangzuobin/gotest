all: build

run:
	ENV=prod SYSLOG_ADDR=127.0.0.1:1514 APP_NAME=titanlog go run main.go

build:
	go build ./...

prod:
	GOOS=linux go build ./...

test:
	go test -v ./...

clean:
	rm -f main titanlog

.PHONY: all main
