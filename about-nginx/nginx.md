# Playing Around with Nginx

https://hub.docker.com/_/nginx

## Getting Started

Hosting some simple static content

```
$ docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d nginx
```

Alternatively, a simple Dockerfile can be used to generate a new image that includes the necessary content (which is a much cleaner solution than the bind mount above):

```
FROM nginx
COPY some-content /usr/share/nginx/html
```

then do:

```
docker build -t nginxtest .
```
Finally:

```
docker run --name some-nginx -d nginxtest
```

To expose the port:

```
docker run --name some-nginx -d -p 8080:80 nginxtest
```

Then when you look at the http://localhost:8080, you should see whatever web structure exists within that some-content, which is now being hosted by nginx through the 8080:80 port.


Or, just run directly from the image.

root@26c8bfb92396:/# ls
bin  boot  dev	docker-entrypoint.d  docker-entrypoint.sh  etc	home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var

## What is Nginx

https://www.nginx.com/resources/glossary/nginx/

> The goal behind NGINX was to create the fastest web server around, and maintaining that excellence is still a central goal of the project. NGINX consistently beats Apache and other servers in benchmarks measuring web server performance. Since the original release of NGINX, however, websites have expanded from simple HTML pages to dynamic, multifaceted content. NGINX has grown along with it and now supports all the components of the modern Web, including WebSocket, HTTP/2, gRPC, and streaming of multiple video formats (HDS, HLS, RTMP, and others).

> NGINX Beyond Web Serving

> Though NGINX became famous as the fastest web server, the scalable underlying architecture has proved ideal for many web tasks beyond serving content. Because it can handle a high volume of connections, NGINX is commonly used as a reverse proxy and load balancer to manage incoming traffic and distribute it to slower upstream servers – anything from legacy database servers to microservices.

## Load Balancing

http://nginx.org/en/docs/http/load_balancing.html

Create a file within this directory:

root@f4aab24524d9:/etc/nginx/conf.d# ls
default.conf

```
nano load-balancer.conf
```

```
# Define which servers to include in the load balancing scheme. 
# It's best to use the servers' private IPs for better performance and security.
# You can find the private IPs at your UpCloud control panel Network section.
http {
   upstream backend {
      server 10.1.0.101; 
      server 10.1.0.102;
      server 10.1.0.103;
   }

   # This server accepts all traffic to port 80 and passes it to the upstream. 
   # Notice that the upstream name and the proxy_pass need to match.

   server {
      listen 80; 

      location / {
          proxy_pass http://backend;
      }
   }
}
```

Then from there, configure which servers should be on the backend.

Remove this one: rm /etc/nginx/sites-enabled/default

Then restart the nginx service. systemctl restart nginx

* There are different ways to configure load balancing.
* You can include health checks.

https://upcloud.com/community/tutorials/configure-load-balancing-nginx/