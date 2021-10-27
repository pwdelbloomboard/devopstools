# Online Bash Simulators

You can use a bash or shell simulator if you're on something else like zsh.

* [tutorialspoint](https://www.tutorialspoint.com/execute_bash_online.php)

* [cocalc](https://cocalc.com/projects/0983a0b1-04ab-4392-ae4d-a35c849ef0e1/files/Welcome%20to%20CoCalc.ipynb?session=default)


# Important Notes
## Merging Command Flags

Note that whenever a command line flat is followed by a value, the value for the flag must always follow the flag.

Example...in the following case, you could never merge the flags together as -f1f2 because they have values following them by default, they are expecting an input, so you have to be explicit:

```
command --flag1 <thing1> --flag2 <thing2>
```

# Extremely Common Tools

## Make Executable

```
chmod +x
```

## Check if Executable

```
ls -l
```

![](/img/permissions.png)


# Introduction

* [grep](/commandline-examples/grep.md)
* [positionalparameters](/commandline-examples/positionalparameters.md)
* [set](/commandline-examples/set.md)