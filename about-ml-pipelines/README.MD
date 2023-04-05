PAYLOAD='[
    {"Age": 85, "Sex": "male", "Embarked": "S"},
    {"Age": 24, "Sex": '"female"', "Embarked": "C"},
    {"Age": 3, "Sex": "male", "Embarked": "C"},
    {"Age": 21, "Sex": "male", "Embarked": "S"}
]'

curl -X POST -H "Content-Type: application/json"                    \
  -d '{"payload": $PAYLOAD' http://localhost:8080/predict


[
    {"Age": 85, "Sex": "male", "Embarked": "S"},
    {"Age": 24, "Sex": "female", "Embarked": "C"},
    {"Age": 3, "Sex": "male", "Embarked": "C"},
    {"Age": 21, "Sex": "male", "Embarked": "S"}
]


[{"Age":85,"Sex":"male","Embarked":"S"},{"Age":24,"Sex":"female","Embarked":"C"},{"Age":3,"Sex":"male","Embarked":"C"},{"Age":21,"Sex":"male","Embarked":"S"}]



curl http://localhost:8889/query -d '{"model": "sklearn-iris", "version": "1", "data": {"instances": [[1.2, 1.3, 1.2, 1.0]]}}' -H 'Content-Type: application/json'


{"result": {"predictions": [0]}}%  



curl http://localhost:8889/query -d '{"model": "sklearn-iris", "version": "1", "data": {"instances": [[1.2, 1.3, 1.2, 1.0]]}}' -H 'Content-Type: application/json'


curl -X POST http://localhost:8889/query \
  -H 'Content-Type: application/json' \
  -d '{"model": "sklearn-iris", "version": "1", "data": {"instances": [[1.2, 1.3, 1.2, 1.0]]}}'




```
import cherrypy
from cherrypy.lib.sessions import Session

class MySession(Session):
    pass

class MyApplication:
    @cherrypy.expose
    def index(self):
        session = cherrypy.session
        session['count'] = session.get('count', 0) + 1
        return "Hello, world! You have visited this page %s times." % session['count']

if __name__ == '__main__':
    cherrypy.config.update({'tools.sessions.on': True,
                            'tools.sessions.storage_type': 'ram',
                            'tools.sessions.timeout': 60,
                           })
    cherrypy.quickstart(MyApplication())
```