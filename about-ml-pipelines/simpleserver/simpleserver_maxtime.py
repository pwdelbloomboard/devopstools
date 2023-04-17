import cherrypy
import time

class MyApplication:
    @cherrypy.expose
    def index(self):
        if 'start_time' not in cherrypy.session:
            cherrypy.session['start_time'] = time.time()
            cherrypy.session['test'] = 'Session created'
            return 'Session created with a fixed maximum session time.'
        else:
            session_age = time.time() - cherrypy.session['start_time']
            max_session_age = 3600  # Maximum session age in seconds (1 hour)

            if session_age > max_session_age:
                # Expire the current session and create a new one
                cherrypy.lib.sessions.expire()
                cherrypy.session.regenerate()
                cherrypy.session['start_time'] = time.time()
                cherrypy.session['test'] = 'New session created'
                return 'Old session expired. New session created with a fixed maximum session time.'
            else:
                return f'Session is still valid. Age: {session_age:.2f} seconds.'

if __name__ == '__main__':
    config = {
        'global': {
            'server.socket_host': '0.0.0.0',
            'server.socket_port': 8080,
        },
        '/': {
            'tools.sessions.on': True,
            'tools.sessions.timeout': 1,  # Default session timeout in minutes (1 hour)
        },
    }
    cherrypy.quickstart(MyApplication(), '/', config)