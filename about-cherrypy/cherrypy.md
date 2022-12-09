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

* Nothing happens.  This is because we need to make the file into an executable and run it in the background.  We can put this at the top of our api.py file:

```
#!/usr/bin/env python3
```

Then, we can run it with the following, giving it permissions first:

```
chmod +x app.py
nohup python /home/volumebindmount/app.py &
```
This reads the output of the file to a new file, "nohup.out" unless we specify a different output file with >.

So now, being that we're back on the command line, we can verify that it's still running with:

```
root@47ed88ca5120:/home/volumebindmount# ps -fA
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 03:14 pts/0    00:00:00 /bin/bash
root        14     1  0 03:19 pts/0    00:00:00 python /home/volumebindmount/app.py
root        28     1  0 03:21 pts/0    00:00:00 ps -fA
```

So, we should be able to curl the address shown in our nohup.out file:

```
curl http://127.0.0.1:8080/
<!DOCTYPE html PUBLIC
"-//W3C//DTD XHTML 1.0 Transitional//EN"
"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"></meta>
    <title>404 Not Found</title>
    <style type="text/css">
    #powered_by {
        margin-top: 20px;
        border-top: 2px solid black;
        font-style: italic;
    }

    #traceback {
        color: red;
    }
    </style>
</head>
    <body>
        <h2>404 Not Found</h2>
        <p>The path '/' was not found.</p>
        <pre id="traceback">Traceback (most recent call last):
  File "/usr/local/lib/python3.9/site-packages/cherrypy/_cprequest.py", line 638, in respond
    self._do_respond(path_info)
  File "/usr/local/lib/python3.9/site-packages/cherrypy/_cprequest.py", line 697, in _do_respond
    response.body = self.handler()
  File "/usr/local/lib/python3.9/site-packages/cherrypy/lib/encoding.py", line 223, in __call__
    self.body = self.oldhandler(*args, **kwargs)
  File "/usr/local/lib/python3.9/site-packages/cherrypy/_cperror.py", line 416, in __call__
    raise self
cherrypy._cperror.NotFound: (404, "The path '/' was not found.")
</pre>
    <div id="powered_by">
      <span>
        Powered by <a href="http://www.cherrypy.dev">CherryPy 18.8.0</a>
      </span>
    </div>
    </body>
</html>
```
* So as we can see, there is no path available for, "/" which makes sense, because we have not written a route for this.
* If we move over to index.py and instead serve that, using the methods described above, and then curl the endpoint, we get:

```
:/home/volumebindmount# curl http://127.0.0.1:8089/


        "Hello World!"

```
### Getting Two Servers Running on Different Ports

* If you have one server running, you can run another with a different port.  Note on intro.py we used a different port setup:

```
cherrypy.config.update({'server.socket_port': 8089})
```