# python examples dockerfile with environments included

# setting up jupyter notebook
# create an argument called, "BASE_CONTAINER" and point to jupyter minimal notebook image
ARG BASE_CONTAINER=jupyter/minimal-notebook

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney"

# switch to user root briefly
USER root

# install depdencies
RUN pip install --no-cache-dir pandas==1.3.3

# copy data we want to evaluate and play around with:
COPY /datainclude /home/jovyan/work

# Switch back to jovyan to avoid accidental container runs as root
USER $NB_UID