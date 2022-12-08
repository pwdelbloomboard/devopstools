#!/bin/bash

echo "Building image and tagging with name: jupyter/scipy-notebook:python-3.8-with-deps"
docker build -t jupyter/scipy-notebook:python-3.8-with-deps .