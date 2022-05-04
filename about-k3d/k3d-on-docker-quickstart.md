# Quickstart for Getting Dockerfile Going with docker-compose.yml

0. Build the image in the same folder as the Dockerfile with:

```
docker build -t k3dpurposed_debian:bullseye-slim .
```

1. Navigate to folder containing Dockerfile (root of about-hugo)
2. The following command will get the container up and running:

```
docker-compose up -d
```

3. View the site at localhost:1313

4. The following command will allow you to exec into the container:

```
docker exec -t -i playwithk3d_container /bin/bash
```

Or better yet, to both run and exec in at once, do:

```
./composeexec.sh
```