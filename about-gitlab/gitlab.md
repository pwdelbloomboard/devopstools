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


# Resources

* [Gitlab Quick Actions - YouTube](https://www.youtube.com/watch?v=aNscanHxu8I)

