import cherrypy

class Root(object):
    @cherrypy.expose
    def index(self):
        return "Hello World!"

if __name__ == '__main__':
    cherrypy.config.update({'server.socket_port': 8089})
    cherrypy.quickstart(Root(), '/')