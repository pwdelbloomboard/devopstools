#!/bin/bash

read -d '' PAYLOAD << EOF
[
    {"Age": 85, "Sex": "male", "Embarked": "S"},
    {"Age": 24, "Sex": "female", "Embarked": "C"},
    {"Age": 3, "Sex": "male", "Embarked": "C"},
    {"Age": 21, "Sex": "male", "Embarked": "S"}
]
EOF

echo "$PAYLOAD"

echo "Attempting curl -X POST on /api/query with input data $PAYLOAD"
curl -X POST http://localhost:8889/api/query                            \
    -H 'Content-Type: application/json'                                 \
    -d "$PAYLOAD"                                                       \
    --output ./app/output/prediction_output.json