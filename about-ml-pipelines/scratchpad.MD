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

