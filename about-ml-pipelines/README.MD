# About

### Getting Started - Creating a Simple Server

* In order to have an endpoint which is queryable, we have to have a session.

The simplest way to establish a server which is queryable by a ```curl``` function is to do the following:

```
import cherrypy

class MyApplication:
    @cherrypy.expose
    def index(self):
        # Enable sessions for this endpoint
        cherrypy.session['test'] = 'Session created'
        return 'Session created. You can use the session ID in the following requests.'

if __name__ == '__main__':
    config = {
        'global': {
            'server.socket_host': '0.0.0.0',
            'server.socket_port': 8080,
        },
        '/': {
            'tools.sessions.on': True,
            'tools.sessions.timeout': 60,  # Session timeout in minutes
        },
    }
    cherrypy.quickstart(MyApplication(), '/', config)
```

After running the above server with ```python simpleserver.py```, you can use curl to create a session with:

```
$ curl -i -c cookies.txt http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/html;charset=utf-8
Server: CherryPy/18.8.0
Date: Tue, 11 Apr 2023 21:59:54 GMT
Content-Length: 70
Set-Cookie: session_id=2539232933a6b4a8ae992b937a77509e67245fba; expires=Tue, 11 Apr 2023 22:59:54 GMT; Max-Age=3600; Path=/

Session created. You can use the session ID in the following requests.
```
### Moving Along - Maintaining the Session

To maintain a session after it's created using CherryPy, you need to ensure that the client (in this case, curl) sends the session cookie back to the server with each subsequent request. This allows the server to recognize the client and associate it with the correct session data. Here's how you can maintain the session:

When you first create the session using the curl command, make sure to save the received cookies, including the session ID, in a file (e.g., cookies.txt):
arduino

```
curl -i --cookie-jar cookies.txt http://localhost:8080
```

For each subsequent request to the CherryPy application, use the -b option to send the cookies stored in cookies.txt back to the server:

```
curl -i -cookie cookies.txt http://localhost:8080/another_endpoint
```

or alternatively, the value could be fed in directly with:

```
curl --cookie "session_id=f586737cd3e586a435f9d45df4bf2b779aacfe22" http://localhost:8080/
```

Replace another_endpoint with the desired endpoint in your CherryPy application. This request will use the session established in the previous step.

So with the above setup, when you get a random session_id for a cookie, you get an expiration time. So for example, at the bottom of the HTTP message, you get:

```
Set-Cookie: session_id=b8c717f96f9baa95513b781b27ee0d9c8e7c411e; expires=Wed, 12 Apr 2023 21:47:17 GMT; Max-Age=3600; Path=/
```
If you ping again with the same session_id, the expiration date updates, does not keep the same sesion id, you just get a new expiration date pushed into the future.

```
Set-Cookie: session_id=b8c717f96f9baa95513b781b27ee0d9c8e7c411e; expires=Wed, 12 Apr 2023 21:47:39 GMT; Max-Age=3600; Path=/
```

So, if you wanted to set the Max-Age of the session to be shorter, you can set the ```timeout``` with the config via:

```
if __name__ == '__main__':
    config = {
        'global': {
            'server.socket_host': '0.0.0.0',
            'server.socket_port': 8080,
        },
        '/': {
            'tools.sessions.on': True,
            'tools.sessions.timeout': 60,  # Session timeout in minutes (1 hour)
        },
    }
    cherrypy.quickstart(MyApplication(), '/', config)
```

So pinging with curl again, you see the Max-Age having been shortened:

```
Set-Cookie: session_id=e8022c4816d12917da2e334b94f121ac0f00a3a9; expires=Wed, 12 Apr 2023 21:15:43 GMT; Max-Age=300; Path=/
```

If you wanted to force a logout and make a session_id restart after a particular time period, basically a maximum session time based upon the first cookie, there has to be some kind of custom handle and way to check whether the current session_id has exceeded the maximum session time.