# Go Viper

* [Go Viper](https://github.com/spf13/viper)

# Overview In Different Words

* Viper is a way to configure Go Applications for 12-factor apps.

## What is a 12-factor app?

> The Twelve-Factor App methodology is a methodology for building software-as-a-service applications. These best practices are designed to enable applications to be built with portability and resilience when deployed to the web.

They are basically, "things that apps should have to make them strong enough to be on the web."  Much of these factors are the types of things that one sees commonly 

### What Are The 12 Factors?

| I    | Codebase            | There should be exactly one codebase for a deployed service with the codebase being used for many deployments.      |
|------|---------------------|---------------------------------------------------------------------------------------------------------------------|
| II   | Dependencies        | All dependencies should be declared, with no implicit reliance on system tools or libraries.                        |
| III  | Config              | Configuration that varies between deployments should be stored in the environment.                                  |
| IV   | Backing services    | All backing services are treated as attached resources and attached and detached by the execution environment.      |
| V    | Build, release, run | The delivery pipeline should strictly consist of build, release, run.                                               |
| VI   | Processes           | Applications should be deployed as one or more stateless processes with persisted data stored on a backing service. |
| VII  | Port binding        | Self-contained services should make themselves available to other services by specified ports.                      |
| VIII | Concurrency         | Concurrency is advocated by scaling individual processes.                                                           |
| IX   | Disposability       | Fast startup and shutdown are advocated for a more robust and resilient system.                                     |
| X    | Dev/Prod parity     | All environments should be as similar as possible.                                                                  |
| XI   | Logs                | Applications should produce logs as event streams and leave the execution environment to aggregate.                 |
| XII  | Admin Processes     | Any needed admin tasks should be kept in source control and packaged with the application.                          |


## Built Viper Demonstration

[Within Docker Go - Go Viper](https://github.com/pwdelbloomboard/devopstools/tree/main/about-go/go-docker/volumebindmount/goviper)

### Why Use Viper

* So viper is all about config files. Any time that you run a program, go viper will look for a config file. Perhaps for the purposes of oauth, using an auth token, you can use config files to store them into a temporary struct in your environment.

#### Settings in Viper

* A common other usage is, "SetEnvPrefix" - which basically, MySQL uses when an image starts up on Docker...viper will check for an environment variable matching the SetEnvPrefix variable.