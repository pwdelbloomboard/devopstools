# About Jupyter Notebook with GraphViz Installed

* This is a Docker image which includes a Jupyter Notebook that allows GraphViz to be used within the notebook.
* The Docker Container may be built using buildimageandtag.sh.
* Graphviz allows a controllable way to create software and business diagrams using source code.
* The volumebindmount folder contains a dev mode container with souce code for the Graphviz charts.

* [More on Graphiz](https://graphviz.readthedocs.io/en/stable/manual.html)

* [Even More on Graphviz](https://graphviz.org/doc/info/lang.html)

### Usage

* Within the graphviz folder, do the following to build and tag a GraphViz Jupyter Notebook image:

```
chmod +x buildimageandtag.sh
./buildimageandtag
```
* Once the image is built, you can launch the container using Docker Compose, in the same folder:

```
docker-compose up -d
```
* This launches the container in detached mode, copying code from ./volumebindmount, and allowing any work done while the container is in operation to be saved into the same folder.
* After launching the container, on your console you should see a link to the notebook which looks like the following:

```
http://127.0.0.1:8888/lab?token=58964ba467e08b3b8cb9705ac9ec873cc0a39e361ba9bf20
```

* Open that in your browser, and it will open up the Jupyter Lab notebook.
* Within that notebook, you can execute GraphViz code by pressing the, "play/execute" button as shown below.
* You can view the GraphViz chart within the notebook as shown below.

![](/graphviz/img/graphvizcode.png)

![](/graphviz/img/graphvizimage.png)

* To shut down the notebook once you are done, within your browser go to File >> Shut Down.
* This will shut down the notebook gracefully, and any changes you had made to the notebook will be saved within /volumebindmount.

## Examples

* It's much easier to draw of existing examples to build charts rather than work completely from scratch.  Here are some helpful examples:

### GraphViz Playgrounds:

* https://sketchviz.com/new

* http://magjac.com/graphviz-visual-editor/

* https://dreampuf.github.io/GraphvizOnline/


### GraphViz Examples

* https://renenyffenegger.ch/notes/tools/Graphviz/examples/index
* https://renenyffenegger.ch/notes/tools/Graphviz/attributes/_color/index
* https://graphs.grevian.org/example
* https://towardsdatascience.com/graph-visualisation-basics-with-python-part-iii-directed-graphs-with-graphviz-50116fb0d670 