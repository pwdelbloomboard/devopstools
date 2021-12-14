#!/bin/bash

#We will take standard input
#Will list all files at the path
#We will concate variable and string

echo "Enter the path"
read path

echo "How deep in directory you want to go:"
read depth

echo "All files at path " $path

du -d $depth -all -h $path