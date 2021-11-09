# printf

## Printf Format Specifiers

%c	character

%d	decimal (integer) number (base 10)

%e	exponential floating-point number

%f	floating-point number

%i	integer (base 10)

%o	octal number (base 8)

%s	a string of characters

%u	unsigned decimal (integer) number

%x	number in hexadecimal (base 16)

%%	print a percent sign

\%	print a percent sign

### Printf Special Characters

\a	audible alert

\b	backspace

\f	form feed

\n	newline, or linefeed

\r	carriage return

\t	tab

\v	vertical tab

\\	backslash



# command help 

printf: printf [-v var] format [arguments]
    Formats and prints ARGUMENTS under control of the FORMAT.
    
    Options:
      -v var	assign the output to shell variable VAR rather than
    		display it on the standard output
    
    FORMAT is a character string which contains three types of objects: plain
    characters, which are simply copied to standard output; character escape
    sequences, which are converted and copied to the standard output; and
    format specifications, each of which causes printing of the next successive
    argument.
    
    In addition to the standard format specifications described in printf(1),
    printf interprets:
    
      %b	expand backslash escape sequences in the corresponding argument
      %q	quote the argument in a way that can be reused as shell input
      %(fmt)T	output the date-time string resulting from using FMT as a format
    	        string for strftime(3)
    
    The format is re-used as necessary to consume all of the arguments.  If
    there are fewer arguments than the format requires,  extra format
    specifications behave as if a zero value or null string, as appropriate,
    had been supplied.
    
    Exit Status:
    Returns success unless an invalid option is given or a write or assignment
    error occurs.
