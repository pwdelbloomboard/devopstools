# Gitlab Pipelines

[Gitlab Pipelines](https://docs.gitlab.com/ee/ci/pipelines/)

## Abstract

* Pipelines are considered the top-level component of continuous integration, delivery and deployment.

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

