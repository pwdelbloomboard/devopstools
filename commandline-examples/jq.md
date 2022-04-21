# JQ

Similar to yq (yaml processor) but it's for json rather than yaml.

Nice little tool that allows you to pretty print a json file.

The main example given is, if you had something called, "names.json" and wanted to print it out in a nice-looking, readable json format:

```
[{"id": 1, "name": "Arthur", "age": "21"},{"id": 2, "name": "Richard", "age": "32"}]
```

The, "." filter takes the entire input and sends it to STDOUT.

So for example, the following command shows everything:

```
$ echo '[{"id": 1, "name": "Arthur", "age": "21"},{"id": 2, "name": "Richard", "age": "32"}]' | jq '.'
[
  {
    "id": 1,
    "name": "Arthur",
    "age": "21"
  },
  {
    "id": 2,
    "name": "Richard",
    "age": "32"
  }
]
```

If you had a huge file with the need to inspect with an interface, you might use, "less."

```
$ echo '[{"id": 1, "name": "Arthur", "age": "21"},{"id": 2, "name": "Richard", "age": "32"}]' | jq '.' | less
```

If you want to print out something specific in the data structure, look at the structure itself and access it with dot notation as follows:

```
echo '[{"id": 1, "name": "Arthur", "age": "21"},{"id": 2, "name": "Richard", "age": "32"}]' | jq '.[]' | jq '.name'
"Arthur"
"Richard"
```


* JQ can also perform all sorts of mapping and wrangling, cleaning data with its various options.

# Resources

* [Using JQ](https://shapeshed.com/jq-json/)