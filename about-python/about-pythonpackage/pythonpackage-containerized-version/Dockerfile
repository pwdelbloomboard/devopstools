# create an argument called, "BASE_CONTAINER" and point to jupyter minimal notebook image
ARG BASE_CONTAINER=python:3.9-bullseye

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney"

# switch to user root briefly
USER root

# install dependencies
RUN apt update && apt install       \
    nano
    
RUN pip install                     \
    build==0.9.0

# Switch back to jovyan to avoid accidental container runs as root
USER $NB_UID