https://faun.pub/golangs-fmt-sprintf-and-printf-demystified-

> Here’s a simple example of Go code without fmt.Sprintf()

```
myString := "Results: " + results + " and more: " + more + "."
```

> We can achieve the same thing with cleaner code using fmt.Sprintf()

```
myString := fmt.Sprintf("Results: %s and more: %s.", results, more)
```