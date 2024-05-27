.PHONY: test all clean

build:
	@go build -o bin/gof cmd/gof/main.go

run:build
	@bin/gof


