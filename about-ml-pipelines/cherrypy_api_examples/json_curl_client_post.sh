#!/bin/bash

# attempt POST
echo "attempting curl command and echo result: "
curl --silent --show-error                      \
    -X POST -H "Content-Type: application/json" \
    -d '{"input": "hi"}'                        \
    http://127.0.0.1:8080/; echo
