# dev-mode Container

* Using the concept of volumes, we can attach source code to a container such that if the source code is changed, the application updates on the fly without having to rebuild and restart the application.

The original documentation for dev-mode containers can be found [here](https://docs.docker.com/get-started/06_bind_mounts/#start-a-dev-mode-container).

Essentially, once you have a container running, you can run the following type of command:

```
 docker run -dp 3000:3000 \
     -w /workingdirectory -v "$(pwd):/app" \
     node:12-alpine \
     sh -c "yarn install && yarn run dev"
```

* -dp 3000:3000 - Runs the application in detached (background) mode and create a port mapping
* -w /workingdirectory sets the working directory from which the command will run
* -v "$(pwd):/workingdirectory" bind mounts the current directory on the local machine to the workingdirectory on the container
* In this case, node:12-apline was the image being used, could be the base image from the dockerfile, named/tagged.
* The command is the last line being run, "sh -c "yarn install && yanr run dev" - this would apply to essentially a javascript app being built with yarn as the first command. However if it were for example, to run the fluent-bit binary, that might be the line, "/fluent-bit/bin/fluent-bit"