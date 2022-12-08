# About CherryPy

https://docs.cherrypy.dev/en/latest/

* Minimalist python web framework.

## Installing and Versioning

* Within requirements.txt we have:

```

```

Otherwise manually:

```
pip install CherryPy
```

## Setting Up Experiment Docker Container

* The following 

```
./buildimageandtag.sh
docker-compose run cherrypy_service /bin/bash
```

* Logging in and navigating to /volumebindmount and running, ```python app.py``` will give us:

```
root@b394fde6ce22:/home/volumebindmount# python app.py
[08/Dec/2022:21:15:01] ENGINE Listening for SIGTERM.
[08/Dec/2022:21:15:01] ENGINE Listening for SIGHUP.
[08/Dec/2022:21:15:01] ENGINE Listening for SIGUSR1.
[08/Dec/2022:21:15:01] ENGINE Bus STARTING
CherryPy Checker:
The Application mounted at '' has an empty config.

[08/Dec/2022:21:15:01] ENGINE Started monitor thread 'Autoreloader'.
[08/Dec/2022:21:15:01] ENGINE Serving on http://127.0.0.1:8080
[08/Dec/2022:21:15:01] ENGINE Bus STARTED
```

* This is using the very basic function, which should display a welcome message:

```
class CherryPyExample:
    @cherrypy.expose
    def welcome_page(self):
        return "Welcome!"
```

* So what?  IF we try to curl:

```
# curl -v http://127.0.0.1:8080
*   Trying 127.0.0.1:8080...
```

* Nothing happens.

### Getting Two Servers Running on Different Ports

* If you have one server running


```
root@bc8384074626:/home/volumebindmount# ps -fA
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 21:17 pts/0    00:00:00 /bin/bash
root         9     1  0 21:17 pts/0    00:00:00 python app.py
root        28     1  1 21:20 pts/0    00:00:00 python intro.py
root        43     1  0 21:21 pts/0    00:00:00 ps -fA
```