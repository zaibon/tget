
all: build

install: generate
	go install

build: generate
	go build

generate:
	scripts/extract_term.py
	go generate ./...

test: generate
	go test ./...
