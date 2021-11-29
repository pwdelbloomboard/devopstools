# sed

* SED is a stream editor and can do lots of functions on files such as: searching, find and replace, insertion, deletion.
* This is commonly a substitute for find or replace.
* You can edit files without opening them.
* Much quicker than using VI/VIM/Nano, etc.

### Detailed Documentation

https://man7.org/linux/man-pages/man1/sed.1.html

### Simplified Form of Sed

The sed command in its simplest form can be used to replace a text in a file.

General syntax of using the sed command for replacement:

sed 's/<text_to_replace>/<replacement_text>/' <file_name>

### Common Options

#### sed i

-i[SUFFIX], --in-place[=SUFFIX]

* edit files in place (makes backup if SUFFIX supplied)

#### sed /!d/

* d      Delete pattern space.  Start next cycle.
* After the address (or address-range), and before the command, a ! may be inserted, which specifies that the command shall only be executed if the address (or address-range) does not match.


### Sed Options

Usage: sed [OPTION]... {script-only-if-no-other-script} [input-file]...

  -n, --quiet, --silent
                 suppress automatic printing of pattern space
      --debug
                 annotate program execution
  -e script, --expression=script
                 add the script to the commands to be executed
  -f script-file, --file=script-file
                 add the contents of script-file to the commands to be executed
  --follow-symlinks
                 follow symlinks when processing in place
  -i[SUFFIX], --in-place[=SUFFIX]
                 edit files in place (makes backup if SUFFIX supplied)
  -l N, --line-length=N
                 specify the desired line-wrap length for the `l' command
  --posix
                 disable all GNU extensions.
  -E, -r, --regexp-extended
                 use extended regular expressions in the script
                 (for portability use POSIX -E).
  -s, --separate
                 consider files as separate rather than as a single,
                 continuous long stream.
      --sandbox
                 operate in sandbox mode (disable e/r/w commands).
  -u, --unbuffered
                 load minimal amounts of data from the input files and flush
                 the output buffers more often
  -z, --null-data
                 separate lines by NUL characters
      --help     display this help and exit
      --version  output version information and exit

If no -e, --expression, -f, or --file option is given, then the first
non-option argument is taken as the sed script to interpret.  All
remaining arguments are names of input files; if no input files are
specified, then the standard input is read.