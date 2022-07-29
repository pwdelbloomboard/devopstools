#!/bin/bash

# redirection operators

# redirection of output
# causes the redirection.txt to be opened and stdout (1) to be output to the file

echo "hello world" > redirection.txt

cat redirection.txt

sleep 3

# using >| sets "noclobber" to false, which means the file will be overwritten without error
# noclobber can also be set with "set -o"

set -o noclobber

echo "hello world pipe" >| redirection.txt

echo "The stderr of the last command was: $?"

cat redirection.txt

sleep 1

echo "hello world again" > redirection.txt

echo "The stderr of the last command was: $?"

cat redirection.txt

sleep 1

# redirection of input
# the file redirection.txt to be opened and placed in the stdin (0) input to the command

cat<redirection.txt

sleep 1

# << here-document structure
# read input from the source until a line contiaing only the key/delimiter is seen

echo "the following will be a json item: "

cat << EoF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:*"
      ],
      "Resource": [
        "arn:aws:s3:::*"
      ]
    }
  ]
}
EoF

# <<< here-string

# standard output redirector
# &> is preferred over >& and both are equivalent to word 2>&1
# file descriptor 1 is the standard output, stdout
# file descriptor 2 is the standard error, stderr
# whereas 2>1 might look like, "put stdout on stderr" it actually means, "put stdout on file called 1"
# & indicates the following is a file descriptor, not a filename, so you use 2>&1
# &> is a shortened format

./errorsandnoterrors > errorsnoerrorsresults.txt 2>&1

echo "the results of errorsnoerrorsresults.txt redirection using 2>&1 "
cat errorsnoerrorsresults.txt