#!/bin/bash

while getopts 'abc' OPTION; do
  case "$OPTION" in 
    a) 
      echo "Option a used" ;;

    b)
      echo "Option b used"
      ;;

    c)
      echo "Option c used"
      ;;

    ?) 
      echo "Usage: $(basename $0) [-a] [-b] [-c]"
      exit 1
      ;;
  esac
done
