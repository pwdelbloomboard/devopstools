# Original Link

https://github.com/caarlos0/svu

## Installation

```
brew install caarlos0/tap/svu
```

## Usage

This tool should be used within a repo that is initializd with git. Otherwise, you get:

```
user@whatever:/scripts# svu
svu: error: failed to get current tag for repo: fatal: not a git repository (or any of the parent directories): .git
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

### Usage Help File

```
$ svu --help
usage: svu [<flags>] <command> [<args> ...]

semantic version util

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --metadata                 discards pre-release and build metadata if disabled (--no-metadata)
      --pattern=PATTERN          limits calculations to be based on tags matching the given pattern
      --pre-release              discards pre-release metadata if disabled (--no-pre-release)
      --build                    discards build metadata if disabled (--no-build)
      --prefix="v"               set a custom prefix
      --suffix=SUFFIX            set a custom a custom suffix (metadata and/or prerelease)
      --strip-prefix             strips the prefix from the tag
      --tag-mode=current-branch  determines if latest tag of the current or all branches will be used
  -v, --version                  Show application version.

Commands:
  help [<command>...]
    Show help.

  next* [<flags>]
    prints the next version based on the git log

  major
    new major version

  minor
    new minor version

  patch
    new patch version

  current
    prints current version
```