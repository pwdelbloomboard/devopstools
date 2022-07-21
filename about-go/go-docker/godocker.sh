#!/bin/bash

# use docker to get target image name
TARGETIMAGE=$(docker image ls | awk '{ print $1 }' | grep golang-bullseye)

echo "Pulling down container specified in docker-compose folder."
docker-compose down

# check if image exists already
if [[ $TARGETIMAGE == "golang-bullseye" ]]
then
    echo "Found image golang-bullseye. Skipping image update and starting container."
    docker-compose up -d                                                               && \
    docker exec -t -i playwithgolang_container /bin/bash
else
    echo "Image golang-bullseye not found, updating without downloading image."
    docker build -t golang-bullseye:latest .                                           && \
    docker-compose up -d                                                               && \
    docker exec -t -i playwithgolang_container /bin/bash
fi