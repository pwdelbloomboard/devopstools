# works with json_curl_client_post.sh and json_curl_client_get.sh
import random
import string
import cherrypy


@cherrypy.expose
class RepeatBackWebService(object):

    @cherrypy.tools.json_in() # used to decode request
    @cherrypy.tools.json_out()
    # curl will not work if session not established with cookie-jar
    def GET(self):
        # print("Within the GET function.")
        print(cherrypy.session['mydict'])

    @cherrypy.tools.json_in() # used to decode request
    @cherrypy.tools.json_out() # used to encode response dict to json
    def POST(self, length=8):
        data = cherrypy.request.json
        print("showing json data on console printout: ",data)
        print("showing type of cherrypy.request.json: ",type(data))
        cherrypy.session['mydict'] = data # put value into 'mystring' in session
        print("showing available dict_items in session: ",cherrypy.session.items())
        print("showing type of cherrypy.session['mydict']: ",type(cherrypy.session['mydict']))
        print("showing session data value for ['mydict']: ",cherrypy.session['mydict'])
        return data

    def PUT(self, another_string):
        cherrypy.session['mystring'] = another_string

    def DELETE(self):
        cherrypy.session.pop('mystring', None)


if __name__ == '__main__':
    conf = {
        '/': {
            'request.dispatch': cherrypy.dispatch.MethodDispatcher(),
            'tools.sessions.on': True, # turns sessions on, to identifiy users and synchronize activity
            'tools.response_headers.on': True,
            'tools.response_headers.headers': [('Content-Type', 'text/plain')],
        }
    }
    cherrypy.quickstart(RepeatBackWebService(), '/', conf)