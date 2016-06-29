
all: build

install: generate
	go install

build: generate
	go build

generate: scripts/t411_terms.json scripts/extract_term.py
	scripts/extract_term.py
	go generate ./api/t411

test: generate
	go test ./api/...

clean:
	rm scripts/mapping.json api/t411/bindata.go tget
