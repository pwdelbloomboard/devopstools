# Rundeck 

## Introduction to Rundeck
### Why Use Rundeck in a K8s Environment

Running scripts on staging using rundeck is primarily for engineers testing scripts that will need to run in prod for doing things like data migrations, or some process that needs to run regularly.

* The non-kubernetes analog to a job, a regularly running process (or repeatedly running process) would be a cron job.  Infrastructure which is set up on Kubernetes doesn't use cron jobs, it uses something like Rundeck to set up jobs and run scripts.
* Rundeck does our deployments as well. There is an audit trail, centralized logging. We don't have to try to figure out what server a particular command was on.
* One-off Jobs are done by Rundeck as well...it's kind of like a terminal.

### Video Intro To Rundeck


### Rundeck Installation

## Smoke Test

> A smoketest is a quick check to make sure something works. It's an industry standard term, but has nothing to do with rundeck. Rundeck is just a tool to run jobs, it doesn't care what the jobs do. A rundeck job is a configuration to run one or more scripts or commands that can be run manually via the web interface, via scheduling or via API call.

> the smoketest is just an http endpoint on test-bot our integration test server. 

smoke test is a curl call that blocks, runs a minimal amount of tests, and returns 200 if successful.

## Integration Test

integration tests is the full suite, is triggered asyncronously, and updates an ssm variable when finished, which ops can use to determine the health of the production app
## Pipeline

a pipeline is something that can be spun up to create many parallel users hitting the site at once. Unfortunately, it doesn't really work at the moment.

## Types of Rundeck Jobs



## Running a Job

### Understanding the Context

* First, we have to understand the context 

* Are variables being updated? Why are they being updated? Are we asking certain applications on certain branches being asked to point to one another?
* Is this an override?  Are you overriding a variable in SSM?
* Are you using branch slugs rather than branch names?
* Can you understand the manual implementation of this job?
### Getting to the Job




# Resources

* [Introduction to Rundeck in Under 10 Minutes](https://www.youtube.com/watch?v=QSY_qw9Buic)
* [Rundeck Documentation](https://docs.rundeck.com/docs/)
* [Rundeck Github Repo](https://github.com/rundeck/rundeck)