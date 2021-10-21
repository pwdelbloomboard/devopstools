### Definition of Pipelines

Pipelines Comprise of:
* Jobs, which describe what we do - for example, compile code.
* Stages, which define when to run jobs.

If all jobs within a stage succeed, the pipeline moves to the next stage.

If any job in a stage fails, the next stage usually does not execute and the pipeline ends early.

GitLab CI is configured via .gitlab-ci.yml file in a repo and can reference configs in other repos as well.

* There can be a CI Repo where standard job configurations and scripts are kept.
* There can also be a seperately created, "DevOps" mono-repo in which ALL DevOps scripts and files are kept, much like having a complete codebase monorepo for multiple microapplications.

### Goals and Objectives of Pipelines

* Run a single pipeline per commit regardlesss of how many apps are changed.
* Set the pipeline job status after all apps have been processed.
* Clearly mark output from pipeline jobs with the app it is from.
* Configure job using evnironment variables in GitLab.
* Reuse as much code as possible from pre-Yosemite pipelines.

### Pipeline Overviews

* Identify which apps have changes and store list in env variable for future jobs.
* Run pre-commit checks.
* Run tests for each application.
* Run code quality checks for each application.
* If a .pre-commit-config.yaml exists for an app that has changed, it will run as well.
* E.g. if a particular pair of apps both

### ci repo

* scripts
* test.sh
* gitlab-ci

#### .pre-commit-config.yaml

Can check various rules or tests...:

* Static files sizes
* Video file sizes
* Code no referencing large files under src
* Setting HEAD at various points along MASTER

### Gitlab Status

* Gitlab shows pipeline status on a particular menu within a pipeline under [REPO]/-/pipelines/[PIPELINE NUMBER]

* .pre
* pre Build
* Build
* Deploy
* .post

### Build Chain

Jobs are often chained together. In its simplest form, this is done by triggering a second job after a first job is finished - when several of these are done one after another, this is called a, "Build Chain."  This is the more general term.

* In Jenkins terminology, the first build in a chain is called the upstream build, and the second one of subsequent builds is called the downstream build.

* While this way of chaining builds is often sufficient, there will most likely be a need for greater control of the build chain eventually. Such a build chain is often called a pipeline or workflow.

#### Pipeline vs Workflow

Basically workflows are subsets of pipelines. The pipeline is meant to mean the entire chain of workflows. Workflows and pipelines are made up of buildchains, which is really just the way to describe the links in the chain. Workflows are just finer-grained ways of seperating out parts of the entire pipeline.

*  A pipeline is the full set of processes that run when you trigger work on your projects. Pipelines encompass workflows.

* workflows are smaller and coordinate multiple jobs. Workflows allow you to run and troubleshoot jobs separately so you can see failed builds in real-time. If one job in a workflow fails, you can re-run that job alone instead of re-running the entire set.

Within data science, pipelines may mean, "fully automatic," whereas workflow may mean, "human effort involved," - so there is a slightly different definition.

#### Parallel vs Serial

* Pipelines can be designed to run with workflows that are either parallel or serial to one another. Kubernetes pods, compute units, docker containers may all run in parallel to one another in order to save time and go faster, while another step (e.g. workflow) in the pipeline might involve pushing static assets to a storage (such as AWS S3).
* Parallelism is designed to save time.
* Serialism (subsequent workflows) are designed to prevent failover or simply because event A must occur prior to event B.

Hypothetically if everything could be done in parallel, if you have the budget, this would be the best way to do things...but serialism is used because of necesity.


# Resources

* [Networking and Servers - Job Chaining](https://subscription.packtpub.com/book/networking-and-servers/9781785882876/5/ch05lvl1sec76/job-chaining-and-build-pipelines)