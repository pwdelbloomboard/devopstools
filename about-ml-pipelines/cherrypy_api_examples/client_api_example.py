#!/bin/bash

import requests

s = requests.Session()

r = s.post('http://127.0.0.1:8080/')

print("Executing POST: s.post('http://127.0.0.1:8080/')")
r = s.post('http://127.0.0.1:8080/')
print("Status Code: ",r.status_code)
print("Output: ",r.text)

print("Executing GET: r = s.get('http://127.0.0.1:8080/')")
r = s.get('http://127.0.0.1:8080/')
print("Status Code: ",r.status_code)
print("Output: ",r.text)

f_string = """whatevs"""

print("Executing PUT: r = s.put('http://127.0.0.1:8080/', params={'another_string': 'hello'})")
r = s.put('http://127.0.0.1:8080/', params={'another_string': {f_string}})

print("Executing GET: r = s.get('http://127.0.0.1:8080/', headers={'Accept': 'application/json'}")
r = r = s.get('http://127.0.0.1:8080/')
print("Status Code: ",r.status_code)
print("Output: ",r.text)

print("Executing DELETE: r = s.delete('http://127.0.0.1:8080/')")
r = s.delete('http://127.0.0.1:8080/')
print("Status Code: ",r.status_code)