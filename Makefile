.PHONY: build
build:
	mkdir -p bin
	go build -trimpath -o bin/single-dev-env

.PHONY: run
run:
	go run main.go