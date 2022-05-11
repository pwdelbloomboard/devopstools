#!/bin/bash
# stop testcluster if it exists
[ "$(k3d cluster list | grep testcluster)" ]                                       && \
k3d cluster stop testcluster                                                       && \
k3d cluster delete testcluster

# spin up and start new container which starts cluster
docker-compose down                                                                && \
docker build -t kustomize_purposed_debian:bullseye-slim .                          && \
docker-compose up -d                                                               && \
docker exec -t -i playwithkustomize_container /bin/bash
