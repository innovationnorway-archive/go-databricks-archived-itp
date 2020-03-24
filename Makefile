all: deps validate generate build test clean

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --v3 --clusters

test:
	go test ./...

validate:
	autorest --v3 --input-file=clusters.yaml
