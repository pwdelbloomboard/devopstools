# import requests
import cherrypy
import joblib
import json
import pandas as pd


# load our model
lr = joblib.load('/home/app/model/model.joblib')
model_columns = joblib.load("/home/app/model/model_columns.joblib")

class Server(object):

    @cherrypy.expose
    def index(self):
        # simple function to test if Server working.
        return "Hello World"
    
class API(object):
    exposed = True

    @cherrypy.expose
    @cherrypy.tools.json_in()
    @cherrypy.tools.json_out()
    def query(self, *args, **kwargs):
        # Read the incoming JSON data from the request body
        data = cherrypy.request.json

        # convert categorical variables into dummy/indicator variables indicating presence
        query = pd.get_dummies(pd.DataFrame(data))
        # reindex the columns of query using the model_columns list, adding new columns if missing
        query = query.reindex(columns=model_columns, fill_value=0)
        # use lr, the trained ml model and predict based upon reindexed dataframe 
        prediction = list(lr.predict(query))
        
        # this can be set in the default. This indicates to the serer that the response should be interpreted as json data by the client
        cherrypy.response.headers['Content-Type'] = 'application/json'

        # numpy int64 data type is not JSON seriallizeable by default
        # need to convert each numpy int64 into a python int
        json_prediction = json.dumps([int(x) for x in prediction])

        # Return the JSON response as a dictionary.
        # note, We need to return this as a dict rather than a json object.
        # If we return it as a jsob object, the key will include backslashes ```\"prediction\":``` to indicate that the double quotes are part of the string.
        return {"prediction": json_prediction}

def main():
    # global server config - seperate this from the application config
    global_config = {
        'server.socket_host': '0.0.0.0',
        'server.socket_port': 8889,
    }    

    server_config = {
        '/': {
            'tools.sessions.on': True,
            'tools.sessions.timeout': 3600,
        }
    }

    api_config = {
        '/api': {
            # map HTTP methods (GET, POST, PUT, DELETE, etc.) to methods in the application class, API()
            # Sub'request.dispatch': cherrypy.dispatch.MethodDispatcher(API()),
            'request.dispatch': cherrypy.dispatch.MethodDispatcher(),
            'tools.sessions.on': True,
            'tools.response_headers.on': True,
            # This indicates to the serer that the response should be interpreted as json data by the client
            'tools.response_headers.headers': [('Content-Type', 'application/json')]
        }
    }   

    cherrypy.config.update(global_config)
    cherrypy.tree.mount(Server(),'/',server_config)
    cherrypy.tree.mount(API(), '/api', api_config)
    cherrypy.engine.start()
    cherrypy.engine.block()


if __name__ == "__main__":
    main()