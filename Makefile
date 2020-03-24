all: deps validate generate build test clean

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --clusters

test:
	go test ./...

validate:
	autorest --input-file=clusters.yaml
