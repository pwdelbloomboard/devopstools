# About Feast

### About Feast - Kubeflow Perspective

* [Feast Feature Store Documentation - Kubeflow](https://www.kubeflow.org/docs/external-add-ons/feature-store/overview/)

> Feature Stores address key challenges:

* Feature sharing and re-use. Building engineering features can be very time consuming in an end-to-end ML system, yet teams develop features in silos.
* Serving features at scale - 
* Consistency between training and serving
* Point-in-time correctness
* Data quality and validation

### About Feast - Feast Documentation Perspective

* Feast is a way of serving features online
* Ensure consistency across training and serving
* Integrate existing tooling and infrastructure
* Standardizing data workflows across teams

#### What is a Feature Store in General

* [Blog Post on the Topic](https://feast.dev/blog/what-is-a-feature-store/)

> Data teams are starting to realize that operational machine learning requires solving data problems that extend far beyond the creation of data pipelines.
> Production data systems, whether for large scale analytics or real-time streaming, aren’t new. However, operational machine learning — ML-driven intelligence built into customer-facing applications — is new for most teams.
> The challenge of deploying machine learning to production for operational purposes (e.g. recommender systems, fraud detection, personalization, etc.) introduces new requirements for our data tools.

* "Feature" refers to an entire column in a dataset.
* "Feature Value," refers to a value within a cell.

> Feature stores make it easy to: 1. Productionize new features without extensive engineering support, 2. Automate feature computation, backfills, and logging 3. Share and reuse feature pipelines across teams 4. Track feature versions, lineage, and metadata 5. Achieve consistency between training and serving data 6. Monitor the health of feature pipelines in production

![](/img/feastarch.png)

![](/img/feastserve.png)

There are different types of feature serving that may happen:  Batch, Streaming and On-Demand:

![](/img/feasttransform.png)

* Note that the feature store may usually include some kind of Registry, a central interface for user interactions with the feature store.

##### Comparing Feature Stores:

> Feast is a great option if you already have transformation pipelines to compute your features, but need a great storage and serving layer to help you use them in production. Feast is GCP/AWS only today, but we’re working hard to make Feast available as a light-weight feature store for all environments. Stay tuned.

> Tecton is a feature-store-as-a-service. A big difference between Feast and Tecton is that Tecton supports transformations, so feature pipelines can be managed end-to-end within Tecton. Tecton is a managed offering, and a great feature store choice if you need production SLAs, hosting, advanced collaboration, managed transformations (batch/streaming/real-time), and/or enterprise capabilities.

##### Uber's Michealangelo Program

* [Michealangelo Blog Post](https://www.uber.com/en-CA/blog/michelangelo-machine-learning-platform/)

> Uber's internal ML as a service.

* Prior to Michelangelo, it was not possible to train models larger than what would fit on data scientists’ desktop machines, and there was neither a standard place to store the results of training experiments nor an easy way to compare one experiment to another. Most importantly, there was no established path to deploying a model into production–in most cases, the relevant engineering team had to create a custom serving container specific to the project at hand.

* Anti-patterns had emerged, as those described in [this paper](https://proceedings.neurips.cc/paper/2015/file/86df7dcfd896fcaf2674f757a2463eba-Paper.pdf), including:

* Glue Code - generic packages, massive amount of supporting code is written to get data into and out of general-purpose packages.
* Pipeline Jungles - arise within data preperation, and can evolve oganically.
* Dead Experimental Codepaths - experimental, "scratch" codepaths get created and abandoned.
* Abstraction Debt - there is a distinct lack of strong abstractions to support ML systems, nothing comes close to a relational database abstraction. It's ambiguous as to what the right interface should be to describe a stream of data, model or prediction. There is a lack of widely accepted abstractions.
* Common Smells - underlying problems in a component are found through, "smells," some general ideas to create smell tests for ML include: 1. Plain Old Data Type Smells, meaning there are just floats and integers, without objects that include information about whether the type is a multiplier or a decision-threshold, or how it should be consumed, etc. 2. Built with multiple languages. 3. Prototypes, e.g. relying on a prototyping environment indicates that the full-scale system is brittle.

> Use cases included - UberEATS estimated time for delivery method. Delivery time models predict how much time a meal will take to prepare and deliver before an order is issued. 

> UberEATS data scientists use gradient boosted decision tree regression models to predict this end-to-end delivery time. Features for the model include information from the request (e.g., time of day, delivery location), historical features (e.g. average meal prep time for the last seven days), and near-realtime calculated features (e.g., average meal prep time for the last one hour). Models are deployed across Uber’s data centers to Michelangelo model serving containers and are invoked via network requests by the UberEATS microservices. These predictions are displayed to UberEATS customers prior to ordering from a restaurant and as their meal is being prepared and delivered.

There was a six-step workflow:

> Manage data, Train models, Evaluate models, Deploy models, Make predictions, Monitor predictions

* Finding good features is often the hardest part of ML, and is often the most costly part.

>  we added a layer of data management, feature store that allows teams to share, discover, and use a highly curated set of features for their machine learning problems.

> there is substantial value in enabling teams to share features between their own projects and for teams in different organizations to share features with each other.

> It allows users to easily add features they have built into a shared feature store, requiring only a small amount of extra metadata (owner, description, SLA, etc.) on top of what would be required for a feature generated for private, project-specific usage.
Once features are in the Feature Store, they are very easy to consume, both online and offline, by referencing a feature’s simple canonical name in the model configuration. Equipped with this information, the system handles joining in the correct HDFS data sets for model training or batch prediction and fetching the right value from Cassandra for online predictions.

> At the moment, we have approximately 10,000 features in Feature Store that are used to accelerate machine learning projects, and teams across the company are adding new ones all the time. Features in the Feature Store are automatically calculated and updated daily.

From there on, Training Models, Evaluating Models, Deploying Models was easier.

### Setting Up Feast - Kubeflow Perspective

* [Feature Store Getting Started](https://www.kubeflow.org/docs/external-add-ons/feature-store/getting-started/)

#### Installing Feast to Development Environment

The standard way to install Feast is via K8s:

* [Running Feast in K8s](https://docs.feast.dev/how-to-guides/running-feast-in-production)

* It appears that Feast gets installed with Helm.

[Feast Python / Go Feature Server Helm Charts](https://github.com/feast-dev/feast/tree/master/infra/charts/feast-feature-server)

* There was an issue started requesting to convert this to a kustomize deployment: [Feast Issue](https://github.com/feast-dev/feast/issues/258)
* The official answer was:

> but we have decided to not maintain a Kustomize installation as part of Feast. We will use Helm as the primary way to install into Kubernetes.

That being said, AWS and other services offer their own Feast services via Redis, and it is possible to customize Feast in such a way that resources get integrated with notebooks.


#### Installing via Docker to a Jupyter Notebook

* Feast may be installed to a Jupyter Notebook.