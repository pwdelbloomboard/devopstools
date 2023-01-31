# Original Link

https://github.com/TomWright/dasel

## About Dasel

> Dasel (short for data-selector) allows you to query and modify data structures using selector strings.
> Comparable to jq / yq, but supports JSON, YAML, TOML, XML and CSV with zero runtime dependencies.

## Installation

On mac, ```brew install dasel```

On Linux:

```
curl -sSLf "$(curl -sSLf https://api.github.com/repos/tomwright/dasel/releases/latest       | \
    grep browser_download_url                                                               | \ 
    grep linux_amd64                                                                        | \ 
    grep -v .gz                                                                             | \ 
    cut -d\" -f 4)" -L -o dasel                                                            && \
    chmod +x dasel

mv ./dasel /usr/local/bin/dasel
```


$(curl -sSLf https://github.com/TomWright/dasel/releases/download/v2.1.1/dasel_darwin_amd64)

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

### Using Dasel on a File

```
PROJECTNAME=$(dasel -f whatever.toml '.project.name')
```