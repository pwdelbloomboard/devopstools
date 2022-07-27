#!/bin/bash

# [command] & [command]
# An ampersand does the same thing as a semicolon or newline in that it indicates the end of a command, 
# but it causes Bash to execute the command asynchronously. 
# That means Bash will run it in the background and run the next command immediately after, 
# without waiting for the former to end.

echo "hello world" & echo "hello asynchronously world"

# [command] && [command]


# [command] || [command]


# [command] | [command]


# [command] ; [command] [newline]



