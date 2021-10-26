# Exec'ing Directly into Alpine

You can exec immediately in to an alpine linux VM with:

```
docker run -i -t alpine /bin/sh
```

This would be in contrast to logging in to a specific alpine image such as node:16-alpine:

```
docker run -i -t node:16-alpine
```
Which would actually not put you into the linux machine but rather a node interpreter. Similarly, a python image such as:

```
docker run -i -t python:3.8-slim-buster
```

Would run a python interpreter.  These are based upon linux VM's as well, as listed below.

So, basically we could run a [Debian 10 image as follows](https://hub.docker.com/_/debian?tab=tags&page=1&name=buster-slim):

```
docker run -i -t debian:bullseye-slim /bin/bash
```



# Linux Versions

```
Alpine
Url: https://alpinelinux.org/
Shorty: Its very small.
Packagemanger: apk
Shells: /bin/sh
Size: Few MBs - current tag needs 2.7MB

Jessie aka Debian 8
Url: https://wiki.debian.org/DebianJessie
Shorty: All what you need, but LTS is running out. Click me for details
Packagemanager: apt
Shells: /bin/bash and many more
Size: Round about 50MB

Stretch aka Debian 9
Url: https://wiki.debian.org/DebianStretch
Shorty: All what you need
Packagemanager: apt
Shells: /bin/bash, many more
Size: Round about 40MB

Buster aka Debian 10
Url: https://wiki.debian.org/DebianBuster
Shorty: All what you need, but newer
Packagemanager: apt
Shells: /bin/bash and many more
Size: Round about 50MB

Bullseye aka Debian 11
URL: https://wiki.debian.org/DebianBullseye


Ubuntu based on debain
Url: https://wiki.debian.org/DebianBuster
Shorty: All what you need - and some more
Packagemanager: apt
Shells: /bin/bash and more
Size: Round about 25MB
```