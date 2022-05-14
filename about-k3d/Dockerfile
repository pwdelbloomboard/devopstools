# syntax=docker/dockerfile:1
FROM debian:bullseye-slim


# install dependencies necessary for installing k3d
# do both updates and upgrades to make up for any missing package repo's if this is a slim image
RUN apt-get -y update                                           && \
    # update first
    apt-get install -y --no-install-recommends                     \
        wget=1.21-1+deb11u1                                        \
        curl=7.74.0-1.3+deb11u1                                    \
        ca-certificates=20210119                                   \
        software-properties-common=0.96.20.2-2.1                   \
        nano=5.*                                                   \
        git=1:2.30.2-1
    # apt-get clean
    # don't remove source lists until not planning on installing anything more (freezing dockerfile)
    # rm -rf /var/lib/apt/lists/*

# install k3d
RUN wget -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
# install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

RUN chmod +x kubectl                                           && \
    mv kubectl /usr/local/bin

WORKDIR /home

RUN git clone https://github.com/vfarcic/k3d-demo.git

WORKDIR /home/k3d-demo/