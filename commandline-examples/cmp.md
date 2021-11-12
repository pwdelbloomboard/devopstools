 cmp(1) - compare two files byte by byte
cmp [OPTION]... FILE1 [FILE2 [SKIP1 [SKIP2]]]
The  optional  SKIP1 and SKIP2 specify the number of bytes to skip at the beginning of each file (zero by
default).
-b, --print-bytes
       print differing bytes
-i, --ignore-initial=SKIP
       skip first SKIP bytes of both inputs
-i, --ignore-initial=SKIP1:SKIP2
       skip first SKIP1 bytes of FILE1 and first SKIP2 bytes of FILE2
-l, --verbose
       output byte numbers and differing byte values
-n, --bytes=LIMIT
       compare at most LIMIT bytes
-s, --quiet, --silent
       suppress all normal output
--help display this help and exit
-v, --version
       output version information and exit
If  a  FILE  is  `-'  or  missing,  read  standard  input.  Exit status is 0 if inputs are the same, 1 if
different, 2 if trouble.