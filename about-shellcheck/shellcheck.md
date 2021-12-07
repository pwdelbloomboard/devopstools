## About Shellcheck

https://github.com/koalaman/shellcheck

> The goals of ShellCheck are:

* To point out and clarify typical beginner's syntax issues that cause a shell to give cryptic error messages.

* To point out and clarify typical intermediate level semantic problems that cause a shell to behave strangely and counter-intuitively.

* To point out subtle caveats, corner cases and pitfalls that may cause an advanced user's otherwise working script to fail under future circumstances.

### Shellcheck Browser Tool

https://www.shellcheck.net/

### Within Builds and Test Suites

From the shellcheck documentation:

> While ShellCheck is mostly intended for interactive use, it can easily be added to builds or test suites. It makes canonical use of exit codes, so you can just add a shellcheck command as part of the process.

## Using Shellcheck on Local

After installation, Shellcheck can be invoked on local with:

```
shellcheck SCRIPT
```

With script being the executable script that you wish to evaluate.

### Shellcheck Help Options

Usage: shellcheck [OPTIONS...] FILES...

  -a                  --check-sourced            Include warnings from sourced files

  -C[WHEN]            --color[=WHEN]             Use color (auto, always, never)

  -i CODE1,CODE2..    --include=CODE1,CODE2..    Consider only given types of warnings

  -e CODE1,CODE2..    --exclude=CODE1,CODE2..    Exclude types of warnings

  -f FORMAT           --format=FORMAT            Output format (checkstyle, diff, gcc, json, json1, quiet, tty)
                      --list-optional            List checks disabled by default
                      --norc                     Don't look for .shellcheckrc files

  -o check1,check2..  --enable=check1,check2..   List of optional checks to enable (or 'all')

  -P SOURCEPATHS      --source-path=SOURCEPATHS  Specify path when looking for sourced files ("SCRIPTDIR" for script's dir)

  -s SHELLNAME        --shell=SHELLNAME          Specify dialect (sh, bash, dash, ksh)

  -S SEVERITY         --severity=SEVERITY        Minimum severity of errors to consider (error, warning, info, style)

  -V                  --version                  Print version information

  -W NUM              --wiki-link-count=NUM      The number of wiki links to show, when applicable

  -x                  --external-sources         Allow 'source' outside of FILES
                      --help                     Show this usage summary and exit


### Example Usage

* We can draw from some bash examples from [this github repo](https://github.com/alexanderepstein/Bash-Snippets).

So using the, "weather" script found here:

* [weather](https://github.com/alexanderepstein/Bash-Snippets/blob/master/weather/weather)

We can run this through shellcheck with:

```
shellcheck weather.sh
```

...and we get the following results:

```

In weather.sh line 90:
        git clone -q "https://github.com/$githubUserName/$repositoryName" && touch .BSnippetsHiddenFile || { echo "Failure!"; exit 1; } &
                                                                          ^-- SC2015: Note that A && B || C is not if-then-else. C may run when A is true.


In weather.sh line 170:
elif [[ "${@: -1}" == "m" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).


In weather.sh line 175:
elif [[ "${@: -1}" == "M" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).


In weather.sh line 180:
elif [[ "${@: -1}" == "mM" || "${@:-1}" == "Mm" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).
                              ^-------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).


In weather.sh line 185:
elif [[ "${@: -1}" == "iM" || "${@:-1}" == "Mi" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).
                              ^-------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).


In weather.sh line 190:
elif [[ "${@: -1}" == "i" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).

For more information:
  https://www.shellcheck.net/wiki/SC2199 -- Arrays implicitly concatenate in ...
  https://www.shellcheck.net/wiki/SC2015 -- Note that A && B || C is not if-t...
```

So the above basically gives different recommendations for the weather shell script.

#### Clearing the Error

We can attempt to clear an error:

```
In weather.sh line 170:
elif [[ "${@: -1}" == "m" ]];then
        ^--------^ SC2199: Arrays implicitly concatenate in [[ ]]. Use a loop (or explicit * instead of @).

```

By adding various shellcheck skip errors at the beginnings of the various statements, for example:

```
# shellcheck disable=SC2199
```

After running, "shellcheck weather.sh" - we recieve no errors or recommendations.