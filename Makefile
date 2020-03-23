all: deps validate generate build test clean

build:
	go build ./...

deps:
	npm install autorest

generate:
	autorest --v3 --input-file=clusters.yaml --go --output-folder=clusters --namespace=clusters --user-agent="go-databricks/clusters"

test:
	go test ./...

validate:
	autorest --v3 --input-file=clusters.yaml
