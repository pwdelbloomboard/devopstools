# Rundeck 

* We have jobs that run on a schedule to do certain things...Non-kubernetes analog would be a cron job.  We don't set up cron jobs.
* Rundeck does our deployments as well. There is an audit trail, centralized logging. We don't have to try to figure out what server a particular command was on.
* One-off Jobs

Nginx is its own seperate app.


## Smoke Test


> A smoketest is a quick check to make sure something works. It's an industry standard term, but has nothing to do with rundeck. Rundeck is just a tool to run jobs, it doesn't care what the jobs do. A rundeck job is a configuration to run one or more scripts or commands that can be run manually via the web interface, via scheduling or via API call.

> the smoketest is just an http endpoint on test-bot our integration test server. 

smoke test is a curl call that blocks, runs a minimal amount of tests, and returns 200 if successful.

## Integration Test

integration tests is the full suite, is triggered asyncronously, and updates an ssm variable when finished, which ops can use to determine the health of the production app
## Pipeline

a pipeline is something that can be spun up to create many parallel users hitting the site at once. Unfortunately, it doesn't really work at the moment.