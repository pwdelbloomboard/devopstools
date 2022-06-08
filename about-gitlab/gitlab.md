### Viewing a Git Diff

* [Gitlab Git Diffs](https://docs.gitlab.com/ee/development/diffs.html)

#### From a Merge Request

1. Go to the merge request.
2. Click on "Changes," at the top.

### How to Do Draft Merge Requests

[Draft Merge Requests](https://docs.gitlab.com/ee/user/project/merge_requests/drafts.html)

There are several ways to flag a merge request as a draft:

* Viewing a merge request: In the top right corner of the merge request, click Mark as draft.
* Creating or editing a merge request: Add [Draft], Draft: or (Draft) to the beginning of the merge request’s title, or click Start the title with Draft: below the Title field.
* Commenting in an existing merge request: Add the /draft quick action in a comment. This quick action is a toggle, and can be repeated to change the status again. This quick action discards any other text in the comment.
* Creating a commit: Add draft:, Draft:, fixup!, or Fixup! to the beginning of a commit message targeting the merge request’s source branch. This is not a toggle, and adding this text again in a later commit doesn’t mark the merge request as ready.

![](/img/gitlabcreatingdraft.png)

#### Commenting In An Existing Merge Request

1. Finish desired changes.

2. "git add ."

3. "git commit -m "Commit message."

4. "git push"

5. Go into the Gitlab branch and find the commit.

6. Within the comments section under, "write" enter "/draft" and click, "save."

7. The commit to that branch should not be marked as, "draft."

8. Create a merge request.


### Gitlab Quick Actions

[Quick Actions](https://docs.gitlab.com/ee/user/project/quick_actions.html#issues-merge-requests-and-epics)

> Quick actions are text-based shortcuts for common actions that are usually done by selecting buttons or dropdowns in the GitLab user interface. You can enter these commands in the descriptions or comments of issues, epics, merge requests, and commits.

> If you manually enter a parameter, it must be enclosed in double quotation marks ("), unless it contains only these characters:

> ASCII letters
> Numbers (0-9)
> Underscore (_), hyphen (-), question mark (?), dot (.), or ampersand (&)

### General Info About How Gitlab Works

* Gitlab organizes things into Groups and Projects
* A project is made up of several things. 
-- The git repo
-- Image registry
-- MR infrastructure

* While the git repo and the image registry are part of the same project, the way the docker images are tagged are up to the user.
* The first part of the tag for an image will be the same for every image in the registry for a given project. 
* For example, for projectA it would be: "registry.gitlab.com/organisationname/projectA

* What we add after that is up to us.
* Using the app_name and branch_name as part of the registry name is done because it makes it easy to understand/manage/automate. But it could be literally any alpanumeric string that meets the requirements of a docker image tag.
* So it doesn't matter if a branch exists or not, you can always tag the image and push it to the registry.
* If you try to pull an image that has not yet been built, tagged and pushed to the registry it will fail.

* There are build scripts which can be set up to help with automation, building images on Gitlab. These scripts can automate the process of pushing things to the registry once they are tagged appropriately with the appopriate branch name and pushed to the registry.

* One automation design choice that can be made is to default to using the master branch for the base image.
* When building an image for a new application, since there is no, "master branch," that exists at the start, a one-time manual build and push of an image tagged as a master branch can be done, to get the registry ready.

* The act of pushing to the registry, where it is effectively built by Gitlab (through the process of pushing to a registry), is That enables testing before merging to master. Then when merging to master a new image will be built and tagged as master.

* build-secrets, if used within a Dockerfile, should never be created manually. There should be scripts set up within a repo to help build a particular part of a repo.

#### Gitlab Runners vs. Pipeline

When setting up Continuous Integration (CI), this is done by providing a custom pipeline definition, which is within a yaml file. The Runner does not need to be configured in order to set up a custom pipeline...it is the pipeline which dictates the testing, while the runner is the specific server running the pipeline.

Configuration of a server might be needed if a particular pipeline is for example, undergoing a huge amount of testing, or testing on something that requires a massive amount of data - then you might need a non-standard runner.

https://docs.gitlab.com/runner/


### Gitlab Pipelines

* Pipelines are the ordered part of Gitlab CI.

#### Caching in Gitlab CI/CD

Caching is used for dependencies, like packages you download from the internet. Cache is stored where Gitlab Runners are installed and uploaded to S3 if distribued cache is enabled.
#####  Cache

* Defines per job using the cache keyword.
* Subsequent pipelines can use the cache
* Subsequent jobs in the same pipeline can use the cache
* Different projects 

#### Job Artifacts

* [Job Artifacts](https://docs.gitlab.com/ee/ci/pipelines/job_artifacts.html)

* An archive of files and directories is known as an artifact.
* 

#### Pipeline Artifacts

## Gitlab Auto-Devops for Gitlab CI

https://docs.gitlab.com/ee/topics/autodevops/

A good starting place for setting up an automation pipeline is the Gitlab auto-devops job.

There can be a default pipeline definition used across all repos, with additional standard pipelines added as needed. For example, a different pipeline could be created for a front-end repo.

Pipeline jobs are defined in yaml files which can be in a single file or split across multiple files.

#### Autodevops Features

https://docs.gitlab.com/ee/topics/autodevops/#auto-devops-features

* Autobuild - builds from Dockerfile
* Auto Dependency Scanning - scans for dependencies

Test your app:

* Auto Test
* Auto Browser Performance Testing
* Auto Code Intelligence
* Auto Code Quality
* Auto Container Scanning
* Auto License Compliance

Deploy your app:

* Auto Review Apps
* Auto Deploy
* Monitor your app:

* Auto Monitoring

Secure your app:

* Auto Dynamic Application Security Testing (DAST)
* Auto Static Application Security Testing (SAST)
* Auto Secret Detection


## Gitlab Applications - Configuring GitLab as an Oauth 2.0 Authentication Identity Provider

https://docs.gitlab.com/ee/integration/oauth_provider.html 

### What is an Oauth 2.0 Authentication Identity Provider?

> Authentication in the context of a user accessing an application tells an application who the current user is and whether or not they're present. A full authentication protocol will probably also tell you a number of attributes about this user, such as a unique identifier, an email address, and what to call them when the application says "Good Morning". Authentication is all about the user and their presence with the application, and an internet-scale authentication protocol needs to be able to do this across network and security boundaries.

> However, OAuth tells the application none of that. OAuth says absolutely nothing about the user, nor does it say how the user proved their presence or even if they're still there. As far as an OAuth client is concerned, it asked for a token, got a token, and eventually used that token to access some API. It doesn't know anything about who authorized the application or if there was even a user there at all. In fact, much of the point of OAuth is about giving this delegated access for use in situations where the user is not present on the connection between the client and the resource being accessed. This is great for client authorization, but it's really bad for authentication where the whole point is figuring out if the user is there or not (and who they are).

* This is why Gitlab refers to Oauth providing client applications, "Secure Delegated Access" to server resources on behalf of a resource owner.

* Once an application is created, external services can manage tokens using hte OAuth 2 API.

* Oauth can be used to allow users to sign in to your application with a Gitlab Account (single sign-on)
* ... to set up Gitlab.com for authentication to your Gitlab instance.
* The above can be achieved either through User-Owned or Group-Owned Applications.
* These applications are given TOKEN's and ID's, which can be stored on a key-value store such as AWS-SSM, where they can be accessed by applications within different environments as a way to gain access.



# Resources

* [Gitlab Quick Actions - YouTube](https://www.youtube.com/watch?v=aNscanHxu8I)

