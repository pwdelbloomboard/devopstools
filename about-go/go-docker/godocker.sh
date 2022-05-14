#!/bin/bash

# spin up and start new container which starts cluster
docker-compose down                                                                && \
docker build -t golang-bullseye:latest .                                           && \
docker-compose up -d                                                               && \
docker exec -t -i playwithgolang_container /bin/bash
