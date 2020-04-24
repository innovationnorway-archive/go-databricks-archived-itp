all: deps validate generate build test

build:
	go build ./...

deps:
	npm install -g autorest

generate:
	autorest --clusters
	autorest --groups
	autorest --workspace
	autorest --secrets

test:
	go test ./...

validate:
	autorest --input-file=clusters.yaml
	autorest --input-file=groups.yaml
	autorest --input-file=workspace.yaml
	autorest --input-file=secrets.yaml