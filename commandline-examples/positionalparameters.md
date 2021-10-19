# 

```
$1
```

> $1 is the first command-line argument passed to the shell script. Also, know as Positional parameters. For example, $0, $1, $3, $4 and so on. If you run: 

```
./script filename1 dir1
```

then:

* $0 is the name of the script itself (script.sh)
* $1 is the first argument (filename1)
* $2 is the second argument (dir1)
* $9 is the ninth argument
* ${10} is the tenth argument and must be enclosed in brackets after $9.
* ${11} is the eleventh argument.

Example

```
#!/bin/bash
script="$0"
first="$1"
second="$2"
tenth="${10}"
echo "The script name : $script"
echo "The first argument :  $first"
echo "The second argument : $second"
echo "The tenth and eleventh argument : $tenth and ${11}"
```
Then:

```
chmod +x demo-args.sh
./demo-args.sh foo bar one two a b c d e f z f z h
```

This will print out the arguments as shown above, including "z" as the 11th argument.

So basically these are numbered arguments fed into the command.

So with, "./tree build x" - the numbered arguments would be: $0, $1, $2, with $0 being "tree" which is the name of the shell script itself. $1 would be build, which could have some special meaning in our script, and x could be perhaps a directory or have some other special meaning in our script.

# Resources

* [Positional Parameters](https://bash.cyberciti.biz/guide/$1)