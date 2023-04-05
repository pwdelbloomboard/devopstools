# syntax=docker/dockerfile:1.3-labs
FROM python:3.9-bullseye

SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8
ENV PIP_NO_CACHE_DIR=true
ENV PYTHONUNBUFFERED=1

RUN <<EOF
    adduser defaultuser
    echo 'PS1="\u@\h in \w \\$ "' >> /root/.bashrc
    echo 'PS1="\u@\h in \w \\$ "' >> /home/defaultuser/.bashrc
EOF

ARG ARCH=amd64

USER defaultuser

COPY --chown=defaultuser:defaultuser requirements.txt /home/launch/requirements.txt
RUN pip install -r /home/launch/requirements.txt

COPY --chown=defaultuser:defaultuser /model/model.joblib /home/app/model/model.joblib

COPY --chown=defaultuser:defaultuser /launch/start /launch/entrypoint /home/launch/
COPY --chown=defaultuser:defaultuser /app /home/app/

WORKDIR /home

ENTRYPOINT ["/home/launch/entrypoint"]
CMD ["/home/launch/start"]