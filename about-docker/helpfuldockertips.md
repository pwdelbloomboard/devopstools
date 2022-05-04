#### Clean Up Installs


```
    apt-get clean                                               && \
    rm -rf /var/lib/apt/lists/*
```

##### Removing the Source List /var/lib/apt/lists/*

Don't remove source lists until not planning on installing anything more (freezing dockerfile)

```
    # rm -rf /var/lib/apt/lists/*
```


#### Use Bind Mounts


#### Choose The Most Exact Possible Versions of Things You Are Installing

This one is a little tricky, because the documentation for versioning of whatever the, "things" are that you may be installing varies vs. whatever the, "thing," is.  One foolproof method I have found is to simply spin up a vanilla container, then install whatever tool it is you are looking to install manually, and then after installing it, check the version number.

For example, at one point I needed to install, "software-properties-common" on Debian but I didn't know what version it might be, and I was too lazy to check out and debug what version it may be at [packages.debian.org](https://packages.debian.org/), so I simply installed it on a vanilla version of a Debian container that I spun up without that package installed, then I installed it with apt-get, and then checked the version with:

```
Setting up software-properties-common (0.96.20.2-2.1) ...
```

So right there was the version number.  So within my Dockerfile to build the image, I put:

```
RUN apt-get update                                              && \
    # update first
    apt-get install -y --no-install-recommends                     \
        software-properties-common=0.96.20.2-2.1                   \
```

Now, I still looked at the documentation and read through to get an idea of what the package was for, but I know from experience that different versions of Linux someimtes have funny version 

> This software provides an abstraction of the used apt repositories. It allows you to easily manage your distribution and independent software vendor software sources.

Now that being said - the title of this section is, "Most Exact Possible," because sometimes, it's not really clear what apt-get is looking for within a particular package.  For example, when trying to install wget, the following was attempted:

```
    apt-get install -y --no-install-recommends                     \
        wget=1.2*                                                  \
```

Which resulted in:

```
E: Version '1.21' for 'wget' was not found
```
So to fix this, the following worked:

```
    apt-get install -y --no-install-recommends                     \
        wget=1.2*                                                  \
```
The versioning wildcard * replaces that last character with the most recent possible version that Dockerfile understands, which means if the Dockerfile is rebuilt in the future, at least wget is frozen at 1.2X and not some potentially unexpected 1.3X version.

Now that being said, once the above was executed, wget=1.2* was installed, and we were able to exec into the container, we found that the actual version installed was:

```
dpkg -l | grep wget

wget        1.21-1+deb11u1
```
So from there, we were able to go back into the Dockerfile and use the version, "wget=1.21-1+deb11u1" which worked like a charm.  So one can think of the wildcard functionality as either a hack to get the most precise version possible, or a way to go in and actually find out what the exact version number of that most possibly precise version actually is.


#### Using dpkg -l to Double Check Things are Installing as Expected

As you are going along building a Dockerfile, you may build a beautiful Dockerfile with every line in its place, and you are left thinking that this fantastic Dockerfile is doing its business, installing everything that you need line-by-line just like you told it to, so that when you enter into the Dockerfile using "docker-compose up", everything will be in place.

Of course this isn't always the case, so you can go and use, "dpkg" to see if something you intended to install was in fact installed.  You can further use the, "grep" command to filter out that result rather than having to scroll through a big long list that comes from using just, "dpkg -l".

For example:

```
dpkg -l | grep software-properties-common
```

#### Starting Out With a Slimmer Install Makes Things Harder, But You Can Update

Docker comes with all sorts of variants of linux versions, some of which may be stripped down to the bare minimum of highly specialized images which are only mean to run particular applications, barely an operating system.

As a result, while you're installing packages and checking them with dpkg as discussed above you may get a blank result.  Or you may recieve the infamous, “unable to locate package” error as you are attempting to install things.

What I mean by that is you might use, "FROM debian:bullseye-slim" rather than, "FROM debian:bullseye", resulting in these packages not being available.

One way to deal with this, if you have the leeway in terms of the base image size, or if you're just playing around with things, is to just use the non-slim version of the base image.

However, you could also use, "apt-get update," which goes and retrieves live data and creates a cache required to do all of the installs necessary.

So within bash on the container itself, if using a slim image, you can do::

```
apt-get -y update && apt-get install -y software-properties-common
```

Then, if you run the, "dpkg -l" command discussed above, you may get:

```
# dpkg -l | grep software-properties-common
ii  software-properties-common   0.96.20.2-2.1                  all          manage the repositories that you install software from (common)
```
Of course on the Dockerfile itself, this could be achieved with the following format:

```
# do both updates and upgrades to make up for any missing package repo's if this is a slim image
RUN apt-get -y update                                            && \
    # update first
    apt-get install -y --no-install-recommends                     \
```

##### Difference Between "upgrade" and "update"

While, "update" is helpful to get a list of the available packages, what, "upgrade" does is different and should be used with caution, because of some of the conerns around not knowing precisely how different packages may behave, or what different security concerns may be, mentioned above.

The, "upgrade" function will go in and take all of the packages within a linux-box and upgrade them to the latest version. So if, "apt-get upgrade" is used on a Dockerfile while simultaneously picking versions, it's like saying, "hey, use version X," and then changing one's mind and saying, "OK now upgrade version X to version Y."


#### Watch the Size of Packages As You Install Them To Build a Mental Map Of What Contributes To Your Image Size

If you're building a Dockerfile on the fly by first starting with a vanilla image and then installing parts piece by piece, it may be helpful to pay attention to the size of various entities you install along the way so you can get an idea of what is contributing to the overall final size of your image.  This way you may be able to go back and make more educated decisions on what might be able to be eliminated.

```
After this operation, 141 MB of additional disk space will be used.
Do you want to continue? [Y/n] y
```

If you want to really be persnickety or just keep a nice accounting of what's going on with your Dockerfile, go ahead and even put the disk space used in a comment within the Dockerfile as an, "Estimated disk space used: 141 MB," or something to that effect.

#### Choose Working Directory



#### Use docker-compose

I think it's helpful to get good at starting with using docker-compose and work with the docker-compose.yaml file rather than messing with all sorts of different, "docker run -flag1 -flag2" commands. For me personally, it's hard to keep track of all of the different flags I need, it gets messy and it's much more, "professional," in a sense to just spin up a container in as purely declarative of a manner as I can, which fundamentally means getting better at using docker-compose.yml and knowing how to set all of the settings in that yaml file. The time it takes to set all of those settings appropriately, though it feels frustrating because you don't have a container spun up right away, is worth it for the time spent messing around with various run scripts and flags and trying to remember how you want to spin up your container -- e.g., it just makes things cleaner.

* docker-compose, and principally docker-compose up is designed to be an easy way to spin up a docker container according to a specification, such that if you want to shut it down and have the container itself deleted, you can just do, "docker-compose down," rather than having to rely on entering in a flag that will automatically delete the container after usage as it would with, "docker-compose run".
* Rather, "docker-compose run" is designed more for one-off tasks, where perhaps you need something static written by a docker container of a certain specification. In this case, you could use the command, "docker-compose run --rm" to remove the container after the task has completed...there's a fundamental use case difference between docker-compose up/down and docker-compose run.

##### Within docker-compose, Build Everytime

Under the container you are looking to rebuild, particularly when you're starting off and modifying the base image to get it going in the first place, just set the settings such that:

```
build: .
```
This will rebuild the image every time as you go through your iterations of perfecting your docker image.

##### Difference Between docker-compose up and docker-compose run

From the [Docker docs](https://docs.docker.com/compose/faq/#whats-the-difference-between-up-run-and-start) themselves:

> Typically, you want docker-compose up. Use up to start or restart all the services defined in a docker-compose.yml. In the default “attached” mode, you see all the logs from all the containers. In “detached” mode (-d), Compose exits after starting the containers, but the containers continue to run in the background.

> The docker-compose run command is for running “one-off” or “adhoc” tasks. It requires the service name you want to run and only starts containers for services that the running service depends on. Use run to run tests or perform an administrative task such as removing or adding data to a data volume container. The run command acts like docker run -ti in that it opens an interactive terminal to the container and returns an exit status matching the exit status of the process in the container.

##### Going a Step Further - docker-copose up AND exec 

To make things even faster, you can create a shell script which will do four things, including: bringing the container down, building the image (or rather, rebuilding from scratch), re-compose the docker image into a container and exec in to said container all one after another from one handy-dandy script.

So for example, within our docker-compose.yaml file, we set the name of a container we want to play around in as, "playwithk3d_container" - and I kept exiting it, playing with the Dockerfile base image to build it up, and then running, "docker-compose up" again to spin up a new container, then exec'ing into said container.  Rather than running two commands again and again, you can simply create something like, "composeexec.sh" and enter in the following, so that every time thereafter, when you want to get back into that container, rather than entering in two commands you just do, "./composeexec.sh" - which further hastens the Dockerfile development time.

```
#!/bin/bash
docker-compose down                                     && \
docker build -t k3dpurposed_debian:bullseye-slim .      && \
docker-compose up -d                                    && \
docker exec -t -i playwithk3d_container /bin/bash
```

As a reminder, if you want to exit that and start again, you just do, "exit" within that container, and then immediately do, "docker-compose down," to bring that container back down in order to modify the Dockerfile before trying again.

The reason I recommend including, "build" in the above script is because, if your build doesn't work, even if docker-compose has, "build ." included, it's not going to stop and tell you if something went wrong with the build, it's just going to continue forward with whatever image previously existed, which may have been workable, rather than telling the human, e.g. yourself, that there was a problem with this round of build...which means, hey...you might need to go in there and fix that to make sure it's working as expected.

The "&&" tells bash to only execute the following task if the previous task succeeds (e.g. meaning gives an exit code 0).

##### We Heard You Like Docker Containers - Docker Container in a Docker Container

```
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

Essentially what this is doing, is creating a bind mount between our local machine's Unix socket that the Docker daemon is listening to at /var/run/docker.sock and connecting this to the actual Docker container's socket at /var/run/docker.sock, essentially connecting our physical machine disk to the container. This is not recommended for production as it is insecure.  If you mount docker socket inside a container there are trusted codes running inside that container otherwise you can simply compromise your host that is running docker daemon, since Docker by default launches all containers as root.

Simply stated, you're opening up a root-level window into your host machine through that Docker container, so if anyone ever hacks into your Docker container through whatever method, they now have access to your host machine, the actual hardware as well. This would be different than a dev-mode container, which does not connect to that docker.sock, so there's no root access there.

### Smaller Dockerfile Size Correlated to Smaller Attack Surface

In general, the smaller your Dockerfile is, this likely means there is less going on. While it is possible to make an extremely small Dockerfile on purpose which is just absolutely insecure, what is nice about having a super small Docker image, is typically it's tied to a super small Dockerfile as well, which means there are fewer lines of code to go through, which means it's likely more understandable and fixable by a human if something goes wrong...or rather, there are just fewer variables for a single human to keep in mind while going through it.

Here's a quick thought experiment to demonstrate my point.  Imagine you're reviewing a really short Dockerfile, and it looks something like this:

```
FROM debian:super_insecure_version
```
Did you catch the problem?  It's using the super-insecure version!  You probably caught that right away, whereas if it were a Dockerfile that's hundreds of lines long, whatever problem exists might be harder to find.

Of course, the super_insecure_version might be a super large image size, invalidating my point, but generally when starting from a base distribution of some kind, you start off with at least some fixed size of a layer, and then you continue to add layers and layers of tools on top of that, making the image larger --- so in general, keeping those tools to a minimum, keeping the size of that image smaller is likely going to introduce a smaller physical disk space size image, which roughly correlates to - easier to understand by a human.

So perhaps it's better to say, smaller Dockerfiles tend toward attack surfaces that are managable by humans.



## Resources

https://www.linuxfordevices.com/tutorials/ubuntu/fix-unable-to-locate-package