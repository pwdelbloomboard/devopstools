#!/bin/bash

# [command] & [command]
# An ampersand does the same thing as a semicolon or newline in that it indicates the end of a command, 
# but it causes Bash to execute the command asynchronously. 
# That means Bash will run it in the background and run the next command immediately after, 
# without waiting for the former to end.

echo "hello world" & echo "hello asynchronously world"

# do a sleep to be able to finish the entire command chain before anything else gets started.
sleep 1

# [command] && [command]
# the second command is executed only if the first one succeeds (returns a zero exit status).

false && echo "This should never show up."

true && echo "First command executed as true, followed by &&."

# [command] || [command]
# The || represents a logical OR based upon the exit status of the sequential list.

true || echo "This should never show up."

echo "This statement shall execute." || true || echo "This should never show up."

echo "the result of true || echo [this should never show up.] was: $?"

(true || true) && echo "The result of the first || statement was true, so this statement executes."

(true || true)
echo "the result of (true || true) was: $?"

(false || true)
echo "the result of (false || true) was: $?"

(true || false)
echo "the result of (true || false) was: $?"

(false || false)
echo "the result of (false || false) was: $?"


# [command] | [command]


# [command] ; [command] [newline]



