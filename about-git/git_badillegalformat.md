# Specific Git Issue - Bad Illegal Format

## Overview of Error and Background

* Injecting username and token into a string.
* Token

## Full Error

```
git clone -q https://user@domain.com:<TOKEN>@gitlab.com/projectname/subprojectname/thing.git
fatal: unable to access 'https://domain.com:<TOKEN>@gitlab.com/projectname/subprojectname/thing.git': URL using bad/illegal format or missing URL
```

## Breakdown

* There may be another script running underneath which acts as a wrapper, which requires a specific input due to a sed command happening.
* Also, the username might need to be different and not have an @ symbol.