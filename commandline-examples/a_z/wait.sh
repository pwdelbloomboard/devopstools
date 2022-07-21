#!/bin/bash

# creating simple process that will create file and write into it
cat > waitforit.txt <<< "Something is going to save into this file" &

# this will store the process id of the running process
# $! is a special variable in bash 
# that will hold the PID of the last active process i.e. creating a file.
pid=$!

# print process is running with its PID
echo "Process with PID $pid is running"

# Waiting until process is Completed
wait $pid

# print process id with its Exit status
# $? is special variable that holds the return value of the recently executed command.
echo "Process with PID $pid has finished with Exit status: $?"
