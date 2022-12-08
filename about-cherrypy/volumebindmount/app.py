import cherrypy

class CherryPyExample:
    @cherrypy.expose
    def welcome_page(self):
        return "Welcome!"

if __name__ == '__main__':
    cherrypy.quickstart(CherryPyExample())