# https://hub.docker.com/layers/apache/airflow/slim-2.3.4-python3.10/images/sha256-cddd80bbaeecd9678030dfe0db67041f9c882afb7e1dc832a74560c2d18d5964?context=explore


ARG BASE_CONTAINER=apache/airflow:slim-2.3.4-python3.10

FROM $BASE_CONTAINER

LABEL author="Patrick Delaney"

# switch to user root briefly
USER root

# copy desired content from goapps file
COPY volumebindmount /home/volumebindmount

USER $NB_UID