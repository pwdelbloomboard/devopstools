# Fluentbit

The docs for fluentbit can be found [here](https://docs.fluentbit.io/manual/).

Fluentbit is basically a log collector and forwarder 

## FluentBit Key Concepts

### Events and Records

* Every piece of data that belongs to a log is considered an Event or Record.
* Syslog file eample:

```
Jan 18 12:52:16 flb systemd[2222]: Starting GNOME Terminal Server
Jan 18 12:52:16 flb dbus-daemon[2243]: [session uid=1000 pid=2243] Successfully activated service 'org.gnome.Terminal'
Jan 18 12:52:16 flb systemd[2222]: Started GNOME Terminal Server.
Jan 18 12:52:16 flb gsd-media-keys[2640]: # watch_fast: "/org/gnome/terminal/legacy/" (establishing: 0, active: 0)
```
This involves 4 independent events.

Every event has two components, as a vector: [TIMESTAMP, MESSAGE]

### Filtering

* Append specific information to the Event like an IP address or metadata.
* Select a specific piece of the Event content.
* Drop Events that matches certain pattern.

### Tag

* Every event that gets into FluentBit gets assigned a tag.
* This tag is an internal string that is used in a later stage by the router to decide 

### The Data Pipeline

#### Input

[Inputs](https://docs.fluentbit.io/manual/pipeline/inputs) are a number of possible plugins through which Fluentbit can extract its data.

Inputs include:

* Node Exporter Metrics - Prometheus Node Exporter, to collect system level metrics from operating systems, such as CPU/Disk/Network/Process.
* CPU Metrics - Standard CPU Metrics, per core or whole system.
* Docker Metrics - Memory usage and CPU Consumption
* Dummy - Generates Dummy Events
* FluentBit Metrics - So you can monitor the internals of your pipeline.
* Head - allows reading events from the head of a file, similar to head command.
* Health - Check the health level of a TCP server, connection check at certain time interval.
* Memory Metrics - Shows mem, swap, etc.
* Network I/O Metrics - Shows ping-like network metrics.
* Random - Uses either /dev/urandom to generate simple random values samples or use a unix timestamp as a value.
* Standard Input - Allows retreival of valid JSON text messages over stdin
* Syslog - Collects Syslog messages through a Unitx socket UDP/TCP or over network using TCP/UDP.
* Tail - The tail input plugin allows to monitor one or several text files. It has a similar behavior like tail -f shell command.
* Thermal - The thermal input plugin reports system temperatures periodically -- each second by default.
* Collectd - The collectd input plugin allows you to receive datagrams from collectd service.
* Disk I/O Metrics - The disk input plugin, gathers the information about the disk throughput of the running system every certain interval of time and reports them.
* Docker Events - The docker events input plugin uses the docker API to capture server events.
* Exec - The exec input plugin, allows to execute external program and collects event logs.
* Forward - Used by Fluentbit and Fluentd to route messages between peers.
* HTTP - Send custom records to an HTTP endpoint
* Kernal Logs - Linux Kernal messages
* MQTT - Retrieve message and data from MQTT control packets over a TCP connection
* Process Metrics - Checks the health of a process.
* Serial Interface - Retrieve messages/data from a serial interface.
* StatsD - Receieve messages through statsD protocol
* TCP - Obvious
* Windows Event Log - Obvious

#### Parser

> Dealing with raw strings or unstructured messages is a constant pain; having a structure is highly desired. Ideally we want to set a structure to the incoming data by the Input Plugins as soon as they are collected:

> The Parser allows you to convert from unstructured to structured data. As a demonstrative example consider the following Apache (HTTP Server) log entry:

> Parsers are fully configurable and are independently and optionally handled by each input plugin, for more details please refer to the Parsers section.

[Parsers Documentation](https://docs.fluentbit.io/manual/pipeline/parsers)

So basically, every input plugin uses a different set of parser commands.

> By default, Fluent Bit provides a set of pre-configured parsers that can be used for different use cases such as logs from: Apache, Nginx, Docker, Syslog rfc5424, Syslog rfc3164

* Parsers get configured in a parser configuration file, parsers.conf, not in the global Fluent Bit global configuration file.
* Parsers are configured at start time, either from command line or through main FluentBit config file.
* Parsers can have multiple entries. Example:

```
[PARSER]
    Name        docker
    Format      json
    Time_Key    time
    Time_Format %Y-%m-%dT%H:%M:%S.%L
    Time_Keep   On

[PARSER]
    Name        syslog-rfc5424
    Format      regex
    Regex       ^\<(?<pri>[0-9]{1,5})\>1 (?<time>[^ ]+) (?<host>[^ ]+) (?<ident>[^ ]+) (?<pid>[-0-9]+) (?<msgid>[^ ]+) (?<extradata>(\[(.*)\]|-)) (?<message>.+)$
    Time_Key    time
    Time_Format %Y-%m-%dT%H:%M:%S.%L
    Time_Keep   On
    Types pid:integer
```




#### Filter

#### Buffer

#### Routing


## Docker Image - Trying it Out

The Docker image for Fluentbit can be found [here](https://hub.docker.com/r/fluent/fluent-bit/).

The Docker documentation can be found [here](https://docs.fluentbit.io/manual/installation/docker).

## Creating a Dockerfile for Fluentbit


We create a [Dockerfile](/about-fluentbit/Dockerfile) with the following:

```
FROM fluent/fluent-bit:1.8
```

After that point, build the container inside of the directory with that Dockerfile with:

```
docker build -t fluentbittest .
```

Finally, run:

```
docker run -ti fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i cpu -o stdout -f 1
```
Keep in mind:

* ti 
* -d is detached mode, run the container in the background as fluentbittest
* --name is the name of the container
* -i interactive mode, keeping STDIN open even if not attached.
* After "i" we are now in commands within that, "fluentbittest" container, in interactive mode.
* cpu appears to be a command within that container
* -o is the output tupe stdout
* -f 1 ...it's not clear what that does.

> That command will let Fluent Bit measure CPU usage every second and flush the results to the standard output, e.g:

### Output Once Above Run Successfully

Status: Downloaded newer image for fluent/fluent-bit:1.8
Fluent Bit v1.8.12
* Copyright (C) 2019-2021 The Fluent Bit Authors
* Copyright (C) 2015-2018 Treasure Data
* Fluent Bit is a CNCF sub-project under the umbrella of Fluentd
* https://fluentbit.io

[2022/02/07 20:32:14] [ info] [engine] started (pid=1)
[2022/02/07 20:32:14] [ info] [storage] version=1.1.5, initializing...
[2022/02/07 20:32:14] [ info] [storage] in-memory
[2022/02/07 20:32:14] [ info] [storage] normal synchronization mode, checksum disabled, max_chunks_up=128
[2022/02/07 20:32:14] [ info] [cmetrics] version=0.2.2
[2022/02/07 20:32:14] [ info] [sp] stream processor started
[0] cpu.0: [1644265934.746361100, {"cpu_p"=>38.333333, "user_p"=>17.000000, "system_p"=>21.333333, "cpu0.p_cpu"=>36.000000, "cpu0.p_user"=>14.000000, "cpu0.p_system"=>22.000000, "cpu1.p_cpu"=>39.000000, "cpu1.p_user"=>19.000000, "cpu1.p_system"=>20.000000, "cpu2.p_cpu"=>41.000000, "cpu2.p_user"=>19.000000, "cpu2.p_system"=>22.000000}]
[0] cpu.0: [1644265935.746231400, {"cpu_p"=>49.666667, "user_p"=>23.333333, "system_p"=>26.333333, "cpu0.p_cpu"=>64.000000, "cpu0.p_user"=>28.000000, "cpu0.p_system"=>36.000000, "cpu1.p_cpu"=>46.000000, "cpu1.p_user"=>22.000000, "cpu1.p_system"=>24.000000, "cpu2.p_cpu"=>38.000000, "cpu2.p_user"=>19.000000, "cpu2.p_system"=>19.000000}]

### Attempting to Exec Into Container

Once running a container, one can attempt to exec in with:

```
$ docker exec -it 67af36cc6a3e4d9d5229f32f5004bc9ec834da383d3db467fc494675ba62d705 /bin/sh
OCI runtime exec failed: exec failed: container_linux.go:380: starting container process caused: exec: "/bin/sh": stat /bin/sh: no such file or directory: unknown
```

This runs into a failure because the default fluentbit docker image which means the layers only contain the container and the application, automatically running, no linux container layer involved.

```
docker run --name myfluentbittest -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.8-debug /fluent-bit/bin/fluent-bit -i forward -o stdout -p format=json_lines -f 1
```

From there, we can -it/bash into the container with:

```
docker exec -it myfluentbittest /bin/bash
```

This actually works and we are presented with a root @ containerid.

```
root@d3e8c8c11c9c:/# ls
bin  boot  dev	etc  fluent-bit  home  lib  lib64  media  mnt  opt  proc  root	run  sbin  srv	sys  tmp  usr  var
```

From here we can attempt to find the config file.

```
root@d3e8c8c11c9c:/# find . conf | grep conf
./fluent-bit/etc/parsers_java.conf
./fluent-bit/etc/parsers_ambassador.conf
./fluent-bit/etc/plugins.conf
./fluent-bit/etc/parsers_cinder.conf
./fluent-bit/etc/fluent-bit.conf
./fluent-bit/etc/parsers_openstack.conf
./fluent-bit/etc/parsers.conf
./fluent-bit/etc/parsers_extra.conf
```

Within the fluentbit file we have several .conf files:

```
root@d3e8c8c11c9c:/fluent-bit/etc# ls
fluent-bit.conf  parsers.conf  parsers_ambassador.conf	parsers_cinder.conf  parsers_extra.conf  parsers_java.conf  parsers_openstack.conf  plugins.conf
```




### Following the FluentBit Docker Documentation

Run the following to get a container going that uses fluent-bit.

```
docker run -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i forward -o stdout -p format=json_lines -f 1
```

Reading from our section above on what the various, "input" parts of Fluentbit work:

* -p is the port we run on
* fluent/fluent-bit:1.8 is the tagged image we download
* /fluent-bit/bin/fluent-bit is basically us connecting in to the bin folder
* we use -i to send a command, "forward" which does [forwarding](https://docs.fluentbit.io/manual/pipeline/inputs/forward)
* -o stdout means a standard output on -p
* format=json_lines is obvious
* -f 1 not clear


In parallel, run a container that drives a log with the following:

```
docker run --log-driver=fluentd -t ubuntu echo "Testing a log message"
```

We should get the following output from the logs, after running the above parallel messages:

```
$ docker run -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i forward -o stdout -p format=json_lines -f 1
Fluent Bit v1.8.12
* Copyright (C) 2019-2021 The Fluent Bit Authors
* Copyright (C) 2015-2018 Treasure Data
* Fluent Bit is a CNCF sub-project under the umbrella of Fluentd
* https://fluentbit.io

[2022/02/07 20:45:09] [ info] [engine] started (pid=1)
[2022/02/07 20:45:09] [ info] [storage] version=1.1.5, initializing...
[2022/02/07 20:45:09] [ info] [storage] in-memory
[2022/02/07 20:45:09] [ info] [storage] normal synchronization mode, checksum disabled, max_chunks_up=128
[2022/02/07 20:45:09] [ info] [cmetrics] version=0.2.2
[2022/02/07 20:45:09] [ info] [input:forward:forward.0] listening on 0.0.0.0:24224
[2022/02/07 20:45:09] [ info] [sp] stream processor started
{"date":1644266714.0,"source":"stdout","log":"Testing a log message\r","container_id":"7e5aeb1491e70c8b8acfb71f3ae02bbd68d74e85233cb5f7c867d6fd9495b9a8","container_name":"/condescending_maxwell"}
{"date":1644266722.0,"source":"stdout","log":"Here's another.\r","container_id":"c562e931dce5b1a6e2ccda28a15e4b79199eb023caa871f15c6e29c08fa4d20d","container_name":"/peaceful_varahamihira"}
```

### Keeping the Above Logging Container Open To Create a Stream of Custom Messages

Basically, run the fluentbit configuration mentioned above:

```
docker run -p 127.0.0.1:24224:24224 fluent/fluent-bit:1.8 /fluent-bit/bin/fluent-bit -i forward -o stdout -p format=json_lines -f 1
```

Then in parallel, run the log-driver as an -it:

```
docker run --log-driver=fluentd -it ubuntu
```

### Configuring FluentBit on Docker Container

In order to Configure FluentBit, meaning configuring the parser, a debug Container Image must be used, which contains a linux-like environment, allowing one to exec/bash into the container while running.

The default fluentbit images are, "Distroless," while the "Debug" images use alpine linux. [Here](https://hub.docker.com/layers/fluent/fluent-bit/1.8-debug/images/sha256-00249780660f7eeb0e1c12d5dadaed9a511fdc7314b210f45ca3c74529104dcd?context=explore) is an example.

```
fluent/fluent-bit:1.8-debug
```

This follows the above 


### Fluentbit on Kubernetes

The documentation for processing container logs from Systemd/Journald, enriching logs with Kubernetes Metadata is found, and centralizing logs in third party storage is found [here](https://docs.fluentbit.io/manual/installation/kubernetes).

When Fluent Bit runs, it runs as a Daemonset and it will read, parse and filter the logs of every POD and will enrich each entry with the following information (metadata):

* Pod Name
* Pod ID
* Container Name
* Container ID
* Labels
* Annotations

### Container Runtime Interface Parser

Fluent Bit by default assumes that logs are formatted by the Docker interface standard. However, when using CRI you can run into issues with malformed JSON if you do not modify the parser used. Fluent Bit includes a CRI log parser that can be used instead. An example of the parser is seen below:

```
# CRI Parser
[PARSER]
    # http://rubular.com/r/tjUt3Awgg4
    Name cri
    Format regex
    Regex ^(?<time>[^ ]+) (?<stream>stdout|stderr) (?<logtag>[^ ]*) (?<message>.*)$
    Time_Key    time
    Time_Format %Y-%m-%dT%H:%M:%S.%L%z
```
To use this, change the input section for your configuration from docker to cri.

```
[INPUT]
    Name tail
    Path /var/log/containers/*.log
    Parser cri
    Tag kube.*
    Mem_Buf_Limit 5MB
    Skip_Long_Lines On
```