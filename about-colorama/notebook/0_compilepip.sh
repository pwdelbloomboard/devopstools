#!/bin/bash

pip-compile --allow-unsafe --annotation-style=line --generate-hashes --output-file=requirements.txt requirements.in