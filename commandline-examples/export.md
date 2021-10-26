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

# Example

```
export CLIENT_CURRENT_DOMAIN="$DNS"
```

The above takes the bash variable, $DNS and sets the environmental variable, CLIENT_CURRENT_DOMAIN.  This CLIENT_CURRENT_DOMAIN can then be used later in subsequent functions called on that particular server (or client in the case of client-side javascript).

In this situation:

* name is CLIENT_CURRENT_DOMAIN
* =[value] is ="$DNS"
