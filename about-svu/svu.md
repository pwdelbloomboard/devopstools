# Original Link

https://github.com/caarlos0/svu

## Installation

```
brew install caarlos0/tap/svu
```

## Print the Next Tag

```
$ svu
v0.0.0
```

or 

```
$ svu current
v0.0.0
```

* Compares the log between the latest tag and HEAD.
* Should follow SEMVAR MAJOR.MINOR.PATCH convention.

## Increment Major, Minor or Patch

```
$ svu major
v1.0.0
```

* Follows SEMVAR principles, you can feed in major, minor or patch to increment the version.

## Tagging Option Flags

* There are various tag modes available which relate to the git branch being used - current-branch and all-branches.
* We can also discard pre-release and build metadata.

### Stripping Tag Prefix

If you want to take the "v" off for example:

```
$ svu major --strip-prefix
1.0.0
```

### Suffix

You can also add a suffix.

```
$ svu major --suffix rc
v1.0.0-rc
```