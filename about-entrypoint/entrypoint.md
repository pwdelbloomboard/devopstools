# Entrypoints for Dockerfiles / Containers

* When Docker launches, there are two parts, the "entrypoint," and the "command."
* When both are specified, the, "command" part is passed as command-line arguments into the, "entrypoint." part.

## ENTRYPOINT ONLY

In the Dockerfile:

```
ENTRYPOINT /path/to/docker-entrypoint.sh ${ARG1} ${ARG2}
```

## ENTRYPOINT AND COMMAND

In the Dockerfile:

```
ENTRYPOINT /path/to/docker-entrypoint.sh ${ARG1} ${ARG2}
CMD ["whatever"]
```

In the above, CMD ["whatever"] could be, "nginx" for example.

The Entrypoint Script:

```
#!/bin/sh
# ... do some setup ...
# then run the CMD passed as command-line arguments
exec "$@"
```

## Keeping a Container Running After Command Finished

Docker containers, when run in detached mode (the most common -d option), are designed to shut down immediately after the initial entrypoint command (program that should be run when container is built from image) is no longer running in the foreground. 

If you would like to keep your container running in detached mode, you need to run something in the foreground. An easy way to do this is to tail the /dev/null device as the CMD or ENTRYPOINT command of your Docker image.

```
CMD tail -f /dev/null
```

## Bringing the Above Together

There should ideally be a conditional which determines whether a conditional was sent into the entrypoint script. Then within the entrypoing script, toward the bottom should be the following:

```
# If a command was passed in, use it. Otherwise use /app/start.
# Needed to support passing flags to run script for local use.
if [[ -n "$*" ]]; then
    # shellcheck disable=SC2068
    exec $@
else
    /app/start
fi
```

* Where if [[ -n "$*" ]]; determines if anything was fed in as a command to the original script.
* "exec $@" tells the file to execute the program.
* recall that if [[ -n ~ ]] is shorthand for the bash, "test" function, with -n meaning, "if not true," and -s meaning "if true."
* "$*" signifies, "anything"
* $@ signifies the input

If we don't have a specific service we want to run, but we have a setup script, an example way of running things would be:

* The Dockerfile:

```
ENTRYPOINT /path/to/docker-entrypoint.sh
CMD ["tail -f /dev/null]
```

* The entrypoint script:

```
#!/bin/sh

# create a test cluster
k3d cluster create testcluster

# keep container running
exec "$@"
```

## Resources

[Entrypoint vs. Command](https://aws.amazon.com/blogs/opensource/demystifying-entrypoint-cmd-docker/)