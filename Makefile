all: deps validate generate build test clean

build:
	go build ./...

clean:
	rm -f swagger

deps:
	go get ./...
	wget https://github.com/go-swagger/go-swagger/releases/download/v0.22.0/swagger_linux_amd64 -O swagger
	chmod +x swagger

generate:
	./swagger generate client -A databricks -c plumbing

test:
	go test ./...

validate:
	./swagger validate swagger.yaml
