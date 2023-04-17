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
            'tools.sessions.timeout': 5,  # Session timeout in minutes
        },
    }
    cherrypy.quickstart(MyApplication(), '/', config)