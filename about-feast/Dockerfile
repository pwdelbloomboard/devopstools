# create an argument called, "BASE_CONTAINER" and point to jupyter minimal notebook image
ARG BASE_CONTAINER=jupyter/minimal-notebook:ubuntu-22.04

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney"

# switch to user root briefly
USER root

# Install dependencies
RUN conda install -y -c conda-forge mmh3

RUN pip install --upgrade pip &&            \
    pip install feast==0.24.0

RUN pip uninstall -y fsspec

RUN pip install aiobotocore=2.1.2
RUN pip install 'feast[aws]'==0.24.0

# Switch back to jovyan to avoid accidental container runs as root
USER $NB_UID