##

https://github.com/moby/buildkit/

### Frontend Syntaxesx

* commands added to the Dockerfile frontend

https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md

#### Mount Secrets

https://github.com/moby/buildkit/blob/master/frontend/dockerfile/docs/syntax.md#run---mounttypesecret

RUN --mount=type=secret
This mount type allows the build container to access secure files such as private keys without baking them into the image.

Option	Description
id	ID of the secret. Defaults to basename of the target path.
target	Mount path. Defaults to /run/secrets/ + id.
required	If set to true, the instruction errors out when the secret is unavailable. Defaults to false.
mode	File mode for secret file in octal. Default 0400.
uid	User ID for secret file. Default 0.
gid	Group ID for secret file. Default 0.