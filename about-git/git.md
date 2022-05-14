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

A better way to do this is to consistenly interactively rebase over time.

```
git checkout feature
git rebase -i main
```

What this does is eliminates the multitude of commits which have been done on your feature branch, and just shows the last few since the time it was branched off.

This reduces, "clutter" and doesn't create the appearance that there was suddently 100+ or however many changes that all of the sudden were done in one hour since the previous commit.

#### Merging Master into Feature Branch

To take care of problems, another way to go about things is to merge a master or main branch into a feature branch on local.

```
# move back over to the main branch
git checkout master

# pull main branch features
git pull

# go back over to the feature branch
git checkout feature

# merge the main into the feature branch
$ git merge main

```

Note at this point, there may be conflicts, and a message:

* "Automatic merge failed; fix conflicts and then commit the result."

From here, you can open up a merge tool to work through the differences in the branch.

```
git mergetool
```

##### Using Git Mergetool

Before using the Git Mergetool, you may need to configure it with the following:

```
git config merge.tool vimdiff
git config merge.conflictstyle diff3
git config mergetool.prompt false
```

* The top left is the LOCAL version of a file, with the feature.
* the middle is the BASE version, before any changes were made.
* the right is the REMOTE version, or main branch, which we are trying to merge into the left.

* The objective is to get the file on the bottom to look like we want it to look like.

![](/img/gitmergetool.png)

The text editor to use while operating gitmerge is VIM.

###### No Conflicts

Sometimes there are no conflicts, you're good to go!


## Selected Git Errors

### directory does not have a commit checked out

```
error: 'about-go/go-docker/volumebindmount/quickstart/' does not have a commit checked out
fatal: adding files failed
```
* Basically there is a leftover git in a directory.


# Resources

* [Merging vs. Rebasing](https://www.atlassian.com/git/tutorials/merging-vs-rebasing)
* [Using Git Mergetool](https://stackoverflow.com/questions/161813/how-to-resolve-merge-conflicts-in-a-git-repository)
* [Git Mergetool Explained](https://www.youtube.com/watch?v=wxh-AOxPX_A)