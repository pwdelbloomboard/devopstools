#!/bin/bash

# attempt POST
curl --silent --show-error                      \
    -X POST -H "Content-Type: application/json" \
    -d '{}'                                     \
    http://127.0.0.1:8080/; echo
