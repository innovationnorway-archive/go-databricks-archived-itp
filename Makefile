all: deps validate generate build test

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --clusters
	autorest --dbfs
	autorest --groups
	autorest --workspace

test:
	go test ./...

validate:
	autorest --input-file=clusters.yaml
	autorest --input-file=dbfs.yaml
	autorest --input-file=groups.yaml
	autorest --input-file=workspace.yaml
