## Useful Docker Commands


* [Remove a volume.](https://docs.docker.com/engine/reference/commandline/volume_rm/)

```
docker volume rm [OPTIONS]
```


* [Remove all Images](https://docs.docker.com/engine/reference/commandline/image_prune/)

```
docker image prune -a
```

* [List all Docker Images]

```
docker image ls -a

docker image ls --all

```

Why don't all images show up on terminal that show up in Docker dashboard?

This might be because docker desktop needs a while to process what has been deleted and what has not.

## Setting Docker Virtual Memory Space

https://docs.docker.com/desktop/mac/space/