# Official Python Package UserGuide

* [Python Package Userguide](https://packaging.python.org/en/latest/)

## Gitlab Documentation

* [Gitlab Documentation](https://docs.gitlab.com/ee/user/packages/pypi_repository/)

## Private Python Repository

### Starting Up The Project

* As with other projects within devops tools, we'll start up a dev-mode-container and Docker Project to keep everything isolated.
* The crux of the package will simply be a, "hello world" python script.  Our file directory looks like the following:

```
/volumebindmount
/tests
/app
    greet.py
    __init__.py
```

* We attempt install wheel but see that in our base image it's already installed:

```
pip install wheel
Requirement already satisfied: wheel in /usr/local/lib/python3.9/site-packages (0.38.4)
```

* We could optionally include this in our requirements file to lock the version.
* After verifying that wheel is installed we do:

```
python app/greet.py bdist_wheel
```

* If we do the above, nothing happens. The code for bdist_wheel is found here: [github/pypa/wheel](https://github.com/pypa/wheel). 

According to their documentation, within the [user guide](https://wheel.readthedocs.io/en/stable/user_guide.html#building-wheels), we need to build with an option wheel:

```
python -m pip install build
python -m build --wheel
```

However, we don't have the, "build" package installed by default within python, so we can find that [here](https://pypi.org/project/build/).

After installing build, we get:

```
python -m build --wheel
ERROR Source /home/volumebindmount does not appear to be a Python project: no pyproject.toml or setup.py
```

* So obviously in order to set up an appropriate python project, we need a pyproject.toml or setup.py!
* In fact, attempting to install ```python app/greet.py bdist_wheel``` was fallacious, because evidently, setup.py needs to be a special format, if that is indeed the format to be used for packaging up a package.

Back to, [the main packaging.python.org instructions](https://packaging.python.org/en/latest/guides/distributing-packages-using-setuptools/):

* There must be a setup.py file which is the root of the project directory.
* There's another file, setup.cfg is used for additional commands:

Re: the documentation:

> It’s the file where various aspects of your project are configured. The primary feature of setup.py is that it contains a global setup() function. The keyword arguments to this function are how specific details of your project are defined. The most relevant arguments are explained in the section below.
> It’s the command line interface for running various commands that relate to packaging tasks. To get a listing of available commands, run python setup.py --help-commands.

* The arguments which [are shown here](https://packaging.python.org/en/latest/guides/distributing-packages-using-setuptools/#setup-args).

* Evidently we can also use a pyproject.toml to represent the project as well, and in fact, the use of setup.py is discouraged in favor of pyproject.toml.

 > pyproject.toml is the specified file format of PEP 518.  Modern Python packages can contain a pyproject.toml file, first introduced in PEP 518 and later expanded in PEP 517, PEP 621 and PEP 660. This file contains build system requirements and information, which are used by pip to build the package.
 
* More information about pyproject.toml can be found [here](https://pip.pypa.io/en/stable/reference/build-system/pyproject-toml/).
* PEP is the python enhancement protocols. More information specifically about PEP518 can be found [here](https://peps.python.org/pep-0518/), which was established in 2016.

> This PEP specifies how Python software packages should specify what build dependencies they have in order to execute their chosen build system. As part of this specification, a new configuration file is introduced for software packages to use to specify their build dependencies (with the expectation that the same configuration file will be used for future configuration details).

* In fact, any information about any PEP can be found [here](https://peps.python.org/).

* So that being said, how do we structure a pyproject.toml?

#### Detail about pyproject.toml

* [Overall Specification - Starting with Build Time Dependencies](https://pip.pypa.io/en/stable/reference/build-system/pyproject-toml/#build-time-dependencies)

* We can set up build time dependencies as we hae done below, including setuptools and cython.
* Looking further into what exactly toml is, we [see this blog post](https://martin-thoma.com/pyproject-toml/)
* Which points toward this [packaging tutorial](https://packaging.python.org/en/latest/tutorials/packaging-projects/)
* These above resources note that:

> The pyproject.toml file allows package creators to define the build system as a dependency as well as a projects metadata. Also, other kinds of meta-data and the install requirements can be defined in it.

* Our basic starter setup is:

```
[build-system]
requires = ["setuptools>=61.0","cython ~= 0.29.0"]

build-backend = "setuptools.build_meta"

[project]
name = "app"
version = "0.0.1"
authors = [
    { name="Patrick Delaney", email="patrick@fakeemail.com" },
]
description = "Sample package application."
requires-python = ">=3.7"
classifiers = [
   "Programming Language :: Python :: 3",
   "Operating System :: OS Independent",
]

[tool.setuptools.packages]
find = {}

```

* Adding onto this further we can provide further metadata, we can look at some of the definitions provided within the tutorials above. Specifically, one definition, "classifiers," seems interesting:

> classifiers gives the index and pip some additional metadata about your package. In this case, the package is only compatible with Python 3, is licensed under the MIT license, and is OS-independent. You should always include at least which version(s) of Python your package works on, which license your package is available under, and which operating systems your package will work on. For a complete list of classifiers, see https://pypi.org/classifiers/.

* Looking further into the [classifiers](https://pypi.org/classifiers/) documentation, we see that:

> Each project's maintainers provide PyPI with a list of "Trove classifiers" to categorize each release, describing who it's for, what systems it can run on, and how mature it is.

* Looking at that list, we see that there are a huge variety of options, including one for Production/Stable:

> Development Status :: 5 - Production/Stable

There are also topics, huge varieties of programming languages, natural languages and even what GPU enviroment, such as CUDA:

> Environment :: GPU :: NVIDIA CUDA


#### Moving on and Generating Distribution Archives

[Moving on to this part of the docs](https://packaging.python.org/en/latest/tutorials/packaging-projects/#generating-distribution-archives):


1. Make sure we have the latest version of build installed (this should be done in the Dockerfile):

```
python3 -m pip install --upgrade build
```

2. Run the following where the pyproject.toml is located:

```
python3 -m build
```

* We then get:

```
python3 -m build
* Creating venv isolated environment...
* Installing packages in isolated environment... (setuptools>=61.0)
* Getting build dependencies for sdist...
/tmp/build-env-gwy93yc9/lib/python3.9/site-packages/setuptools/config/pyprojecttoml.py:108: _BetaConfiguration: Support for `[tool.setuptools]` in `pyproject.toml` is still *beta*.
  warnings.warn(msg, _BetaConfiguration)
running egg_info
creating app.egg-info
writing app.egg-info/PKG-INFO
writing dependency_links to app.egg-info/dependency_links.txt
writing top-level names to app.egg-info/top_level.txt
writing manifest file 'app.egg-info/SOURCES.txt'
reading manifest file 'app.egg-info/SOURCES.txt'
writing manifest file 'app.egg-info/SOURCES.txt'
* Building sdist...
/tmp/build-env-gwy93yc9/lib/python3.9/site-packages/setuptools/config/pyprojecttoml.py:108: _BetaConfiguration: Support for `[tool.setuptools]` in `pyproject.toml` is still *beta*.
  warnings.warn(msg, _BetaConfiguration)
running sdist
running egg_info
writing app.egg-info/PKG-INFO
writing dependency_links to app.egg-info/dependency_links.txt
writing top-level names to app.egg-info/top_level.txt
reading manifest file 'app.egg-info/SOURCES.txt'
writing manifest file 'app.egg-info/SOURCES.txt'
running check
creating app-0.0.1
creating app-0.0.1/app.egg-info
creating app-0.0.1/src
copying files to app-0.0.1...
copying README.md -> app-0.0.1
copying pyproject.toml -> app-0.0.1
copying app.egg-info/PKG-INFO -> app-0.0.1/app.egg-info
copying app.egg-info/SOURCES.txt -> app-0.0.1/app.egg-info
copying app.egg-info/dependency_links.txt -> app-0.0.1/app.egg-info
copying app.egg-info/top_level.txt -> app-0.0.1/app.egg-info
copying src/__init__.py -> app-0.0.1/src
copying src/greet.py -> app-0.0.1/src
Writing app-0.0.1/setup.cfg
Creating tar archive
removing 'app-0.0.1' (and everything under it)
* Building wheel from sdist
* Creating venv isolated environment...
* Installing packages in isolated environment... (setuptools>=61.0)
* Getting build dependencies for wheel...
/tmp/build-env-iwzac3b3/lib/python3.9/site-packages/setuptools/config/pyprojecttoml.py:108: _BetaConfiguration: Support for `[tool.setuptools]` in `pyproject.toml` is still *beta*.
  warnings.warn(msg, _BetaConfiguration)
running egg_info
writing app.egg-info/PKG-INFO
writing dependency_links to app.egg-info/dependency_links.txt
writing top-level names to app.egg-info/top_level.txt
reading manifest file 'app.egg-info/SOURCES.txt'
writing manifest file 'app.egg-info/SOURCES.txt'
* Installing packages in isolated environment... (wheel)
* Building wheel...
/tmp/build-env-iwzac3b3/lib/python3.9/site-packages/setuptools/config/pyprojecttoml.py:108: _BetaConfiguration: Support for `[tool.setuptools]` in `pyproject.toml` is still *beta*.
  warnings.warn(msg, _BetaConfiguration)
running bdist_wheel
running build
running build_py
creating build
creating build/lib
creating build/lib/src
copying src/__init__.py -> build/lib/src
copying src/greet.py -> build/lib/src
running egg_info
writing app.egg-info/PKG-INFO
writing dependency_links to app.egg-info/dependency_links.txt
writing top-level names to app.egg-info/top_level.txt
reading manifest file 'app.egg-info/SOURCES.txt'
writing manifest file 'app.egg-info/SOURCES.txt'
installing to build/bdist.linux-x86_64/wheel
running install
running install_lib
creating build/bdist.linux-x86_64
creating build/bdist.linux-x86_64/wheel
creating build/bdist.linux-x86_64/wheel/src
copying build/lib/src/__init__.py -> build/bdist.linux-x86_64/wheel/src
copying build/lib/src/greet.py -> build/bdist.linux-x86_64/wheel/src
running install_egg_info
Copying app.egg-info to build/bdist.linux-x86_64/wheel/app-0.0.1-py3.9.egg-info
running install_scripts
creating build/bdist.linux-x86_64/wheel/app-0.0.1.dist-info/WHEEL
creating '/home/volumebindmount/app/dist/tmprzy4vtmo/app-0.0.1-py3-none-any.whl' and adding 'build/bdist.linux-x86_64/wheel' to it
adding 'src/__init__.py'
adding 'src/greet.py'
adding 'app-0.0.1.dist-info/METADATA'
adding 'app-0.0.1.dist-info/WHEEL'
adding 'app-0.0.1.dist-info/top_level.txt'
adding 'app-0.0.1.dist-info/RECORD'
removing build/bdist.linux-x86_64/wheel
Successfully built app-0.0.1.tar.gz and app-0.0.1-py3-none-any.whl
```
* So now that we have successfully created a wheel, we can look to make sure we have it available. Sure enough:

```
/home/volumebindmount/app/dist# ls
app-0.0.1-py3-none-any.whl  app-0.0.1.tar.gz
```
> The tar.gz file is a source distribution whereas the .whl file is a built distribution.

> A built distribution is a format containing files and metadata that only need to be moved to the correct location on the target system, to be installed.

* Whereas:

> A source distribution provides metadata and the essential source files needed for installing by a tool like pip, or for generating a Built Distribution.

* So by creating the whl file, we basically have packaged up our package.

#### Uploading the Distribution Archives

* You can start out by uploading a sample distribution to [test.pypi.org](https://test.pypi.org/account/register/).

* Once registered and having set up a token, you can use twine to upload a distribution package.

```
python3 -m pip install --upgrade twine
python3 -m twine upload --repository testpypi dist/*
```

* In attempting the above, you may get:

```
HTTPError: 403 Forbidden from https://test.pypi.org/legacy/
```

* This may be due to how the token and api key is being set up within your system.  The recommended way to set up one's $HOME/.pypirc file is like this:

```
[testpypi]
  username = __token__
  password = (the actual password)
```

* Note - be careful that none of this information gets mounted anywhere within the bind mount, e.g. that it gets destroyed along with the container and does not get mounted as a part of the bind mount, as it is a secret.
* If you get an invalid auth, you might be redirected here: https://test.pypi.org/help/#invalid-auth ... I was redirected there and found that it's necessary to literally use username = __token__ as in, two underscores before and after, not the actual token name.
* However even after doing that, I still got an error:

```
The user 'pwdelbloomboard' isn't allowed to upload to project 'app'. See https://test.pypi.org/help/#project-name for more information.
```

* This url leads to the message:

> Your publishing tool may return an error that your new project can't be created with your desired name, despite no evidence of a project or release of the same name on PyPI. Currently, there are four primary reasons this may occur: The project name conflicts with a Python Standard Library module from any major version from 2.5 to present. The project name is too similar to an existing project and may be confusable. The project name has been explicitly prohibited by the PyPI administrators. For example, pip install requirements.txt is a common typo for pip install -r requirements.txt, and should not surprise the user with a malicious package. The project name has been registered by another user, but no releases have been created.See How do I claim an abandoned or previously registered project name?

* Basically, it's quite obvious that there are likely already projects called, "app" out there!  So, we can put a random tag at the end of our app generated from OpenSSL and rebuild:

```
[project]
name = "banana-app-NZCkyBK1zTm4z96s"
```

* After deleting our old files, within our */app directory we have:

```
README.md  banana_app_NZCkyBK1zTm4z96s.egg-info  dist  pyproject.toml  src
```

* Which we can now attempt to upload with:

```
python3 -m twine upload --repository testpypi dist/*
Uploading distributions to https://test.pypi.org/legacy/
Uploading banana_app_NZCkyBK1zTm4z96s-0.0.1-py3-none-any.whl
100% ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 5.3/5.3 kB • 00:00 • ?
Uploading banana-app-NZCkyBK1zTm4z96s-0.0.1.tar.gz
100% ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 4.9/4.9 kB • 00:00 • ?

View at:
https://test.pypi.org/project/banana-app-NZCkyBK1zTm4z96s/0.0.1/
```
* Looking at this website, we can see that we have the package ready to go and available:

![](/about-pythonpackage/pypackage.png)

#### Using Pip To Install a Newly Uploaded Package

```
pip install -i https://test.pypi.org/simple/ banana-app-NZCkyBK1zTm4z96s==0.0.1
```

* This should lead to the message:

```
pip install -i https://test.pypi.org/simple/ banana-app-NZCkyBK1zTm4z96s==0.0.1
Looking in indexes: https://test.pypi.org/simple/
Collecting banana-app-NZCkyBK1zTm4z96s==0.0.1
  Downloading https://test-files.pythonhosted.org/packages/5a/ec/dcae7f471a689926258afac77d783aa71f93a695082615f80072534196a9/banana_app_NZCkyBK1zTm4z96s-0.0.1-py3-none-any.whl (1.9 kB)
Installing collected packages: banana-app-NZCkyBK1zTm4z96s
Successfully installed banana-app-NZCkyBK1zTm4z96s-0.0.1
```

* We may then import parts of the package based upon the folder structure.  However, when we try to do the following, nothing works, because our folder structure is not set up properly.

```
>>> from banana-app-NZCkyBK1zTm4z96s import greet
  File "<stdin>", line 1
    from banana-app-NZCkyBK1zTm4z96s import greet
               ^
SyntaxError: invalid syntax
```

* Going back above, to look at how the folder structure should have originally been built we see:

> __init__.py is required to import the directory as a package, and should be empty.
> example.py is an example of a module within the package that could contain the logic (functions, classes, constants, etc.) of your package.

* So, we re-organize our directory structure:

```
app/
└── src/
    └── addone/
    |   ├── __init__.py
    |   └── addone.py
    └── greet/
        ├── __init__.py
        └── greet.py        
```

* Within greet.py we included a ```if __name__ == '__main__':``` clause to test out the behavior specification, "runs blocks of code only when our Python script is being executed directly from a use", whereas within addone.py we just created a function.
* We left both __init__.py open per specification.
* We can attempt to build and upload the package on pypi once again, without changing the version with:


```
python3 -m build
python3 -m twine upload --repository testpypi dist/*
```

* Note that the version number MUST be iterated as from an error redirect url we found [here](https://test.pypi.org/help/#file-name-reuse) states:

> PyPI will return these errors for one of these reasons: Filename has been used and file exists Filename has been used but file no longer exists A file with the exact same content exists PyPI does not allow for a filename to be reused, even once a project has been deleted and recreated. A distribution filename on PyPI consists of the combination of project name, version number, and distribution type.

* So attempting to iterate the version number to 0.0.2, and completely removing the 0.0.1 version files from the dist folder we get:

```
python3 -m twine upload --repository testpypi dist/*
Uploading distributions to https://test.pypi.org/legacy/
Uploading banana_app_NZCkyBK1zTm4z96s-0.0.2-py3-none-any.whl
100% ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 5.6/5.6 kB • 00:00 • ?
Uploading banana-app-NZCkyBK1zTm4z96s-0.0.2.tar.gz
100% ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 5.1/5.1 kB • 00:00 • ?

View at:
https://test.pypi.org/project/banana-app-NZCkyBK1zTm4z96s/0.0.2/
```

* Which, we can get with pip for the correct version via the following, which would only work after visiting the actual package site [here](https://test.pypi.org/project/banana-app-NZCkyBK1zTm4z96s/0.0.2/):

```
pip install -i https://test.pypi.org/simple/ banana-app-NZCkyBK1zTm4z96s==0.0.2
Looking in indexes: https://test.pypi.org/simple/
Collecting banana-app-NZCkyBK1zTm4z96s==0.0.2
  Downloading https://test-files.pythonhosted.org/packages/10/fd/9bc3f7f6321378dede8d88e55f1a9bde739524a20a311f3224436f4fc66c/banana_app_NZCkyBK1zTm4z96s-0.0.2-py3-none-any.whl (2.2 kB)
Installing collected packages: banana-app-NZCkyBK1zTm4z96s
  Attempting uninstall: banana-app-NZCkyBK1zTm4z96s
    Found existing installation: banana-app-NZCkyBK1zTm4z96s 0.0.1
    Uninstalling banana-app-NZCkyBK1zTm4z96s-0.0.1:
      Successfully uninstalled banana-app-NZCkyBK1zTm4z96s-0.0.1
Successfully installed banana-app-NZCkyBK1zTm4z96s-0.0.2
```
* So as can be seen above, version 0.0.1 was uninstalled before 0.0.2 was installed.  Attempting to import the package we get:

```
>>> from banana-app-NZCkyBK1zTm4z96s import greet
  File "<stdin>", line 1
    from banana-app-NZCkyBK1zTm4z96s import greet
               ^
SyntaxError: invalid syntax
```
* If we look at pip list we see:

```
root@9bcd51c6599d:/home/volumebindmount/app# pip list
Package                     Version
--------------------------- ---------
banana-app-NZCkyBK1zTm4z96s 0.0.2
```

* However within the environment we get:

```
>>> from banana_app_NZCkyBK1zTm4z96s import addone
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ModuleNotFoundError: No module named 'banana_app_NZCkyBK1zTm4z96s'
```
* This might be a permissions issue.  We can find the location of the package with:

```
root@9bcd51c6599d:/home/volumebindmount/app# pip show banana-app-NZCkyBK1zTm4z96s
Name: banana-app-NZCkyBK1zTm4z96s
Version: 0.0.2
Summary: Sample package application.
Home-page:
Author:
Author-email: Patrick Delaney <patrick@fakeemail.com>
License:
Location: /usr/local/lib/python3.9/site-packages
Requires:
Required-by:
```
* Showing the permissions of the file:

```
root@9bcd51c6599d:/home/volumebindmount/app# ls -l /usr/local/lib/python3.9/site-packages | grep banana
drwxr-xr-x  2 root root   4096 Nov 22 20:52 banana_app_NZCkyBK1zTm4z96s-0.0.2.dist-info
```

* Actually this is a directory, looking in the file itself:

```
ls -l /usr/local/lib/python3.9/site-packages/banana_app_NZCkyBK1zTm4z96s-0.0.2.dist-info
total 20
-rw-r--r-- 1 root root    4 Nov 22 20:52 INSTALLER
-rw-r--r-- 1 root root  606 Nov 22 20:52 METADATA
-rw-r--r-- 1 root root 1093 Nov 22 20:52 RECORD
-rw-r--r-- 1 root root    0 Nov 22 20:52 REQUESTED
-rw-r--r-- 1 root root   92 Nov 22 20:52 WHEEL
-rw-r--r-- 1 root root    4 Nov 22 20:52 top_level.txt
```

* We can see that we only have rw permissions on all of these files, no execute permission.

```
from banana_app_NZCkyBK1zTm4z96s import addone
```


* Attempting to hammer everything into shape throughout the whole directory (sudo 777 is always very bad idea ):

```
chmod 777 /usr/local/lib/python3.9/site-packages
```

* However, this didn't work either, because we have to chmod -R, or recursively across the entire directory.

```
root@9bcd51c6599d:/home/volumebindmount/app# chmod 755 -R /usr/local/lib/python3.9/site-packages/banana_app_NZCkyBK1zTm4z96s-0.0.2.dist-info
root@9bcd51c6599d:/home/volumebindmount/app# ls -l /usr/local/lib/python3.9/site-packages/banana_app_NZCkyBK1zTm4z96s-0.0.2.dist-info
total 20
-rwxr-xr-x 1 root root    4 Nov 22 20:52 INSTALLER
-rwxr-xr-x 1 root root  606 Nov 22 20:52 METADATA
-rwxr-xr-x 1 root root 1093 Nov 22 20:52 RECORD
-rwxr-xr-x 1 root root    0 Nov 22 20:52 REQUESTED
-rwxr-xr-x 1 root root   92 Nov 22 20:52 WHEEL
-rwxr-xr-x 1 root root    4 Nov 22 20:52 top_level.txt
```

* Now we can see that the file is actually executable at the 755 level.  However, none of this allows us to import hte module.

[According to this Stackoverflow article](https://stackoverflow.com/questions/61665331/created-a-python-package-in-testpypi-but-i-cant-import-it), evidently, we need a subdirectory with the same name as the entire application that we are looking to create if we want to import via the top level name specified.

* So, going back to our python terminal, since we did have a top level directory named, "src" - which is not a good name for a top level directory, because it's a common name, we can do:

```
>>> from src import addone
>>> from src import greet
>>> import src
>>>
```
* ...without error.  However, when we attempt to invoke the functions within them we get:

```
>>> from src import addone
>>> addone.add_one(2)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: module 'src.addone' has no attribute 'add_one'
```
* This is in part due to how python package structuring works. The way to actually import a specific part of a package and then invoke its function is the following, as discussed [here at this stackoverflow article](https://stackoverflow.com/a/4871671).

```
>>> from src.addone import addone
>>> addone.add_one(2)
Adding:  2  +1:
3

>>> from src.greet import greet
>>> greet.SayHello()
Hello from my package
```

#### Uninstalling and Reinstalling:

```
root@9bcd51c6599d:/# pip uninstall banana-app-NZCkyBK1zTm4z96s
Found existing installation: banana-app-NZCkyBK1zTm4z96s 0.0.2
Uninstalling banana-app-NZCkyBK1zTm4z96s-0.0.2:
  Would remove:
    /usr/local/lib/python3.9/site-packages/banana_app_NZCkyBK1zTm4z96s-0.0.2.dist-info/*
    /usr/local/lib/python3.9/site-packages/src/addone/*
    /usr/local/lib/python3.9/site-packages/src/greet/*
Proceed (Y/n)? y
  Successfully uninstalled banana-app-NZCkyBK1zTm4z96s-0.0.2
```
* Re-installing:

```
pip install -i https://test.pypi.org/simple/ banana-app-NZCkyBK1zTm4z96s==0.0.2
```
* Ensuring still works:

```
>>> from src.addone import addone
>>> addone.add_one(2)
Adding:  2  +1:
3
```
* The above shows that the package can be easily uninstalled and re-installed.

### The Upshot

* All of the previous should show from start to finish how to successfully import a custom package at one's chosen repo, which may be Github, Gitlab or a myriad of other locations.
* In order to duplicate this in a private location, you must read the documentation for the chosen service.


# Tutorials

* https://dev.to/christo22694524/installing-private-python-packages-in-docker-images-1hgm
* https://packaging.python.org/en/latest/tutorials/packaging-projects/