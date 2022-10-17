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

WORKDIR /root

ENV NB_PREFIX /
ENTRYPOINT ["tini", "--"]

CMD ["sh","-c", "jupyter lab --notebook-dir=/home/${NB_USER} --ip=0.0.0.0 --no-browser --allow-root --port=8888 --NotebookApp.token='' --NotebookApp.password='' --NotebookApp.allow_origin='*' --NotebookApp.base_url=${NB_PREFIX}"]


# Switch back to jovyan to avoid accidental container runs as root
USER $NB_UID