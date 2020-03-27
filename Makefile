all: deps validate generate build test

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --clusters
	autorest --groups

test:
	go test ./...

validate:
	autorest --input-file=clusters.yaml
	autorest --input-file=groups.yaml
