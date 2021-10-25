# About Docker

## About Docker Images

* [Explaining Docker Image IDs](https://windsock.io/explaining-docker-image-ids/)

The definition of an image:

> A docker image is a read-only template for creating containers, and provides a filesystem based on an ordered union of multiple layers of files and directories, which can be shared with other images and containers. 

> Sharing of image layers is a fundamental component of the Docker platform, and is possible through the implementation of a copy-on-write (COW) mechanism. During its lifetime, if a container needs to change a file from the read-only image that provides its filesystem, it copies the file up to its own private read-write layer before making the change.

> A layer or 'diff' is created during the Docker image build process, and results when commands are run in a container, which produce new or modified files and directories.

So basically, Docker images are fundamentally sets of layers.

Previously, the way to inspect layers and look at their id was to do:

```
docker inspect IMAGENAME
```
Which would pull up an image id, a full image id, basically a long string which would serve as an identity. However, this structure could result in tampering, in that the image and 

> Since Docker v1.10, generally, images and layers are no longer synonymous. Instead, an image directly references one or more layers that eventually contribute to a derived container's filesystem.

> A Docker image now consists of a configuration object, which (amongst other things) contains an ordered list of layer digests, which enables the Docker Engine to assemble a container's filesystem with reference to layer digests rather than parent images. The image ID is also a digest, and is a computed SHA256 hash of the image configuration object, which contains the digests of the layers that contribute to the image's filesystem definition.

```
 docker inspect ghcr.io/pwdelbloomboard/ps-container
[
    {
        "Id": "sha256:23f9401d2b7d275acf143295324d8f0bf3870988c3ea8432ce1db6b0e11bd36a",
        "RepoTags": [
            "ghcr.io/pwdelbloomboard/ps-container:latest"
        ],
        "RepoDigests": [
            "ghcr.io/pwdelbloomboard/ps-container@sha256:8f9959b48435e7689fa8a091fac9140766e061510daa4ed0db263c05ca4baf20"
        ],
        "Parent": "",
        "Comment": "buildkit.dockerfile.v0",
        "Created": "2021-10-02T01:44:37.9352531Z",
        "Container": "",
        "ContainerConfig": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": null,
            "Cmd": null,
            "Image": "",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": null
        },
        "DockerVersion": "",
        "Author": "",
        "Config": {
            "Hostname": "",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "Tty": false,
            "OpenStdin": false,
            "StdinOnce": false,
            "Env": [
                "PATH=/app/node_modules/.bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                "NODE_VERSION=12.22.6",
                "YARN_VERSION=1.22.5"
            ],
            "Cmd": [
                "yarn",
                "start"
            ],
            "ArgsEscaped": true,
            "Image": "",
            "Volumes": null,
            "WorkingDir": "/",
            "Entrypoint": [
                "docker-entrypoint.sh"
            ],
            "OnBuild": null,
            "Labels": null
        },
        "Architecture": "amd64",
        "Os": "linux",
        "Size": 418081028,
        "VirtualSize": 418081028,
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/wujg3c4g9c77ojuo4d3x4rry7/diff:/var/lib/docker/overlay2/p5yfc2dc30bl8j9ad9fomyeli/diff:/var/lib/docker/overlay2/qtascl2rge6se4astamhnqi78/diff:/var/lib/docker/overlay2/e957619865e0bcee4bd3710bed2bafec9d4d58294659e959c5d39d49e5977446/diff:/var/lib/docker/overlay2/6ee20854b8218a071704a7e797e26ede30d49a35e5e5439f92ead55ae01a02de/diff:/var/lib/docker/overlay2/98505a339b83011fcbb3519f6fad06ed2fca0d93fb6c08c4365abbe25ca26947/diff:/var/lib/docker/overlay2/dd357e7b3053fcd8e1c0a438612c279602381ad9972979eec7dba617ba30a655/diff",
                "MergedDir": "/var/lib/docker/overlay2/eqfft42g2wwn88kon2d8bdq9x/merged",
                "UpperDir": "/var/lib/docker/overlay2/eqfft42g2wwn88kon2d8bdq9x/diff",
                "WorkDir": "/var/lib/docker/overlay2/eqfft42g2wwn88kon2d8bdq9x/work"
            },
            "Name": "overlay2"
        },
        "RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:39982b2a789afc156fff00c707d0ff1c6ab4af8f1666a8df4787714059ce24e7",
                "sha256:f8700d3a252fffe60e30bc672e8a6560f30a3ce8816f2ad396020553fe4d9210",
                "sha256:b8f0e895f5208b04d533d013ddec6f12642fdd679ef70bc1497ffe733c97428b",
                "sha256:446ec7c50f08cfba388bcebe29f54b2a46a5ddccdabd6b4caac21cbdb7c60b4b",
                "sha256:62595f19d44040cd7b2866304305cc5dcf70c98cafab5d72af748242f88443f4",
                "sha256:4848bc58006830a74703e36d113d2bacc6dd8a2be7fc4b6b2ea289abdc758ba6",
                "sha256:32fc229d01775980019d4261c3536de6c9e15c296db5d861588eb681efd25ac7",
                "sha256:3e796ce6091bde0c0545195c2d46fc20225bf8578e527a3cdcf1490f9b6c8729"
            ]
        },
        "Metadata": {
            "LastTagTime": "2021-10-04T23:56:16.919268Z"
        }
    }
]

```

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
* [Explaining Docker Image IDs](https://windsock.io/explaining-docker-image-ids/)