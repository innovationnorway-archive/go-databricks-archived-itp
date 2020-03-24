# Configuration

> see https://aka.ms/autorest

``` yaml
azure-arm: false
```

## Clusters

``` yaml $(clusters)
input-file: clusters.yaml

go:
  namespace: clusters
  output-folder: clusters
  package-name: github.com/innovationnorway/go-databricks/clusters
  user-agent: 'go-databricks/version clusters'
```
