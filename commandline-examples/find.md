### basic find

For finding a script...

```
find ~ -name script.sh 
```

### Longer Documentation on Find

https://man7.org/linux/man-pages/man1/find.1.html

### find on the filesystem

```
find / -name script.sh
```

This will show the pathname on the filesystem.

If you want to avoid errors, do:

```
find / -name script.sh 2>/dev/null
```

### find by type

```
find -type f
```

This comes from, "-type [bcdpflsD]"

  -type c

File is of type c:

* b      block (buffered) special
* c      character (unbuffered) special
* d      directory
* p      named pipe (FIFO)
* f      regular file
* l      symbolic link; this is never true if the -L option
                     or the -follow option is in effect, unless the
                     symbolic link is broken.  If you want to search for
                     symbolic links when -L is in effect, use -xtype.

* s      socket
* D      door (Solaris)

* To search for more than one type at once, you can supply the combined list of type letters separated by a comma `,' (GNU extension).

### printf Option



```
 -printf format
              True; print format on the standard output, interpreting
              `\' escapes and `%' directives.  Field widths and
              precisions can be specified as with the printf(3) C
              function.  Please note that many of the fields are printed
              as %s rather than %d, and this may mean that flags don't
              work as you might expect.  This also means that the `-'
              flag does work (it forces fields to be left-aligned).
              Unlike -print, -printf does not add a newline at the end
              of the string.  The escapes and directives are:

              \a     Alarm bell.

              \b     Backspace.

              \c     Stop printing from this format immediately and
                     flush the output.

              \f     Form feed.

              \n     Newline.

              \r     Carriage return.

              \t     Horizontal tab.

              \v     Vertical tab.

              \0     ASCII NUL.

              \\     A literal backslash (`\').

...



```

This can format based upon date and many other options.

### Find exec Option

```
-exec command ;

       Execute  command;  true  if 0 status is returned.  All following arguments to find are taken to be
       arguments to the command until an argument consisting of `;' is encountered.  The string  `{}'  is
       replaced  by  the  current  file name being processed everywhere it occurs in the arguments to the
       command, not just in arguments where it is alone, as in some versions  of  find.   Both  of  these
       constructions  might  need  to be escaped (with a `\') or quoted to protect them from expansion by
       the shell.  See the EXAMPLES section for examples of the use of the -exec option.   The  specified
       command  is  run  once  for each matched file.  The command is executed in the starting directory.
       There are unavoidable security problems surrounding use of the -exec action; you  should  use  the
       -execdir option instead.


-exec command {} +
       This variant of the -exec action runs the specified command on the selected files, but the command
       line is built by appending each selected file name at the end; the total number of invocations  of
       the command will be much less than the number of matched files.  The command line is built in much
       the same way that xargs builds its command lines.  Only one instance of `{}' is allowed within the
       command.  The command is executed in the starting directory.
```


### Find Expanded Help

# find --help

Usage: find [-H] [-L] [-P] [-Olevel] [-D debugopts] [path...] [expression]

default path is the current directory; default expression is -print
expression may consist of: operators, options, tests, and actions:
operators (decreasing precedence; -and is implicit where no others are given):

      ( EXPR )   ! EXPR   -not EXPR   EXPR1 -a EXPR2   EXPR1 -and EXPR2
      EXPR1 -o EXPR2   EXPR1 -or EXPR2   EXPR1 , EXPR2

positional options (always true): -daystart -follow -regextype

normal options (always true, specified before other expressions):
      -depth --help -maxdepth LEVELS -mindepth LEVELS -mount -noleaf
      --version -xdev -ignore_readdir_race -noignore_readdir_race
tests (N can be +N or -N or N): -amin N -anewer FILE -atime N 

-cmin N
      -cnewer FILE -ctime N -empty -false -fstype TYPE -gid N -group NAME
      -ilname PATTERN -iname PATTERN -inum N -iwholename PATTERN -iregex PATTERN
      -links N -lname PATTERN -mmin N -mtime N -name PATTERN -newer FILE
      -nouser -nogroup -path PATTERN -perm [-/]MODE -regex PATTERN
      -readable -writable -executable
      -wholename PATTERN -size N[bcwkMG] -true -type [bcdpflsD] -uid N
      -used N -user NAME -xtype [bcdpfls]      -context CONTEXT

actions: -delete -print0 -printf FORMAT -fprintf FILE FORMAT -print
      -fprint0 FILE -fprint FILE -ls -fls FILE -prune -quit
      -exec COMMAND ; -exec COMMAND {} + -ok COMMAND ;
      -execdir COMMAND ; -execdir COMMAND {} + -okdir COMMAND ;

Valid arguments for -D:
exec, opt, rates, search, stat, time, tree, all, help
Use '-D help' for a description of the options, or see find(1)

Please see also the documentation at http://www.gnu.org/software/findutils/.

You can report (and track progress on fixing) bugs in the "find"
program via the GNU findutils bug-reporting page at
https://savannah.gnu.org/bugs/?group=findutils or, if
you have no web access, by sending email to <bug-findutils@gnu.org>.