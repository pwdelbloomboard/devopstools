# create an argument called, "BASE_CONTAINER" and point to jupyter minimal notebook image
ARG BASE_CONTAINER=jupyter/scipy-notebook:python-3.8

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney - Python 3.8 Jupyter Notebook"

# switch to user root briefly
USER root

# install depdencies
RUN apt update && apt install       \
    nano

COPY requirements.txt ./

# Switch back to jovyan to avoid accidental container runs as root
USER $NB_UID

