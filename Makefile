.DEFAULT_GOAL := build

./bin:
	@mkdir bin

./bin/fart: ./bin
	go build -o ./bin/fart ./cmd/...

.PHONY: build
build: ./bin/fart

.PHONY: clean
clean:
	rm -f ./bin/*
