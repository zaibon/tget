include credentials.mk

all: build

install: generate
	go install

build: generate
	go build

generate:
	go run scripts/extrac_term.go -login '$(T411_USERNAME)' -password '$(T411_PASSWORD)'
	go generate ./api/t411

test:
	go test ./...

clean:
	rm scripts/mapping.json tget
