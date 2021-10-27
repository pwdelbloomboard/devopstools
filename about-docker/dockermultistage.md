# Docker Multistage Builds

## Example

Let's say you have two Dockerfiles:

```
Dockerfile.build
FROM node:12.13.0-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
```
```
Dockerfile.main
FROM nginx
EXPOSE 3000
COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf
COPY /app/build /usr/share/nginx/html
```
Then you could build both of those files with:
```
Build.sh
#!/bin/sh
echo Building lukondefmwila/react:build

docker build -t lukondefmwila:build . -f Dockerfile.build

docker create --name extract lukondefmwila:build

docker cp extract:/app/build ./app

docker rm -f extract

echo Building lukondefmwila/react:latest

docker build --no-cache -t lukondefmwila/react:latest . -f Dockerfile.main
```
## Different Example

Let's say we have two docker files:

* base.Dockerfile
* local.Dockerfile

The base.Dockerfile contains everything necessary to build the, "main" things in common across many applications, whereas, local.Dockerfile has everything necessary to build for a local machine.

* base.Dockerfile comes first, and then ends within WORKDIR /app
* local.Dockerfile picks up right where base.Dockerfile left off, and uses WORKDIR /app to start off with.

```
#!/bin/sh

docker build -t [IMAGETAG1] .                                  \
-f base.Dockerfile                                             \
docker create --name extract [IMAGETAG1]                       \
docker cp extract:/app/build ./app                             \
docker rm -f extract                                           \
docker build --no-cache -t [IMAGETAG2] . -f local.Dockerfile   \
```

* IMAGETAG1 is just a placeholder tag and can be anything, as it will be deleted with the --no-cache flag option in the subsequent line.
* If you're going to push to a registry, IMAGETAG2 needs to be the tag used to push to that registry, it's the, "final form."


Finalizing, with appropriatetags:

```
#!/bin/sh

docker build -t testimagetag:build . -f base.Dockerfile        \
docker create --name extract testimagetag:build                \
docker cp extract:/app/build ./app                             \
docker rm -f extract                                           \
docker build --no-cache -t https://registry.gitlab.com/bloomboard/peryton/styleguide/base:master . -f local.Dockerfile
```




# Resources

[Docker Multistage Builds](https://earthly.dev/blog/docker-multistage/)