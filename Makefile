get:
	go get -u -d ./...

run:
	go run cmd/skink/main.go

init:
	go run cmd/skink/main.go init

specs:
	ginkgo ./...