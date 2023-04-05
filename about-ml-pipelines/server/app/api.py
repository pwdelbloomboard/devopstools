import requests
import cherrypy
import sklearn
import joblib

# load our model
lr = joblib.load('/home/app/model/model.joblib')

# establish session
# session = requests.Session()
# session_cookie = ''

#def query(model_name, version, json):
#    model_api = "/api/" + model_name + "/" + str(version) + "/"
#    response = session.post(HOST + model_api, headers=HEADERS, json=json)
#    return response

class Server(object):

    @cherrypy.expose
    def index(self):
        cherrypy.session['fieldname'] = 'fieldvalue';
        return "Hello World!"

    @cherrypy.expose
    @cherrypy.tools.json_out()
    @cherrypy.tools.json_in()
    def query(self):
        try:
            json = cherrypy.request.json
            model = json['model']
            version = json['version']
            data = json['data']
            resp = query(model, version, data)
            if (resp.status_code != 200):
                raise cherrypy.HTTPError(500, "Query Error ({})".format(resp.status_code))
            print(resp.content)
            return {"result": ast.literal_eval(resp.content.decode('utf-8'))}
        except KeyError as err:
            raise cherrypy.HTTPError(400, "Bad Request for key: {}".format(err))
        except Exception as err:
            raise cherrypy.HTTPError(500, "Unknown Error: {}".format(err))


def main():
    # global server config - seperate this from the application config
    cherrypy.config.update({
    'server.socket_host' : '0.0.0.0',
    'server.socket_port' : 8889,
    })

    # applicaton config is provided
    cherrypy.tree.mount(Server(), '/', {
        '/' : {
        'tools.sessions.on' : True,
        'request.dispatch'  : cherrypy.dispatch.MethodDispatcher(),
        'tools.sessions.timeout': 60,
        },
        '/string' : {
        'tools.response_headers.on'      : True,
        'tools.response_headers.headers' : [('Content-Type', 'text/plain')]  
        }
    })

    cherrypy.quickstart(Server())



if __name__ == "__main__":
    main()