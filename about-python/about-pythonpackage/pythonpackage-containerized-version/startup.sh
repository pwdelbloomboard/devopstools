#!/bin/bash

echo "Building image and tagging with name: "
docker build -t debian_pythonpackage_experiment_image .

docker-compose run debian_pythonpackage_experiment_service /bin/bash