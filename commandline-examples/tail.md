# tail

* Tail is basically a command which shows the end of a file, the last 10 lines by default.
* Tail can also be used on logs as a way to follow a log, since ultimately a log file is a continuously annotated file.

## Example Usage

The most common usage within devops or server administration is: tail -f which follows a streaming log. Fundamentally tail is working the same using -f as it would be with other options explained below, but the -f flag is used specifically for that purpose.

We can for example, tail this exact file that we're working with by moving into the directory in question and using the command:

```
tail tail.md
```

The results of that command should be the last 10 lines:

```
### Last Ten Lines for Sample Testing

This here...
is the...
last 10...
or so...
plus or minus...
if you can double check...
countable...
lines.
```
You can adjust the number of lines being tailed with -n, so if you use, -n 11 "file.whatever" you will see the lines above it, in our case, we hid, "hehehe" above it so you should be able to see that using -n 12.

There is also an -n +NUMBER feature which allows the tail feature to start printing off lines starting at a particular point in the file forward, rather than from the bottom, so for example at the time of writing, our #hehehe started on line 30, so we could completely print ou the end of the file with:

```
tail -n +32 tail.md
```

The tail command can also use "-c" to count off by bytes rather than by characters.

```
tail -c -50 tail.md
...
if you can double check...
countable...
lines
```

So basically the above prints off the last 50 bytes rather than the last 50 characters.

-v is a nice little flag that just prints out the file name between characters at the start of the tail output, analogous in a way to the, "version" since the file is kind of like the, "exact thing," we are tailing, kind of like the version of software points to an exact instance of said software.:

```
$ tail -v tail.md
==> tail.md <==
### Last Ten Lines for Sample Testing
...

```

That being said, -v is not the version of tail, --version is!

```
$ tail --version
tail (GNU coreutils) 8.32
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Written by Paul Rubin, David MacKenzie, Ian Lance Taylor,
and Jim Meyering.
```


## hehehe
### Last Ten Lines for Sample Testing

This here...
is the...
last 10...
or so...
plus or minus...
if you can double check...
countable...
lines.