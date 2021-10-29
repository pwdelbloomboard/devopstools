# Source

Source: source filename [arguments]

Execute commands from a file in the current shell.
    
Read and execute commands from FILENAME in the current shell.  The entries in $PATH are used to find the directory containing FILENAME. If any ARGUMENTS are supplied, they become the positional parameters when FILENAME is executed.
    
Exit Status:

Returns the status of the last command executed in FILENAME; fails if FILENAME cannot be read.

## More About Source

* source and . (a period) are the same command.
* If the FILENAME is not a full path to a file, the command will search for the file in the directories specified in the $PATH environmental variable . If the file is not found in the $PATH, the command will look for the file in the current directory.
* If any ARGUMENTS are given, they will become positional parameters to the FILENAME.
* If the FILENAME exists, the source command exit code is 0, otherwise, if the file is not found it will return 1.

## A Quick Warning ./ vs Source

* ./ and source are not quite the same.

* ./script runs the script as an executable file, launching a new shell to run it
* source script reads and executes commands from filename in the current shell environment

> Both sourcing and executing the script will run the commands in the script line by line, as if you typed those commands by hand line by line.

> The differences are:

> When you execute the script you are opening a new shell, type the commands in the new shell, copy the output back to your current shell, then close the new shell. Any changes to environment will take effect only in the new shell and will be lost once the new shell is closed.
> When you source the script you are typing the commands in your current shell. Any changes to the environment will take effect and stay in your current shell. 
> the "environment" are things like the current working directory and environment variables. also shell settings (among others history and completion features). there are more but those are the most visible.

> Use source if you want the script to change the environment in your currently running shell. use execute otherwise.


## Example

1. First need to change a script permission to be executable with chmod +x
2. The below function, "check_root" will check and print if it's being run as root and print the location.

```
#!/bin/bash

say_hello () {
  echo "hello world";
}

say_goodbye () {
  echo "goodbye.";
}

"$@"
```
* Note the #!/bin/bash helps determine the shell which will be created by, "source" as /bin/bash as opposed to /bin/sh if not already in that shell.

3. Run, "source helloworld.sh say_hello"

The script should print out, "hello world" - this is being executed from the current shell environment.

4. Run ". helloworld.sh say_hello"

This script should print out, "hello world" - this is being executed from a *new shell environment.