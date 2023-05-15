
# syntax = docker/dockerfile:1.3-labs
ARG BASE_CONTAINER=python:3.9-bullseye

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney"

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY volumebindmount /home/volumebindmount

# copy so that we can use the same requirements file in tox
COPY requirements.txt /home/volumebindmount/requirements.txt 

ENTRYPOINT ["/bin/bash"]