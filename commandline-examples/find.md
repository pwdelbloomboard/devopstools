### basic find

For finding a script...

```
find ~ -name script.sh 
```

### find on the filesystem

```
find / -name script.sh
```

This will show the pathname on the filesystem.

If you want to avoid errors, do:

```
find / -name script.sh 2>/dev/null
```