.PHONY: run
run:
	go run cmd/gitthub_parser/main.go 

.PHONY: build
build:
	go build -o github-parser.exe cmd/gitthub_parser/main.go