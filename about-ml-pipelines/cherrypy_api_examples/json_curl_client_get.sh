#!/bin/bash

if [ -f "cookies.txt" ]; then
    echo " -----> found cookies.txt, using with -cookie flag to connect. <-----"
    # attempt GET
    curl -i --cookie cookies.txt --silent --show-error                      \
        -X GET                                                              \
        -H "Content-Type: application/json"                                 \
        http://127.0.0.1:8080/; echo
else
    echo " -----> did not find cookies.txt, establishing connection with --cookie-jar. <-----"
    # attempt POST
    echo "attempting curl command and echo result: "
    curl -i --cookie-jar cookies.txt --silent --show-error                  \
        -X POST -H "Content-Type: application/json"                         \
        -d '{"input": "hi"}'                                                \
        http://127.0.0.1:8080/; echo
        
    # attempt GET
    curl -i --cookie cookies.txt --silent --show-error                      \
        -X GET                                                              \
        -H "Content-Type: application/json"                                 \
        http://127.0.0.1:8080/; echo
fi