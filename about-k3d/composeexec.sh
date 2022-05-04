#!/bin/bash
docker-compose down                                     && \
docker build -t k3dpurposed_debian:bullseye-slim .      && \
docker-compose up -d                                    && \
docker exec -t -i playwithk3d_container /bin/bash
