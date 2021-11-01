# Dockerfile

## Dockerfile Best Practices

In general, Docker has an extensive list of best practices for creating dockerfiles, based upon the nature of how docker (or perhaps containerized systems in general) work.

* [General Best Practices for Dockerfiles - docs.Docker.com](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

Beyond this, there may be organizational goals which actually priortize and rank these above best practices for actual strategic reasons.

For example, the following may be priorities, ranked and hypothetically discussed:

* Fast Build Time (perhaps the development cycle is very fast)
* Maximize Use of Cache on Rebuilds
* Small Image Size (perhaps there are many pods in a giant cluster, and large image size becomes expensive)
* Security (perhaps your organization is being constantly attacked, maybe it's a bank)
* Consistency (perhaps you run a lean organization and need to reduce change to keep individuals from being overloaded with tasks in the case of change or a downtime emergency)

### Seperating Build Files for Different Environements to Optimize Build Times

Given the layering inherent in Docker builds, the importance of designinger Dockerfiles to support faster build times, and the presence of different pipeline stages (e.g. local vs. production), it is helpful to create a branching strategy that seperates Dockerfiles either serially or in parallel.

One typical practice that is often seen, is that a dockerfile is built which includes multi-stage building, and is organized in such a way that the, "basic, less frequently changing elements," occur earlier in the dockerfile than the, "more frequently changing elements."

The weakness of a single-file approach is that it is strictly linear and has no branching capability based upon what environment may be running. An alternative approach would be to break into three seperate files which branch based upon the environment:

* Base Docker Files (base.dockerfile), which runs no matter what, and contains the, "most important stuff."
* Local Docker Files (local.dockerfile) which runs only on local machines, development machines.
* Prod Docker Files (prod.dockerfile) which runs on production environments, and has elements specific to production.

Prod may also be called, "cloud" and cover Staging as well as Production.

It is also possible to create a, "test.dockerfile," which tests tools and dependencies which take a long time to install. For example, with Node.js, it could create a simple way to run tests on javascript dependencies to ensure they are working, before proceding on to, "cloud.dockerfile" as an example.

### Maximizing Use of Cache

One of the basics of how Docker operates is that it uses layers to build an application from the virtual machine level upward to dependencies and then the application itself (depending of course on how the Dockerfile is designed).

Docker will cache the results of the first build of a Dockerfile, allowing subsequent builds to be faster.  Keeping this principle in mind is critical to gaining some of the benefits of docker (otherwise why not just download and work with entire linux virtual machines every time!?).

#### Use of .dockerigore

Right away, it's important to cut out the fat and speed things up.

#### Use Seperate Run Instructions to Install Dependencies and Copy Code

While this block of code reads cleaner:

```
COPY . /app
RUN pip install -r requests.txt
```
What it does is causes requirements to be installed every time changes to any file are made in an entire codebase. A better practice would be the below, which instead copies the actual app directory itself (meaning the code written by a developer on a team), after the requirements and dependencies are built in the above command.

```
COPY requests.txt /app/requests.txt
RUN pip install -r requests.txt

COPY . /app
```
Other best practices include:

* try to group similar cacheable things together.
* pay close attention to the order of these commmands!

#### Cleaning Up After Package Managers

Package mangers, such as apk, apt-get, bower, composer, npm, pip, leave a bunch of crap on the machine in cache after they run. This crap does not automatically go away just because an image is built, there is nothing explicitly telling the image to remove caches. An image is basically the state of a machine after a particular point, frozen in carbonite.

So, you can delete the cache from these package managers with various commands specific to those package mangers, for example:

```
npm ci && rm -rf "$HOME/.npm"
```

...accomplishes this for NPM.

### Not Running as Root within an Environment (Security Practice)

As with any environment, there's just no reason to run things in root unless explicitly required - e.g. setting up a machine in the first place.

Instead, a userID of for eample 1000 can be set up and then a k8s configuration can be set up to allow a user with an ID of 1000 should run things:

```
securityContext:
    runAsNonRoot: true
    runAsUser: 1000
```

That being said, during the build of a particular file on Docker itself, you may run certain setup(s) as root, and then switch to a non-root user once appropriate permissions have been assigned to that subsequent user.


## Use of Different Base Files for Docker Buildfiles

Within an entire cornacopia of base files to choose from, it's important to choose from altarnatives and ensure that the base image size is minimized.

One notable example is that a base image size using python-slim:buster which later installs npm was found to be smaller in size than a base image using npm which then installs python.

## Setting Up an Environment vs. Running Jobs

### The Difference Between CMD and ENTRYPOINT

When you execute, "docker run," on the command line, you must specify an ENTRYPOINT or CMD.

When you start a docker container, you are essentially starting a linux distribution, and within that, a shell (meaning a command line interpreter) such as /bin/sh or /bin/bash to get started, meaning by default, if you use, "docker run --interactive" (which keeps STDIN open even if not attached) or "docker run -t" (tty or TeleTYpewriter, terminal, the text input/output environment)

ENTRYPOINT or CMD both can specify the the default executable when a linux distro starts up. However, the user can override these at runtime by feeding in arguments on the command-line with, "docker run [ARGUMENTS]".

* CMD defines default commands and/or parameters for a container. CMD is an instruction that is best to use if you need a default command which users can easily override. If a Dockerfile has multiple CMDs, it only applies the instructions from the last one.

* On the other hand, ENTRYPOINT is preferred when you want to define a container with a specific executable. You cannot override an ENTRYPOINT when starting a container unless you add the --entrypoint flag.

#### Key Usage Difference

If you use ENTRYPOINT and CMD together, then CMD becomes an argument for ENTRYPOINT.

Example:

```
ENTRYPOINT ["/app/entrypoint"]
CMD ["/app/start"]
```

* So in the above, we use the, "entrypoint" script to run the actual file, with all sorts of imperative commands running in sequence, whereas "CMD" just feeds the argument, "/app/start" into the script, "/app/entrypoint".


The name, "Entrypoint" is a bit of a misnomer, in that "Entrypoint" defines the "set of commands" which runs the command to get something going, unless it is overridden explicitly, whereas "CMD" just runs the most recent CMD as one line fed into the container at the very end of the file.

##### Importance of Using Both Entrypoint and CMD

* Why would a set of applications use both ENTRYPOINT and CMD?  Why go through all the hassle?

Basically, modern applications are run on clusters, such as k8s. Within a cluster, within various different sets of environments in a cluster, an administrator may want to start and run a, "new job," at some point in time, perhaps to check the health of a pod.

The way to access and start jobs, is basically through running code through injecting values or scripts into the environment which is running an application. Of course the way we inject those values into the environment, is the same way we inject environmental variables and config files into that environment - the dockerfile(s).

So essentially, dockerfiles are a sort of bottleneck into an environment (or pod), they must be run through to set up a pod or environment in the first place, but they also must be run through in order to just run a simple new job (without re-starting the entire environment).

The way to reconcile this, "bottleneck," is basically to use, "CMD" last, because CMD can simply feed arguments into entrypoints.

By keeping this CMD variable last, we are basically seperating the setting up of an entire environment from just starting or re-starting an application which is sitting on that environment, making it easier to run scripts by just overriding that CMD value without having to re-run or re-start that entire environment, just the application.

### Differences Between terminal, shell, tty, console

* terminal = tty = text input/output environment
* console = physical terminal
* shell = command line interpreter

> Console, terminal and tty are closely related. Originally, they meant a piece of equipment through which you could interact with a computer: in the early days of unix, that meant a teleprinter-style device resembling a typewriter, sometimes called a teletypewriter, or “tty” in shorthand. The name “terminal” came from the electronic point of view, and the name “console” from the furniture point of view. Very early in unix history, electronic keyboards and displays became the norm for terminals.

* So basically, paper output became terminal outputs, you're talking about the Machine Level.

* The term, "tty" goes as far back as 1869, the teletype machine (tty) which was used for stock tickers.  tty is a super old technology, with telegraph-age origin.  Tty's were eventually connected to computers through adapters called UARTs, which was known as a, "terminal" because it was the terminal of a physical line connecting a tty to a computer.

* Today, a tty is a particular type of device file which implements a bunch of commands beyond read and write. Some terminals are provided by the linux kernal (os) on behalf of whatever hardware device or VM being used.

> A console is generally a terminal in the physical sense that is by some definition the primary terminal directly connected to a machine. The console appears to the operating system as a (kernel-implemented) tty. On some systems, such as Linux and FreeBSD, the console appears as several ttys (special key combinations switch between these ttys); just to confuse matters, the name given to each particular tty can be “console”, ”virtual console”, ”virtual terminal”, and other variations.

* So basically, a console may be another form of thinking about a tty, a kind of, "custom" tty that is provided by the linux distro to help things work. It's like an adapter to the terminal.

> A shell is the primary interface that users see when they log in, whose primary purpose is to start other programs. (I don't know whether the original metaphor is that the shell is the home environment for the user, or that the shell is what other programs are running in.)  In unix circles, shell has specialized to mean a command-line shell, centered around entering the name of the application one wants to start, followed by the names of files or other objects that the application should act on, and pressing the Enter key. Other types of environments don't use the word “shell”; for example, window systems involve “window managers” and “desktop environments”, not a “shell”.

* So basically, a shell is a higher-level, more application-level type interface which talks to the below console/terminal.

* Shells may include, "bash" or "zsh" - they are like user-friendly windows, but text-input rather than click input. They may be customized.

* Shells may also give the capability to write scripts, hence the term, "shell programming."  Shells are almost Bourne Shells or Bourne Again Shells (bash) but there are different varieties out there, and programming on those different varieties may differ slightly, and you have to reference the documentation.

* Job control (launching programs and managing them) is mostly performed by the shell. It's kind of like a user interface. However, the terminal would control and work with interupts such as, CTL+ALT+DELETE or whatever other interupt...it is a machine-layer, closer to the metal layer that listens and has higher priority for those types of things.

## Other Topics

### Syntax

# syntax = docker/dockerfile:1.0-experimental

### Details on Important Dockerfiles Commands

#### SHELL

[SHELL](https://docs.docker.com/engine/reference/builder/#shell)

> The SHELL instruction allows the default shell used for the shell form of commands to be overridden. The default shell on Linux is ["/bin/sh", "-c"]

* Basically with "SHELL" we are explicitly calling out a particular type of shell, for example, using, "/bin/bash" may be preferable because bash has more features than sh.
* Other arguments which can be fed in, within the context of the set operation include:

* -e ... means exit on error.
* -u ... means unbind all keys bound to the named function
* -o ... set according to the option-name, which can include "pipefail," which if set returns the value of the last (rightmost) command to exit with a non-zero status. This option is disabled by default.

* [the gnu set function](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#The-Set-Builtin)
* [gnu bash arguments](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Invoking-Bash)

More specifically stated:

* SHELL is used to define the shell that should execute the commands while building the image from the Dockerfile, from within the Docker runtime environment.  Basically Docker the application is running on top of an OS, this Docker runtime environment can have commands run, and we can pick which shell runs those commands with, "SHELL". It does not have any impact on the image that gets built. https://docs.docker.com/engine/reference/builder/#shell


#### ENV

[ENV](https://docs.docker.com/engine/reference/builder/#env)

* Quite simply these are environment variables.

These can be used to set the language, defaults across all dockerfiles used across a k8s cluster, various settings to make things work.

#### ARG

[ARG](https://docs.docker.com/engine/reference/builder/#arg)

> The ARG instruction defines a variable that users can pass at build-time to the builder with the docker build command using the --build-arg <varname>=<value> flag. If a user specifies a build argument that was not defined in the Dockerfile, the build outputs a warning.

ARG's are not persisted to the built images while ENV variables are. However, if an ARG value changes from a previous build, then a, "cache miss" occurs upon its first usage, which can slow down a build by forcing a start at an earlier layer.

##### Default ARG's

* HTTP_PROXY
* http_proxy
* HTTPS_PROXY
* https_proxy
* FTP_PROXY
* ftp_proxy
* NO_PROXY
* no_proxy

There are also automatic ARGs which are set to specify platform by Docker itself, based upon the system Docker is running on.

#### RUN

* [RUN](https://docs.docker.com/engine/reference/builder/#run)

> RUN has 2 forms:

> RUN <command> (shell form, the command is run in a shell, which by default is /bin/sh -c on Linux or cmd /S /C on Windows)

> RUN ["executable", "param1", "param2"] (exec form)

##### RUN exec Form

> The exec form makes it possible to avoid shell string munging, and to RUN commands using a base image that does not contain the specified shell executable.

The exec form is parsed as a JSON array, which means that you must use double-quotes (“) around words not single-quotes (‘).

##### RUN shell Form

The default shell for the shell form can be changed using the SHELL command.

So as an example, you could [add a user](/commandline-examples/useradd.md) based upon a UID and a [group id](/commandline-examples/groupmod.md) (GID) fed in like so:

```
ARG UID=1000
ARG GID=1000
RUN useradd -u $UID -m -s /bin/bash node && \
    groupmod -g $GID node
```

#### COPY



#### VOLUME

* [VOLUME](https://docs.docker.com/engine/reference/builder/#volume)

> The VOLUME instruction creates a mount point with the specified name and marks it as holding externally mounted volumes from native host or other containers. The value can be a JSON array, VOLUME ["/var/log/"], or a plain string with multiple arguments, such as VOLUME /var/log or VOLUME /var/log /var/db

[About Volumes in the Context of Docker](/about-docker/docker-volumes.md)

> Volumes are the preferred mechanism for persisting data generated by and used by Docker containers. While bind mounts are dependent on the directory structure and OS of the host machine, volumes are completely managed by Docker. Volumes have several advantages over bind mounts:


# Using Different Names for Dockerfiles

* [How to Name Dockerfiles - Stackoverflow](https://stackoverflow.com/questions/26077543/how-to-name-dockerfiles)

If you're attempting to build with Docker using, "docker build X" but your dockerfile name is something like, "test1.Dockerfile" - then it won't work right off the bat - you may get the following error:

> failed to solve with frontend dockerfile.v0: failed to read dockerfile: error from sender: walk base.Dockerfile: not a directory

It would need to be along the lines of:

```
test1.Dockerfile

docker build -f dockerfiles/test1.Dockerfile  -t test1_app .
```





# Resources

* [General Best Practices for Dockerfiles - docs.Docker.com](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
* [Exact Difference Between TTY, Console, Shell, etc](https://unix.stackexchange.com/questions/4126/what-is-the-exact-difference-between-a-terminal-a-shell-a-tty-and-a-con)
* [CTL - Dockerfile Entrypoint vs Command](https://www.ctl.io/developers/blog/post/dockerfile-entrypoint-vs-cmd/)
* [Docker Entrypoint vs CMD](https://phoenixnap.com/kb/docker-cmd-vs-entrypoint)
* [Understanding Docker Cache](https://thenewstack.io/understanding-the-docker-cache-for-faster-builds/)
* [How to Name Dockerfiles - Stackoverflow](https://stackoverflow.com/questions/26077543/how-to-name-dockerfiles)