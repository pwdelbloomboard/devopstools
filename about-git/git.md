# git

This is a collection of notes on, "git" in general, not a beginner how-to on git, but rather some more advanced philosophies and procedural thoughts.

### Merging vs. Rebasing

* Merging may be the preferred chosen strategy over rebasing in order to preserve history.

* Both, "merge" and "rebase" solve the same problems -- resolving differences. However they accomplish the solve in different ways.

#### Merging

Mechanistically, this can be done by either of the two methods:

```
git checkout feature
git merge main

or

git merge feature main
```
![](/img/mergecommit.png)

* Merging is non-descructive, the existing branches are not changed in any way.
* This avoids potential pitfalls of rebasing.
* However, there will be an extra merge commit every time you need to incorporate upstream changes.
* If main is very active, this can pollute your feature banch history quite a bit.
* This can make it hard for developers to understand the history of the project.

#### Rebasing

Mechanistically, rebasing can be done like the following:

```
git checkout feature
git rebase main
```

* This moves the entire feature branch to begin on the tip of the main branch, effectively incorporating all of the new commits in main.
* Instead of using a merge commit, rebassing re-writes the project history by creating brand new commits for each commit in the original branch.
* This means a much cleaner project history.
* Re-writing project history can be catastrophic for project collaboration workflow, and loses the context provided by a merge commit, e.g. you can't see when upstream changes were made.

![](/img/rebasecommit.png)

#### Interactive Rebasing


# Resources

* [Merging vs. Rebasing](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)