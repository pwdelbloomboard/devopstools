#!/bin/bash

echo "Building image and tagging with name: jupyter/minimal-notebook/feastjupyter:ubuntu-22.04 "
docker build -t jupyter/minimal-notebook/feastjupyter:ubuntu-22.04 .