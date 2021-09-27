# Instructions on Running this Docker Instance

```
docker build -t jupyter/custom-notebook .
docker run --name datanotebook -p 8888:8888 jupyter/custom-notebook
```

# Notes on the Docker File Linting

* [hadolint](https://github.com/hadolint/hadolint) is a tool used to lint docker images. When attempting to build docker images, if the docker file installation code is not written correctly, hadolint will prevent the file from being built.

## Example

If trying to run the following line within a Dockerfile...

```
RUN pip install pandas
```
...you will get the below hadolint warning...

```
[hadolint] warning: Pin versions in pip. Instead of `pip install <package>` use `pip install <package>==<version>` or `pip install --requirement <requirements file>`DL3013
[hadolint] warning: Avoid use of cache directory with pip. Use `pip install --no-cache-dir <package>`DL3042
```

So the correction would be to specify a precise version of [pandas](https://pypi.org/project/pandas/), directly from pypi.org like so:

```
RUN pip install --no-cache-dir pandas==1.3.3
```
When running with the above dependency installer, the lint test passes.
