# Kustomize Value Inheretence

* This file was written to help explain what is available within Kustomize to be able to populate values in one yaml file, drawing from another yaml file.

## Word on Definitions - Inheretence, Composition

* When searching the Kustomize documentation and the broader web on the concept of kustomize inheretence, several lines of discussion come up, using the words, "inheretence," and "composition."

This is a brief definition of terms from a purely Computer Science standpoint, to help clarify what people might be talking about in various posts and threads.

### Inheretence

> In object-oriented programming, inheritance refers to the ability of an object to take on one or more characteristics from other classes of objects. The characteristics inherited are usually instance variables or member functions. An object that inherits these characteristics is known as a subclass.

### Composition

> Composition is one of the fundamental concepts in object-oriented programming. It describes a class that references one or more objects of other classes in instance variables. This allows you to model a has-a association between objects.

> You can find such relationships quite regularly in the real world. A car, for example, has an engine and modern coffee machines often have an integrated grinder and a brewing unit.

### Comparing Inheretence vs. Composition

While inheretence seems to describe something like this:

```
ObjectA <- ObjectB.function

or

ObjectA <- ObjectB.variable
```
* ObjectA is known as the, "subclass."

Whereas composition sounds more like:

```
ClassA.variable <- ClassB.ObjectX.variable
ClassA.ObjectY.variable <- ClassB.ObjectX.variable
```
* Basically,  whereas within inheretence, we can have objects taking on functions and variables from other types of objects...within composition, you can have classes reference objects in other classes.

* So the key difference is that composition deals with classes, essentially structured things that include objects and functions, whereas inheretence just seems to refer to the objects, not the classses.

### Tying This All Together

* Within Kustomize, we're ultimately dealing with YAML files, so it sounds reasonable that we should refer to these YAML files and the key/values within them as variables, which would imply we are really truly looking for inheretence when we say, "having a yaml file populate a value from another yaml file."
* That being said, it is of course always possible that people online confuse the concepts, or use them interchangeably.

## Reference Guides Online, StackOverflow Search, Etc.

* [Kustomize: Set attribute value from file or URI](https://stackoverflow.com/questions/58721575/kustomize-set-attribute-value-from-file-or-uri)

### Secrets

* The typical way this is achieved within Kustomize is with Secrets.
* That being said, there may be a reason not to use Secrets, e.g. perhaps there may be a convention that parameters are stored as environmental variables within Pods, rather than as Secrets. If this is the case, then another manner of achieveing inheretence may be necessary.

### Overlays

> ...overlay and patches (which is one or multiple files) that are kind of merged into your base file. A Patch overrides an attribute. With those two features you predefine some probable manifest-compositions and just combine them right before you apply them to your cluster.

Per [this documentation](https://github.com/kubernetes-sigs/kustomize#2-create-variants-using-overlays), there may be an overlays folder which includes values that replace those referenced in a base.

```

```