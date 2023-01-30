# Original Link

https://github.com/TomWright/dasel

## About Dasel

> Dasel (short for data-selector) allows you to query and modify data structures using selector strings.
> Comparable to jq / yq, but supports JSON, YAML, TOML, XML and CSV with zero runtime dependencies.

## Installation

On mac, ```brew install dasel```

### Reading

```
echo '{"name": "Tom"}' | dasel -r json 'name'
"Tom"
```

### Writing

echo '{"name": "Tom"}' | dasel put -r json -t string -v 'contact@tomwright.me' 'email'
{
  "email": "contact@tomwright.me",
  "name": "Tom"
}

### Deletion

```
echo '{  "email": "contact@tomwright.me", "name": "Tom"}' | dasel delete -r json '.email'
{
  "name": "Tom"
}
```