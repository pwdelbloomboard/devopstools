# Rundeck 

https://docs.rundeck.com/docs/
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


## Types of Rundeck Jobs

Rundeck can be split up by, "Project," and those projects can contain anything that w

* Common - common jobs might be jobs called by other jobs, sort of like a global variable.
* CronJobs - cron jobs, jobs that run on a schedule.
* Deploy - jobs for actually deploying apps
* EngineerTools - jobs that engineers, software developers may use to run or fix or switch something over.
* Infra - jobs that help support infrastructure
* Testing/Temp - you might want a temp for testing or experimental stuff.

## Running a Job

Jobs may be run by logging into Rundeck, finding the appropriate project and clicking on, "Jobs."  So for example when activating a job under EngineerTools:

EngineerTools >> Jobs

* This shows a list of jobs which could be categorized under Prod and Stage, or some other categorization, or jobs could be searched.

In this case, we can look at a job called, "Point Peryton Branches at each other," which is a job designed to point a front end and a back end at each other.

Looking at the job itself, by clicking on the job (not the "play job" button), you get:

![](/img/pointperytonateachother.png)

This interface shows a few different variables that can be used as inputs to the job. Sometimes these variables are optional, sometimes not.

![](/img/rundeck_entervariables.png)

Going deeper into the job, the actual job code itself can be inspected under, "Definition," which is in the upper right hand corner of the screen, above the variables:

![](/img/pointperytonateachother_definition.png)

Rundeck does not have version control, the code for this job is stored directly within Rundeck.


### Understanding the Context

* First, we have to understand the context 

* Are variables being updated? Why are they being updated? Are we asking certain applications on certain branches being asked to point to one another?
* Is this an override?  Are you overriding a variable in SSM?
* Are you using branch slugs rather than branch names?
* Can you understand the manual implementation of this job?
### Getting to the Job


### Rundeck Plugins

* [Rundeck Plugins](https://docs.rundeck.com/docs/plugins/)

Official Rundeck Plugins listing:

https://docs.rundeck.com/plugins/?$

* Rundeck has a variety of plugins, from integrations with third party services to all sorts of neat tricks.




# Resources

* [Introduction to Rundeck in Under 10 Minutes](https://www.youtube.com/watch?v=QSY_qw9Buic)
* [Rundeck Documentation](https://docs.rundeck.com/docs/)
* [Rundeck Github Repo](https://github.com/rundeck/rundeck)