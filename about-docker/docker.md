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

## Docker Exit Codes

* Exit Code 0: Absence of an attached foreground process
* Exit Code 1: Indicates failure due to application error
* Exit Code 137: Indicates failure as container received SIGKILL (Manual intervention or ‘oom-killer’ [OUT-OF-MEMORY])
* Exit Code 139: Indicates failure as container received SIGSEGV
* Exit Code 143: Indicates failure as container received SIGTERM

# Resources

* [Docker Container Exit Codes](https://betterprogramming.pub/understanding-docker-container-exit-codes-5ee79a1d58f6)