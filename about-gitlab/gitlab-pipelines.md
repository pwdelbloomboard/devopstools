# Gitlab Pipelines

[Gitlab Pipelines](https://docs.gitlab.com/ee/ci/pipelines/)

## Abstract

* Pipelines are considered the top-level component of continuous integration, delivery and deployment.
* In this file, we'll go through a hypothetical scenario where we automatically upgrade a package and inject it into a build for a given application using pipelines.

* Pipelines Include:

1. Jobs, which define what is being done.
2. Stages, which define when it's done.

* Types of Pipelines:

Pipelines can be configured in many different ways:

* Basic pipelines run everything in each stage concurrently, followed by the next stage.
* Directed Acyclic Graph Pipeline (DAG) pipelines are based on relationships between jobs and can run more quickly than basic pipelines.
* Merge request pipelines run for merge requests only (rather than for every commit).
* Merged results pipelines are merge request pipelines that act as though the changes from the source branch have already been merged into the target branch.
* Merge trains use merged results pipelines to queue merges one after the other.
* Parent-child pipelines break down complex pipelines into one parent pipeline that can trigger multiple child sub-pipelines, which all run in the same project and with the same SHA. This pipeline architecture is commonly used for mono-repos.
* Multi-project pipelines combine pipelines for different projects together.

### Publishing Package to Package Registry and Installing

The objective of publishing a Package to a Package registry would be the following:

1. When code from a particular package X gets merged into master or a specified branch:
2. Publish the version of package X to the Gitlab package registry.

Then ideally:

3. Build dockerfiles that installs the latest version of package X
4. Trigger a job to build images which get built and injected into the latest version of a software platform.

#### For More Information About Python Packages

* [Python Packages Article](https://github.com/pwdelbloomboard/devopstools/blob/main/about-pythonpackage/pythonpackage.md)

#### Step 1 - Monitor When Code Merged to Master

* This appears to be a, "Merge Results Pipeline" - note that this is a Premium Feature within Gitlab:

> Merged results pipelines are merge request pipelines that act as though the changes from the source branch have already been merged into the target branch.

* Notably, there are, "Merge Request Pipelines," which run every time a change is committed to a branch. This may alternatively be known as a, "branch pipeline."
* More specifically, zooming in further is a, "Merge Results Pipeline," which [occurs after a merge is completed](https://docs.gitlab.com/ee/ci/pipelines/merged_results_pipelines.html).

##### Prerequisites

> Your project’s CI/CD configuration file must be configured to run jobs in merge request pipelines.
> Your repository must be a GitLab repository, not an external repository.
> You must not be using fast forward merges. An issue exists to change this behavior.

* The following must be done:

1. Enable Merged Results Pipelines. Navigate to: Settings/Merge Requests / Merge Options / Enable Merged Results Pipelines.

* Assuming Premium is activated and you activate that option for the designated Project, it should work.

###### The .gitlab-ci.yml File

* [Documentation About This Here](https://docs.gitlab.com/ee/ci/yaml/gitlab_ci_yaml.html)

> When you add a .gitlab-ci.yml file to your repository, GitLab detects it and an application called GitLab Runner runs the scripts defined in the jobs.

Sample .gitlab-ci.yml file:

```
stages:
  - build
  - test

build-code-job:
  stage: build
  script:
    - echo "Check the ruby version, then build some Ruby project files:"
    - ruby -v
    - rake

test-code-job1:
  stage: test
  script:
    - echo "If the files are built successfully, test some files with one command:"
    - rake test1

test-code-job2:
  stage: test
  script:
    - echo "If the files are built successfully, test other files with a different command:"
    - rake test2
```

##### .gitlab-ci.yml keyword reference

* There's a keyword reference for [.gitlab-ci.yml](https://docs.gitlab.com/ee/ci/yaml/).
* This documentation applies to all plans, but keep in mind, premium is required for branch results merge.

* Looking at some various examples of steps we may need:

* "include" is the selector for the project name, file name, branch name.  This is a global keyword along with, "default" which signifies default values, "stages" which names the order of the pipeline stages, "variables" which defines the variables for the job in the pipeline and "workflow," which controls what kind of pipeline to run.
* [include](https://docs.gitlab.com/ee/ci/yaml/#include) has specific possible subjey inputs:

```
include:local
include:project
include:remote
include:template
```

* So for example [include:project](https://docs.gitlab.com/ee/ci/yaml/#includeproject) is documented in its own right.

```
include:
  - project: 'my-group/my-project'
    file: '/templates/.gitlab-ci-template.yml'
    ref: master
```
* We can define our own instruction set, such as, "Run precommit checks" which utilizes keyword, ["extends"](https://docs.gitlab.com/ee/ci/yaml/#extends) which creates Configuration entries that a job inherits from.
* extends an alternative to YAML anchors and is a little more flexible and readable.
* About anchors: YAML has a feature called ‘anchors’ that you can use to duplicate content across your document.  So if you have an anchor such as &thisthing:

```
.hello: &thisthing
    image: ruby
    services:
    - postgres
```
and then you extend that anchor, like so:

```
test1:
    <<: *thisthing
    script: test1 project
```

* Then all of the items under &thisthing above will be included and copied down below.  Therefore if we go:

```
Run precommit tests:
  extends: .tests
  rules:
  - if:
    when:
```
* The, ".tests" would be equivalent to doing <<: .tests, but not only that, it would be grabbing ".tests" from the above:

```
include:
  - project: 'my-group/my-project'
    file: '/templates/.gitlab-ci-template.yml'
```
* Which can basically serve as a library of a ton of different references to other yml setups, which allows us to store a huge variety of yml options across many different fields and consolidate them into one file, also using the, "include" field, and then re-populate them back into the yaml we are working on now.
* This avoids having to copy and paste tons and tons of yaml all over every single file, but rather to just create pointers to pointers.
* So for example if we anted to set up something to publish a python package we could do:

```
Build and publish python package:
  extends: .pythonPublish

variables:
  VAR:
```
* So what would our .pythonPublish look like?

```
.pipPublish:
  image: $IMAGE
  stage: build
  script: /scripts/pip_publish.sh
  rules:
    - if: (some rules where we wouldn't want this to happen)
      when: never
    - changes:
      - anything/under/these/files/**/*
      when: on_success
    - if: $PYTHON_PUBLISH_FORCE
      when: on_success
```


##### Building a pip_publish Script

* The pip_publish script included within the extended .pipPublish: yaml is the meat of how pip's get published.
* For more on how to build a pip_publish Script:

* [pip-publish-script.md](/about-pythonpackage/pip-publish-script.md)

##### Predefined CI/CD Pipelines

Any script ran on a Docker container which is running within a Pipeline on a Gitlab Runner are available to be output to be available for a script command within a job.

E.g.: https://docs.gitlab.com/ee/ci/variables/index.html#list-all-variables

Basically, the script can be called as it is ran on a Gitlab Runner, e.g.:

```
job_name:
  script:
    - export
    # - 'dir env:'  # Use this for PowerShell
```

That Gitlab Runner is essentially running a container which we specify, or if we don't specify, then a default Ruby image. This container will then by default have access to all of the Gitlab Predefined Variables, e.g.:

* [Gitlab Predefined Variables](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html)

For example:

```
CHAT_CHANNEL	10.6	all	The Source chat channel that triggered the ChatOps command.
CHAT_INPUT	10.6	all	The additional arguments passed with the ChatOps command.
CHAT_USER_ID	14.4	all	The chat service’s user ID of the user who triggered the ChatOps command.

etc..
```

* Many of these variables are critical to putting together a script and a full pipeline.