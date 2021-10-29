# Docker Bind Mounts

### Background

All files created in a container are stored on a writable container layer, which means that:

1. Data does not persist.
2. The data is coupled to the host machine, you can't move it around very easily.
3. Writing to a container's writable layer requires a storage driver to manage the filesystem.

There are two options to store files on the host machine, so that the files are persisted even after a container stops:

![](/img/typesofmounts.png)

#### Volumes

* [volumes](/about-docker/docker-volumes.md)

* Volumes sit on a specific part of the filesystem, (disk) which is designated by the container, let's call it the, "Docker Area."
* Volumes are created and managed by Docker itself, e.g. you can do, "docker create volume" as a command...docker manages the isolation away from the core functionality of the host machine.
* Volumes can be shared between different containers simultaneously. This management affords flexibility in this manner.
* Volumes *must* be named or be designated as, "annonymous" which are not given an explicit name, but are given a random name, much like the random container naming system.
* These do persist on a disk.

##### Use Cases for Volumes

* Sharing data amoung containers.
* When Docker is not guaranteed to have a given directory or file structure...e.g. different types of machines like Windows vs. Linux/Unix. Volumes decouple the configuration of the Docker host from the container runtime.
* When you want to store a container's data on a remote host or cloud provider rather than a local machine.
* When you need to backup, restore or migrate data from one Docker host to another.
* Performance, or fully native file system behavior on Docker Desktop (for dev/prod compatibility)

#### Bind Mounts

* [bind mounts](/about-docker/docker-bindmounts.md)

* Quite simply, bind mounts sit within the filesystem, not on the pre-designated, "Docker Area," but rather on a location specified by the user.
* You can't use Docker CLI commands to explicitly to manage bind mounts.

##### Use Casess for Bind Mounts

* Sharing configuration files from host machine to containers
* Sharing source code or build artifacts between dev environment on a Docker host and a container. So for example, you could run a NodeJS project on a local machine in a particular folder, and as artifacts get created, build artifacts get created and stored in that folder, which the --mount bind mount points to, granting access to those build artifacts to the container itself.
* When the file or directory structure of the 

#### Other Mount Types

##### tmpfs Mount

This is a completely seperate type of mount which is stored right in the system memory (RAM), not on the file system (DISK).

* These are not persisted on the disk.
* tmpfs mounts are best used for cases when you do not want the data to persist either on the host machine or within the container. This may be for security reasons or to protect the performance of the container when your application needs to write a large volume of non-persistent state data.


##### named pipes

named pipes: An npipe mount can be used for communication between the Docker host and a container. Common use case is to run a third-party tool inside of a container and connect to the Docker Engine API using a named pipe.


### About Bind Mounts

* [Docker Bind Mounts](https://docs.docker.com/storage/bind-mounts/)

> Bind mounts have been around since the early days of Docker. Bind mounts have limited functionality compared to volumes. When you use a bind mount, a file or directory on the host machine is mounted into a container. The file or directory is referenced by its absolute path on the host machine. By contrast, when you use a volume, a new directory is created within Docker’s storage directory on the host machine, and Docker manages that directory’s contents.

![](/img/typesofmounts.png)

> Choose the -v or --mount flag

> Difference is that the -v syntax combines all the options together in one field, while the --mount syntax separates them. Here is a comparison of the syntax for each flag.

### Building a Basic Bind Mount with and Empty Directory

> Use the following command to bind-mount the target/ directory into your container at /app/. Run the command from within the source directory. The $(pwd) sub-command expands to the current working directory on Linux or macOS hosts.

```
docker run -d \
  -it \
  --name devtest \
  --mount type=bind,source="$(pwd)"/target,target=/app \
  nginx:latest
```

* In order for the above to work, we have to run the above code from the /project directory within terminal, and have a folder structure which looks like the following:

```
project
│   README.md
│   Dockerfile
│
└───target
    │
    └───app
```

* The above command, if executed from, the "project" folder, will create a container called, "devtest" which is, "interactive / tty" (-it), uses nginx and has a bind mount in the, "app" directory.

If we run "docker inspect devtest" we get the following areas of json output which show what our Mounts are:

```
docker inspect devtest | jq '.[] | .HostConfig' | jq '.' | jq '.Mounts'
[
  {
    "Type": "bind",
    "Source": "/Users/patrick.delaneybloomboard.com/Projects/dockerbindmount/target",
    "Target": "/app"
  }
]
```

```
$ docker inspect devtest | jq '.[] | .Mounts'
[
  {
    "Type": "bind",
    "Source": "/Users/patrick.delaneybloomboard.com/Projects/dockerbindmount/target",
    "Destination": "/app",
    "Mode": "",
    "RW": true,
    "Propagation": "rprivate"
  }
]
```
These above show that the mount is a bind mount, with the source and destinations shown, shows that it is read-write and that the propagation is set to rprivate.

So this is literally a container running and then having access directly to a part of the filesystem on the host machine which it can interact with.

### Building a Basic Bind Mount in a Docker File

*** ---> Unknown / Unsure Here <---

Therefore, to transcribe this over to a Dockerfile, we can use experimental syntax and the Docker Buildpack.

```
RUN --mount=type=bind,source="$(pwd)"/target,target=/app
```

* Once this has been added, the application can be built. [Here](https://github.com/pwdelbloomboard/dockerbindmount) is a sample application which uses the Docker Bind Mount within the Dockerfile build file as a demonstration.

* After the image has been built, it can be run per the instructions and options desired (listed on [this projects](https://github.com/pwdelbloomboard/dockerbindmount) README.md file).



After this point, we could optionally set up directories, change user owners, launch programs, etc., within the container itself.



### Building a Basic Bind Mount with a non-Empty Directory





### Types of Bind Mounts

shared - 

slave - 

private - 

rshared -

rslave - 

rprivate - 