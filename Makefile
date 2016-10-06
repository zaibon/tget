include credentials.mk

all: build

install: generate
	go install

build: generate
	go build

generate:
	go run scripts/extrac_term.go -login '$(T411_USERNAME)' -password '$(T411_PASSWORD)'
	go generate ./api/t411

test: generate
	go test ./api/...

clean:
	rm scripts/mapping.json api/t411/bindata.go tget
