#!/bin/bash
# stop testcluster if it exists
[ "$(k3d cluster list | grep testcluster)" ]                                       && \
k3d cluster stop testcluster                                                       && \
k3d cluster delete testcluster

# stop overall container
docker-compose down

