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


# Resources

* [Gitlab Quick Actions - YouTube](https://www.youtube.com/watch?v=aNscanHxu8I)

