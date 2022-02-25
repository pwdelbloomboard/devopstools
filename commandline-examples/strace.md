# strace

https://man7.org/linux/man-pages/man1/strace.1.html

Strace Overview
strace can be seen as a light weight debugger. It allows a programmer / user to quickly find out how a program is interacting with the OS. It does this by monitoring system calls and signals.

Uses
Good for when you don't have source code or don't want to be bothered to really go through it.
Also, useful for your own code if you don't feel like opening up GDB, but are just interested in understanding external interaction.


## Tutorial

http://timetobleed.com/hello-world/

So basically, you can write a program, perhaps a hello.c program or perhaps a hello.py or hello.rb program and then run strace on that program.

1. Create the c program with nano hello.c

hello.c:

```
#include <stdio.h>
int
main(int argc, char *argv[])
{
  printf("hi!\n");
  return 0;
}
```

2. Save, then compile with "gcc hello.c"

3. Run the c program with: "./a.out" which is the default saved name, otherwise change the name and do, "./changedname"

Now you have a c file.

You can run an strace of that file with:

```
strace -ttT ./changedname
```

Which then will give an output showing how your computer runs that program and what happens, each step, at each part of the process, with a timestamp, for example:.

```
19:18:25.825481 execve("./hello", ["./hello"...], [/* 33 vars */]) = 0 <0.000071>
```

The first point is the timestamp when this was run, 19:18.<seconds> and the remainder shows the step.

So in summary, an extremely accurate view of a program runtime can be given.

## Usage

In the simplest case strace runs the specified command until it
       exits.  It intercepts and records the system calls which are
       called by a process and the signals which are received by a
       process.  The name of each system call, its arguments and its
       return value are printed on standard error or to the file
       specified with the -o option.
