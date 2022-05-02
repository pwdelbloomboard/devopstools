# Hugo

## Getting Started

[hugo quickstart](https://gohugo.io/getting-started/quick-start/)

```
apt-get -y install hugo
hugo version
Hugo Static Site Generator v0.80.0/extended linux/amd64 BuildDate: 2021-07-18T09:31:51Z (debian 0.80.0-6+b5)
```

Then do:

```
hugo new site quickstart
Congratulations! Your new Hugo site is created in /home/quickstart.

Just a few more steps and you're ready to go:

1. Download a theme into the same-named folder.
   Choose a theme from https://themes.gohugo.io/ or
   create your own with the "hugo new theme <THEMENAME>" command.
2. Perhaps you want to add some content. You can add single files
   with "hugo new <SECTIONNAME>/<FILENAME>.<FORMAT>".
3. Start the built-in live server via "hugo server".

Visit https://gohugo.io/ for quickstart guide and full documentation.
```

There is now a new folder, "quickstart," in the directory that we just created.

```
cd quickstart
git init
git submodule add https://github.com/theNewDynamic/gohugo-theme-ananke.git themes/ananke
Cloning into '/home/quickstart/themes/gohugo-theme-ananke'...
remote: Enumerating objects: 2496, done.
remote: Counting objects: 100% (519/519), done.
remote: Compressing objects: 100% (286/286), done.
remote: Total 2496 (delta 261), reused 403 (delta 199), pack-reused 1977
Receiving objects: 100% (2496/2496), 4.45 MiB | 7.15 MiB/s, done.
Resolving deltas: 100% (1357/1357), done.
```

Then, put this into the configuration.

```
cd ..
echo theme = \"ananke\" >> config.toml
cd themes
mv gohugo-theme-ananke/ ananke/
```
We had to use that mv command to rename the gohugo-theme-anake/ folder to anake/ to work with our config.toml file.

```
cd ..
hugo new posts/my-first-post.md
WARN 2022/04/28 20:55:45 Module "ananke" is not compatible with this Hugo version; run "hugo mod graph" for more information.
posts/my-first-post.md created
```

Since that theme is not compatible with the hugo version we can go ahead and try to get more information if it doesn't work after serving.

```
hugo server -D
WARN 2022/04/28 20:56:58 Module "ananke" is not compatible with this Hugo version; run "hugo mod graph" for more information.
Start building sites …

                   | EN
-------------------+-----
  Pages            |  7
  Paginator pages  |  0
  Non-page files   |  0
  Static files     |  1
  Processed images |  0
  Aliases          |  0
  Sitemaps         |  1
  Cleaned          |  0

Built in 92 ms
Watching for changes in /home/quickstart/{archetypes,content,data,layouts,static,themes}
Watching for config changes in /home/quickstart/config.toml, /home/quickstart/themes/ananke/config.yaml
Environment: "development"
Serving pages from memory
Running in Fast Render Mode. For full rebuilds on change: hugo server --disableFastRender
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop
```

Unfortunately, to view the website we would have had to run the container differently, connecting to a local port before viewing this Web Server.

It may be necessary to build a different kind of Docker image with Go included and connect up the default port options with a .yaml file to be able to more easily view and repeat, iterate and play around with Hugo.

Using a simple run command may involve something like the following:

```
docker run -d --add-host host.docker.internal:host-gateway my-container:latest
```

Using docker compose, this might be via adding the following to a service:

```
    extra_hosts:
      - "host.docker.internal:host-gateway"
```
Or alternatively:

```
    extra_hosts:
      - "host.docker.internal:172.17.0.1"
```

We may also need to add ports:

```
    ports: 
      - 80:3000
```

Which, may do the work of exposing this to the outside world.

However, when we attempted to do this, we got:

```
ERROR: for playwithgolang_container  Cannot start service playwithgolang_service: driver failed programming external connectivity on endpoint playwithgolang_container (a036463e23c45af7cd1132af135947c7a23a1cd6825b224ed2cd23816b0186d9): Bind for 0.0.0.0:80 failed: port is already allocated
```

So, we can use a different port:


```
    ports: 
      - 1313:3002
```

Which should ideally connect our inside ports where hugo gets served, 1313 to the outside world, 3002.

If we serve a docker container with hugo properly, as discussed under the [about go](/about-go/) folder, we should be able to exec in and start working with hugo immediately, along with a bind mount available.

After running all of the above commands getting hugo up and running again, on our local machine we can go to:

```
localhost:3002
```

This didn't work, so it's possible we have to reverse the ports:

```
    ports: 
      - 3002:1313
```
Which also did not work, so we can try:

```
    ports: 
      - 1313:1313
```
So all of the above being said, part of the issue was to be able to run the Docker container in such a way that the port was bound to an outside port, but also the server was bound to 0.0.0.0 internally, and an extra host was used that bounded the internal to the host-gateway, as follows:

```
version: '3.1'

services:
  playwithgolang_service:
    image: golang-bullseye:latest
    build: .
    container_name: playwithgolang_container
    working_dir: /home/startcontent/quickstart/
    command: hugo server --bind 0.0.0.0 -D
    ports: 
      - 1313:1313
    volumes:
      - type: bind
        source: ./volumebindmount
        target: /home/startcontent
    extra_hosts:
      - "host.docker.internal:host-gateway"
    tty: true
```

So from here, the question would be, how do we actually generate a static site from the served demo?  Basically, just the command, "hugo" with no other flags.

```
/home/startcontent/quickstart# hugo
WARN 2022/05/01 12:02:49 Module "ananke" is not compatible with this Hugo version; run "hugo mod graph" for more information.
Start building sites …

                   | EN
-------------------+-----
  Pages            |  7
  Paginator pages  |  0
  Non-page files   |  0
  Static files     |  1
  Processed images |  0
  Aliases          |  0
  Sitemaps         |  1
  Cleaned          |  0

Total in 300 ms
```

A folder called, "public" will be generated which contains the css, images, tags, index, etc. However, looking into that folder we don't see the page rendered as expected. This could hypothetically be because of the warning message mentioned above, stating that the theme, "ananke" is not compatible with this Hugo version.

```
hugo version
Hugo Static Site Generator v0.80.0/extended linux/amd64 BuildDate: 2021-07-18T09:31:51Z (debian 0.80.0-6+b5)
```
When we do an, "inspect" on the webpage via chrome, and look under, "Console," we can see that there is an error:

```
/ananke/css/main.min.css:1          Failed to load resource: net::ERR_FILE_NOT_FOUND
```

This may perhaps be because the webpage is looking at, "/anake/otherstuff" rather than "anake/otherstuff" - essentially, anake is not a root folder.  If we were to host this on a location such as s3, it may work.

We can attempt a free static site upload test with [Tiiny.host](https://tiiny.host/).

Attempting to copy the files to this host, we still get an error:

```
GET https://testitout.tiiny.site/ananke/css/main.min.css net::ERR_ABORTED 404 
```

Part of this now could be due to how we have the configuration settings setup within Hugo.  If we change the baseURL to the following rather than, "https://exampledomain.com/" this may resolve the css.

```
baseURL = "https://testitout.tiiny.site/"
languageCode = "en-us"
title = "My New Hugo Site"
theme = "ananke"
```

After re-rendering with the command, "hugo" we can see that one of the meta tags was re-rendered differently as:

```
<meta property="og:url" content="https://testitout.tiiny.site/" />
```

Attempting to re-upload still results in an error.

If we switch over to AWS S3 and upload all of our documents to try again, we can try the endpoint:

```
http://testitout.tiiny.site.s3-website-us-west-2.amazonaws.com/
```

Which we used as a default after attempting testitout.tiny.site and it didn't work. However on S3 we get a 403 error, even when static hosting is enabled.  This was because we failed to attach the following policy:

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "PublicReadGetObject",
            "Effect": "Allow",
            "Principal": "*",
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::testitout.tiiny.site/*"
        }
    ]
}
```

Once this was attached properly, the site worked at this location:

```
http://testitout.tiiny.site.s3-website-us-west-2.amazonaws.com/
```

However, there was no post included, e.g.: ```posts/my-first-post/```

Of course, this could be because there are no posts that were actually generated into the, "public" folder upon serving our actual site.


After running hugo again, we see immediately that there are 10 pages created rather than 7.

```
WARN 2022/05/01 18:13:52 Module "ananke" is not compatible with this Hugo version; run "hugo mod graph" for more information.
Start building sites …

                   | EN
-------------------+-----
  Pages            | 10
  Paginator pages  |  0
  Non-page files   |  0
  Static files     |  1
  Processed images |  0
  Aliases          |  1
  Sitemaps         |  1
  Cleaned          |  0

Total in 387 ms
```

And sure enough, within the, "public" folder, there was a posts/ folder with posts included.

Uploading to S3, we had to go and upload everything rather than just the new posts/ folder to ensure that it worked completely and was listed on the front page.