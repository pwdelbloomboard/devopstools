# Set

> set command allows you to change the value of shell options or to display the names and values of shell variables. Rarely used, it is a bash builtin, but is quite a bit more complicated than most builtins.

```
set -o errexit
```

Same as -e Exit immediately if the pipeline (single or multiple commands) returns a non-zero status. Does not effect a while or until keyword, or part of a test/if statement, part of any command executes in a && or || list. If a compound command other than a sub shell returns a non-zero status because command failed while -e was being ignored, shell does not exit. A trap on ERR, if set is executed before the shell options. Applies to the shell environment and each sub shell environment separately.

```
set -o errtrace
```

Same as -E ... any trap on ERR is inherited by shell functions, command substitutions, and command executed in subshell environment.

```
set -o nounset
```

Same as -u Treat unset variables and parameters other than special parameters * or @ as an error when performing parameter expansion. Will be written to the standard error, non-interactive shell will exit.

```
set -o pipefail
```

Returns the value of a pipeline, the value of the last (rightmost) command to exit with a non-zero status, or zero if all commands in the pipeline exit successfully. Option is disabled by default.

if [[ -n ${DEBUG_MODE+x} ]]; then
    set -o xtrace
    set -o functrace
fi

xtrace is the same as -x


# Resources

* [Using set](https://www.networkworld.com/article/3631415/using-the-linux-set-command.html)
* [gnu - set](https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html )