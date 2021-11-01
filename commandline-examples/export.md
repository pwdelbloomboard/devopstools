# Export

export: export [-fn] [name[=value] ...] or export -p

* Set export attribute for shell variables.
    
* Marks each NAME for automatic export to the environment of subsequently executed commands.  If VALUE is supplied, assign VALUE before exporting.
    
* Options:

      -f	refer to shell functions

      -n	remove the export property from each NAME

      -p	display a list of all exported variables and functions
    
    An argument of `--' disables further option processing.
    
    Exit Status:

    Returns success unless an invalid option is given or NAME is invalid.

## Example

```
export CLIENT_CURRENT_DOMAIN="$DNS"
```

The above takes the bash variable, $DNS and sets the environmental variable, CLIENT_CURRENT_DOMAIN.  This CLIENT_CURRENT_DOMAIN can then be used later in subsequent functions called on that particular server (or client in the case of client-side javascript).

In this situation:

* name is CLIENT_CURRENT_DOMAIN
* =[value] is ="$DNS"

## Using Debian Bullseye Docker Server

We can quickly spin up a debian-11 (bullseye docker image) with the following command and test out export:

```
docker run -i -t debian:bullseye-slim /bin/bash
```
* Note - the command, [source](commandline-examples/source) runs the exportscript.script below.
* the command "set" shows the environmental variables, we grep it to filter for THIS_VARIABLE.

```
/# echo 'export THIS_VARIABLE=1234' > exportscript.sh

/# cat exportscript.sh
export THIS_VARIABLE=1234

/# chmod +x exportscript.sh

/# source exportscript.sh 

/# set | grep THIS_VARIABLE
THIS_VARIABLE=1234

```