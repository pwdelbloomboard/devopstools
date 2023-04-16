#!/bin/bash

if [ -f "cookies.txt" ]; then
    echo " -----> found cookies.txt, using with -cookie flag to connect. <-----"
    curl -i --cookie cookies.txt http://localhost:8889
else
    echo " -----> did not find cookies.txt, establishing connection with --cookie-jar. <-----"
    curl -i --cookie-jar cookies.txt http://localhost:8889
fi
