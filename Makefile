.PHONY: run

get:
	go get -u -d ./...

run:
	go run cmd/skink/main.go

run-dev:
	go run cmd/skink/main.go -config=etc/skink.conf
