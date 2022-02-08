# GoLang

### Basic Setup

* We can get started with GoLang using Docker and a [GoLang Docker image](https://hub.docker.com/_/golang).
* There are other distributions available found on the [GoLang Github](https://github.com/docker-library/golang/tree/dbdde931579e4a3d446b17167c67f573658d6989/1.17).

Using the dockerfile within this folder, we can set up Golang on a local environment running the Docker runtime with:

First, build the image:

```
docker build -t golang-bullseye .
```
Then, run the image, assigning it a name and port:
```
docker run -it --rm                                   \
-p 8883:8883                                          \
--name playwithgolang                                 \
golang-bullseye
```
You should then have access to a go interpreter / shell:

```
root@a631c2ab0aa8:/go# ls
```

This is the bash shell, starting out in the folder /go.

From here, you can start up a go file by installing nano.

```
apt-get update

apt-get -y install nano
```
Now, you can create a simple program:

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
Then, you can build, install packages and run the program with:

```
root@c2ec3ca8f4de:/go# go build packages.go
root@c2ec3ca8f4de:/go# go install packages.go
root@c2ec3ca8f4de:/go# go run packages.go
My favorite number is 1
```

#### Go Shell


### Setting Up Volume for Bind Mount

Let's say we want to keep some of the files we create within the GoLang shell on our local machine for safe keeping, to have for the next time we sign into practice Go.
## Go Playground

* "Go Playground" is a web service that runs on GoLang.org's servers.

https://go.dev/tour/welcome/4
## About Go

[Tour Of Go](/about-go/tourofgo.md)

# Resources

* [Tour of Go](https://go.dev/tour/welcome/1)
* [Golang in Juputer Notebook](https://levelup.gitconnected.com/running-golang-on-jupyter-notebook-f7f9fba37812)