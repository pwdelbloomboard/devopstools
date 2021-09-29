# Debug Scripts

## About Debugging Scripts

A debug script is a way to look at the status of various aspects of system which is running a set of applications. Debug scripts are useful to help with debugging because they tell the devops team the operation system configuration of the underlying hardware as well as the container, cluster, databasae and service ingress health, among other reports that may wish to be created.

## Code Analysis

The overall code for a debug script looks like the following

```
#!/usr/bin/env bash

# Dump info about your system to a text file, useful for getting help debugging.

(some settings)

~set directory

neofetch --config "${DIR}/config/neofetch-config.conf" | tee "${DIR}/debug.txt"
echo "Above output has been saved to ${DIR}/debug.txt"

kubectl cluster-info dump > "${DIR}/cluster.txt"
echo "Output of kubectl cluster-info dump saved to ${DIR}/cluster.txt"
```

The key parts of the debug task are the following:

* neofetch
* kubectl cluster-info dump

#### neofetch

[neofetch](https://github.com/dylanaraps/neofetch) 

* Neofetch is a command-line system information tool written in bash 3.2+. Neofetch displays information about your operating system

```
neofetch --config "${DIR}/config/neofetch-config.conf" | tee "${DIR}/debug.txt"
echo "Above output has been saved to ${DIR}/debug.txt"
```

So everything from the neofetch will be saved to [debug.txt](debuglog.md) based upon the configurations.

##### neofetch configurations

The [neofetch configuration documentation](https://github.com/dylanaraps/neofetch/wiki/Customizing-Info) contains much more documentation on how the configuration is set up.

This file is organized as a, ".conf" file.

###### .conf Files

A CONF file is a configuration or "config" file used on Unix and Linux based systems. It stores settings used to configure system processes and applications. CONF files are similar to .CFG files found on Windows and Macintosh systems.

This simply is used to set the parameters to set the initial settings for neofetch.

```
| sed 's/:/|/' | column -ts'|'
```

##### Important Points

The most important points to know about the architecture of how a neofetch file is set up are:

* custom_section is a list of custom sections that are withdrawn by neoconfig from the system
* kernal
* distro
* uptime
* memory
* packages
* shell
* CPU
* GPU
* resolution
* IP Address
* Public IP timeout
* local IP interface
* desktop environment
* disk

There are custom sections which can be written, for example:

```
{
...
custom_section docker-mac-info
...
}
```

...leads to a function which is set up to withdraw information about containers and groups stored within the appropriate directory on a Mac.

##### Troubleshooting Problems Using debug.txt

* Refer to [debuglog.md](/understanding-scripts/debuglog.md)

#### kubectl cluster-info dump > "${DIR}/cluster.txt"

Basically, this dumps a lot of relevant info for diagnosis.

```
kubectl cluster-info dump > "${DIR}/cluster.txt"
```

https://www.mankier.com/1/kubectl-cluster-info-dump

##### Troubleshooting Problems Using clusterlog.txt

Check out [clusterlog.md](/understanding-scripts/clusterlog.md)


#### kubectl get events > events.txt

Get a list of events and put them into a .txt file.

* The results of this comand are described more at [eventslog.md](/understanding-scripts/eventslog.md)

##### Troubleshooting Problems using events.txt

* Check out [eventslog.md](/understanding-scripts/eventslog.md)