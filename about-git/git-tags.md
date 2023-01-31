# Git Tagging

* Just as you can tag things in docker, you can tag things in git for release.

* If you haven't set up a tag, and attempt to list tags with ```git tag``` then you get nothing in the terminal:

```
git tag

```

There are two main types of tags - annotated and lightweight.

## Annotated Tags

* Starting out, creating an annotated tag as follows:

```
$ git tag -a 0.0.30 -m "Version 0.0.30"
```

* Then if you list out the tags, you get:

```
git tag -l
0.0.30
```

* If you wanted further data shown you would have:

```
git show 0.0.30

...

tag 0.0.30
Tagger: Patrick Delaney <email@email.com>
Date:   Tue Jan 31 13:09:35 2023 -0600

Version 0.0.30

commit fa4a48a600b7d860ed1bf512257aa85e36227a95 (HEAD -> master, tag: 0.0.30)
Author: Patrick Delaney <email@email.com>
Date:   Tue Jan 31 12:42:33 2023 -0600

    PATCH: the thing.

diff --git a/script.sh b/script.sh
new file mode 100755
index 0000000..4eceebc
--- /dev/null
+++ b/script.sh
```

## Lightweight Tags

* Basically, no annotated data!

```
$ git tag 0.0.31
$ git tag -l
0.0.30
0.0.31
```

## Tagging Commits In the Past

* You can pull up a list of previous commits:

```
$ git log --pretty=oneline
9fceb02d0ae598e95dc970b74767f19372d61af8 Update rakefile
```

* And optionally tag a particular commit, commit by commit.

```
$ git tag -a v1.2 9fceb02
```
* This will add a particular tag to that individual commit.

## Deleting Tags

* Tags may also of course be deleted from commits.

```
$ git tag -d v1.4-lw
Deleted tag 'v1.4-lw' (was e7d5add)
```

## Checking Out Tags

* Once a tag is created, it can be checked out, similar to a branch being checked out.

```
$ git checkout v2.0.0
Note: switching to 'v2.0.0'.
```

## Structured Re-Versioning

* Given the helpfulness of tags, another layer for structuring tags is warranted.
* This can be accomplished with [svu](/about-svu/svu.md)