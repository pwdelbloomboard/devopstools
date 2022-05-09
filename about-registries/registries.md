# Registries

## Registries in General

* A registry, or Docker registry in the case of Docker, is a system for storing and distributing images with specific names. There may be several versions of the same image, each with its own set of flags. A docker registry is seperated into docker repositories, each of which holds all image modifications.
* A registry may be used by users to fetch images locally and push new images to the registry, given adaqute permissions.
* A registry is usually a server-side application that stores and distributes images, but with k3d there can be a local registry as well to speed up development time.
* Registries are stateless and extremely scalable.
* There are also private registries with Github or Gitlab.

### Pushing an Image to Remote on Heroku Example

As an example, when using Heroku and deploying Docker images to Heroku, after an app was finished to an acceptable state on local, a user had to push an image to a remote heroku registry before it was deployed to Heroku.  That was done as follows:

```
# set the proper env variables if necessary to prevent naming conflicts

# build for production given settings in the docker-compose yaml file
docker-compose -f docker-compose.prod.yml up -d --build

# log into the Heroku docker registry using the docker login tool:
docker login --username=_ --password=$(heroku auth:token) registry.heroku.com

# create the application with the heroku cli
heroku create

# (the cli would output an application name)
# we would set environmental variables and pre-provision any services, including databases as needed prior to pushing
# use the docker command line to tag an image "docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]"

docker tag SOURCE_IMAGE registry.heroku.com/TARGET_IMAGE_NAME/web

# the /web was specified by Heroku
# "release" the image to web as a command fed into Heroku via the heroku command line
heroku container:release web
```
When using Kubernetes, there is a different procedure, because we're going about an entirely different pipeline of instructions. Rather than pushing an image up into a Heroku registry and then asking Heroku to build from that image, which Heroku then uses to run a pre-prepared pipeline of tasks, we instead build an image locally, push it up to a specified registry such as Docker, and then pull that image down with some settings in the various .yaml files in the deployment. Kubernetes is after all from a user perspective, a set of yaml specifications.

### Pulling an Image from a Remote Registry

Months previous to putting together this guide, I had put together and stored an image within a github registry as a way of getting more familiar with how it works.

The Github repo for this project can be found at [pwdelbloomboard.dockerreactjs-yarn](https://github.com/pwdelbloomboard/dockerreactjs-yarn/), while the registry itself can be found at [/pwdelbloomboard/dockerreactjs-yarn](https://github.com/pwdelbloomboard/dockerreactjs-yarn/pkgs/container/ps-container).

Public packages from github can always be found under, "Packages," on the right-hand side of a repo.

Running the following docker command, the local machine will pull down the remote public image from the registry, and run it in a container:

```
docker run -it --rm \
-v ${PWD}:/app \
-v /app/node_modules \
-p 3001:3000 \
-e CHOKIDAR_USEPOLLING=true \
--name buysellguessapp \
ghcr.io/pwdelbloomboard/ps-container
```
To view the container, we can see at:

```
http://localhost:3001/
```
Which shows the app up and running.

So how did we push the image to Github in the first place?

Basically, we followed the steps [talked about here](https://github.com/pwdelbloomboard/devopstools/blob/main/about-k3d/k3d.md#creating-a-github-personal-access-token-pat),

...Which involved run of the mill tagging of the image in the correct way that Github was expecting, and then using the, "docker push" function once the image was properly tagged.

By default, under the free plan this image gets recieved as a public image.

### Using a Local Registry on K3d



## k3d Registries

* [k3d.io Registries](https://k3d.io/usage/guides/registries/)

### Creating a k3d Registry

https://k3d.io/v5.0.0/usage/commands/k3d_registry_create/


## Resources

