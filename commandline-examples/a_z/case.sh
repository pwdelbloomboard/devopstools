#!/bin/bash

echo -n "Enter the name of a country: "
read COUNTRY

echo -n "The official language of $COUNTRY is "

case $COUNTRY in

  Lithuania)
    echo -n -e "Lithuanian \n"
    ;;

  Romania | Moldova)
    echo -n -e "Romanian \n"
    ;;

  Italy | "San Marino" | Switzerland | "Vatican City")
    echo -n -e "Italian \n"
    ;;

  *)
    echo -n -e "unknown \n"
    ;;
esac
