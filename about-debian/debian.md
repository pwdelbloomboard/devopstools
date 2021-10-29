# Exec'ing Directly into Debian Docker

One of the best ways to test out Debian is just to set up a Docker image, or even just exec straight into a Debian Docker Image:

Basically we could run a [Debian 11 image as follows](https://hub.docker.com/_/debian?tab=tags&page=1&name=bullseye-slim):

```
docker run -i -t debian:bullseye-slim /bin/bash
```
