$ gitgrep

You can search a git repo using, "gitgrep."

### Help File

NAME
       git-grep - Print lines matching a pattern

SYNOPSIS
       git grep [-a | --text] [-I] [--textconv] [-i | --ignore-case] [-w | --word-regexp]
                  [-v | --invert-match] [-h|-H] [--full-name]
                  [-E | --extended-regexp] [-G | --basic-regexp]
                  [-P | --perl-regexp]
                  [-F | --fixed-strings] [-n | --line-number] [--column]
                  [-l | --files-with-matches] [-L | --files-without-match]
                  [(-O | --open-files-in-pager) [<pager>]]
                  [-z | --null]
                  [ -o | --only-matching ] [-c | --count] [--all-match] [-q | --quiet]
                  [--max-depth <depth>] [--[no-]recursive]
                  [--color[=<when>] | --no-color]
                  [--break] [--heading] [-p | --show-function]
                  [-A <post-context>] [-B <pre-context>] [-C <context>]
                  [-W | --function-context]
                  [--threads <num>]
                  [-f <file>] [-e] <pattern>
                  [--and|--or|--not|(|)|-e <pattern>...]
                  [--recurse-submodules] [--parent-basename <basename>]
                  [ [--[no-]exclude-standard] [--cached | --no-index | --untracked] | <tree>...]
                  [--] [<pathspec>...]


DESCRIPTION
       Look for specified patterns in the tracked files in the work tree, blobs registered in the index file, or
       blobs in given tree objects. Patterns are lists of one or more search expressions separated by newline
       characters. An empty string as search expression matches all lines.

OPTIONS
       --cached
           Instead of searching tracked files in the working tree, search blobs registered in the index file.

       --no-index
           Search files in the current directory that is not managed by Git.

### Usage

Searching for the term, "context" across a folder structure, starting from the folder structure itself:

```
git grep -a context
```