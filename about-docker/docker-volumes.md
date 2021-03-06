# Docker Volumes
## About Docker Volumes

[About Volumes in the Context of Docker](https://docs.docker.com/storage/volumes/)

> Volumes are the preferred mechanism for persisting data generated by and used by Docker containers. While bind mounts are dependent on the directory structure and OS of the host machine, volumes are completely managed by Docker. Volumes have several advantages over bind mounts:

![](/img/typesofmounts.png)

## More Background

For more background go to [docker bindmounts](/about-docker/docker-bindmounts.md)

## Setting Up a Docker Volume

* [Using Volumes Documentation](https://docs.docker.com/storage/volumes/)

### -v vs --mount flag

* --mount is more verbose
* -v combines all the options together in one field, whereas --mount seperates them, so you have to be explicit

#### -v

> -v or --volume: Consists of three fields, separated by colon characters (:). 

* The fields must be in the correct order, and the meaning of each field is not immediately obvious.
* In the case of named volumes, the first field is the name of the volume, and is unique on a given host machine. For anonymous volumes, the first field is omitted.
* The second field is the path where the file or directory are mounted in the container.
* The third field is optional, and is a comma-separated list of options, such as ro. These options are discussed below.

In Summary:

```
-v [volumename]:[mount_location],[options]
```

Example:

```
-v "$PWD/somedir":/somedir
```
* This is attaching the volume at $PWD/somedir to the mount_location on the container at /somedir
#### --mount

> -mount: Consists of multiple key-value pairs, separated by commas and each consisting of a <key>=<value> tuple. The --mount syntax is more verbose than -v or --volume, but the order of the keys is not significant, and the value of the flag is easier to understand.

* The type of the mount, which can be bind, volume, or tmpfs. This topic discusses volumes, so the type is always volume.
* The source of the mount. For named volumes, this is the name of the volume. For anonymous volumes, this field is omitted. May be specified as source or src.
* The destination takes as its value the path where the file or directory is mounted in the container. May be specified as destination, dst, or target.
* The readonly option, if present, causes the bind mount to be mounted into the container as read-only. May be specified as readonly or ro.
* The volume-opt option, which can be specified more than once, takes a key-value pair consisting of the option name and its value.

```
--mount 'type=volume,src=<VOLUME-NAME>,dst=<CONTAINER-PATH>,volume-driver=local,volume-opt=type=nfs,volume-opt=device=<nfs-server>:<nfs-path>,"volume-opt=o=addr=<nfs-address>,vers=4,soft,timeo=180,bg,tcp,rw"'

```

### Creating a Volume

Volumes can be managed within the scope of docker, unlike a [bind mount](/about-docker/docker-bindmounts.md).

```
docker volume create my-vol
```

### Inspecting a Volume

You can list volumes by doing:

```
docker volume ls
```
And inspect the volume in json format with:

```
docker inspect my-vol

{
  "CreatedAt": "2021-10-28T16:19:00Z",
  "Driver": "local",
  "Labels": {},
  "Mountpoint": "/var/lib/docker/volumes/my-vol/_data",
  "Name": "my-vol",
  "Options": {},
  "Scope": "local"
}

docker inspect my-vol | jq '.[0].Name'             

"my-vol"
```

### Removing a Volume

```
docker volume rm my-vol
```

### Start Container with Volume

```
docker run -d \
  --name devtest \
  --mount source=myvol2,target=/app \
  nginx:latest
```

* The above starts a container called, "devtest" which runs nginx, and attaches a volume, "myvol2" to /app.  
* If we inspect this we get:

```
docker inspect devtest | jq '.[0].Mounts'
[
  {
    "Type": "volume",
    "Name": "myvol2",
    "Source": "/var/lib/docker/volumes/myvol2/_data",
    "Destination": "/app",
    "Driver": "local",
    "Mode": "z",
    "RW": true,
    "Propagation": ""
  }
]

```
Or, converting to yaml:

```
docker inspect devtest | jq '.[0].Mounts' | yq eval -P
- Type: volume
  Name: myvol2
  Source: /var/lib/docker/volumes/myvol2/_data
  Destination: /app
  Driver: local
  Mode: z
  RW: true
  Propagation: ""
```
We see above that on, "devtest", the "Mount" is myvol2, with a source within the /var/lib/docker/volumes area of our machine, managed by the Docker Runtime.

We can stop the container, remove it, and then remove the volume all as seperate commands, the volume exists independently of the container.

```
docker container stop devtest

docker container rm devtest

docker volume rm myvol2
```

### Pre-Populating a Volume with a Container

* Certain containers (meaning, pre built images running as an app container), contain default content, such as in the example of nginx, which stores default HTML content in /usr/share/nginx/html.
* Sometimes these sets of directories, such as the default HTML content in nginx, need to be mounted. In the above example, we mounted the directory, "/app" to an nginx container. We may want to mount, "usr/share/nginx/html" as a volume as well, so that other containers which happen to be running also have access to this directory as a volume, rather than it being, "trapped," on the VM disk.
* So basically upon startup, we copy the directory's content into a volume, and then mount the volume!

```
docker run -d \
  --name=nginxtest \
  --mount source=nginx-vol,destination=/usr/share/nginx/html \
  nginx:latest
```

*  --mount mount ...Attach a filesystem mount to the container
* --mount source={volume_name},destination={container_default_content_directory}

The, "--mount" flat is a bit confusing here, because the, "source" is actually the location where we are putting the content, e.g. we're copying the content to the volume, or the nginx-vol in this case. The, "destination" is where we are getting the content from on the actual container, the original directory which stores that default content.

From the Docker documentation:

> this example starts an nginx container and populates the new volume nginx-vol with the contents of the container???s /usr/share/nginx/html directory, which is where Nginx stores its default HTML content.

### Read-Only

We can create read-only volumes, which allow multiple containers to access certain static information, perhaps a key, by adding a readonly flag (using the above example).

```
docker run -d \
  --name=nginxtest \
  --mount source=nginx-vol,destination=/usr/share/nginx/html,readonly \
  nginx:latest
```
Inspecting, "Mounts" you will see that the, "RW" mode is set to, "false."

### Fault Tolerant Applications

You can use Volume Drivers to attach the same shared file storage to multiple replicas of the same VM using the following documentation:

https://docs.docker.com/storage/volumes/#share-data-among-machines

### Backup, Restore or Migrate Data Volumes

https://docs.docker.com/storage/volumes/#backup-restore-or-migrate-data-volumes

Volumes are useful for backups, restores and migrations.
#### Backup a Container

* You can create a new container called, "dbstore," and mount a volume, "/dbdata", using the ubuntu image and running bash:

```
docker run -v /dbdata --name dbstore ubuntu /bin/bash
```
Then, to effectively copy and paste this, run:

```
docker run --rm --volumes-from dbstore          \
    -v $(pwd):/backup ubuntu                    \
    tar cvf /backup/backup.tar /dbdata
```
Basically waht this does is launches a new container, 

* Use the --volumes-from flag to create a new container that mounts that volume, "from that container" (or from dbstore in this instance).
* Mounts a local host directory as /backup with "-v $(pwd):/backup"
* Pass a command that tars the contents of the dbdata volume to a backup.tar file inside our /backup directory.


#### Restore a Container from Backup

https://docs.docker.com/storage/volumes/#restore-container-from-backup

#### Attaching a Volume to an Existing Container

There is no way to attach a volume to a currently running container, it can't be, "hot added," however you can copy the exact state of a running container, re-run it and attach the volume in a new container, which is effectively the same thing as, "hot connecting," because you're just duplicating an exact state.

For that, you use, [docker commit](https://docs.docker.com/engine/reference/commandline/commit/)

```
docker commit [CONTAINER ID] newimagename
```
Example:

```
docker commit 5a8f89adeead newimagename

docker run -ti                  \
    -v "$PWD/somedir":/somedir  \ 
    newimagename /bin/bash

```

* So basically, we're attaching the volume -v from the local host directory (managed by Docker)
* This is attaching the volume at $PWD/somedir to the (:)mount_location on the container at /somedir 
* The container is newimagename

### Using Build Secrets

* /moby/buildkit for the Docker frontend contains a feature which allows build secrets.

[documentation for that is here](https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md)


# Resources

* [Persistent Volumes for Notebooks Example](https://stackoverflow.com/questions/53201430/notebooks-not-persistent-for-jupyter-in-docker-container)
* [Attaching Volume to Existing Container](https://stackoverflow.com/questions/28302178/how-can-i-add-a-volume-to-an-existing-docker-container)