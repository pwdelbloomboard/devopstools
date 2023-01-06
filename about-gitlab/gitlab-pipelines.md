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

```
include:
  - project: $PROJECT_NAME
    file: $FILE_NAME
    ref: master
```

```

```