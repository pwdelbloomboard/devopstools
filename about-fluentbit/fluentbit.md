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

Parsers are independently and optionally handled by each input plugin.  The [parsers.conf](/about-fluentbit/parsers.conf) file shows some of the defaults.

All of the services handled by parsers can be viewed under [the parsers configuration documentation](https://docs.fluentbit.io/manual/pipeline/parsers)

* All parsers must be defined in a parsers.conf file, not in the Fluent Bit global configuration file.
#### Filter

> In production environments we want to have full control of the data we are collecting, filtering is an important feature that allows us to alter the data before delivering it to some destination.

* Filtering is implemented through plugins.
* Each filter can be used to match, exclude or enrich your logs.
* Many filters are supported - a common use case us K8s deployments. Every pod log needs the proper metadata associated.
* Filters, similar to input plugins, run in an instance context, which has its own independent configuration.

More information on filters described in the [filters](https://docs.fluentbit.io/manual/pipeline/filters) section.

The default plugins.conf configuration is shown under [plugins.conf](/about-fluentbit/plugins.conf).

There are scripted, command-line based instantiations and declarative based instantiations of each of these below.

The different types of filters includes:

* AWS Metadata - enriches logs with AWS Metadata. Adds EC2 instance ID and availability zone to log records.
* Expect - validate that records match certain criteria in their structure, like validating that a key exists or it has a specific value.  Expect can actually validate data and structure, basically integrating it into our CI systems to ensure that the expected structure is being used. More information on this found [here](https://docs.fluentbit.io/manual/local-testing/validating-your-data-and-structure).
* Grep - allows you to match or eclude specific records based upon a regular expression pattern.
* Lua - allows you to modify incoming records with custom Lua scripts.
* Record Modifier - Alows to append fields or to exclude specific fields.
* Multiline - Helps to concatenate messages that were orignally one context but later were split across multiple records or log lines. There are built-in parsers that allow multi-line functionality for Go, Python, Ruby, or Java.
* Nest - Works with nested data, allows you to take records and put them in a nest.
* Rewrite Tag - There may be scenarios where the tag needs to be modified on the fly, this allows us to re-emit a record under a new tag in order to route it the proper direction.
* Throttle - Sets the average rate of messages per interval based upon a leaky bucket and sliding window algorithm.
* Checklist - Looks up if a value in a specified list exists and then allows the addition of a record to indicate if found.
* GeoIP2Filter - Filter allows you to enrich the incoming data stream using location data from GeoIP2 database.
* Kubernetes - enrich your log files with Kubernetes metadata.
* Parser - The Parser Filter plugin allows for parsing fields in event records.  Basically, Parsers do not happen automatically, they have to go through a Filter in order to be applied. Parsers are built "shapes," which filters can then apply.
* Modify - The Modify Filter plugin allows you to change records using rules and conditions.
* Standard Output - Allows printing to the standard output the data flowed through the filter plugin, which can be useful for debugging.
* Tensorflow - Allows running machine learning interface tasks on the records of data coming from input plugsin or a stream processor. Uses Tensorflow Lite as the interface engine (used for mobile and IoT applications).

###### Record Modifier in Detail

https://docs.fluentbit.io/manual/pipeline/filters/record-modifier

Allows appending certain fields or to exclude specific fields.

###### Modify in Detail

https://docs.fluentbit.io/manual/pipeline/filters/modify

* So for example you could take a json input, and rename a key to another name, and add an additional entry key/value pair.
* Contains operations Set, Add, Remove, (also Remove Wildcard, Remove Regex), Rename, Copy, etc. 
* You can also use conditionals such as, "Key_exists" or "keymatches" and others.
###### Parser in Detail

https://docs.fluentbit.io/manual/pipeline/filters/parser

This is what would allow parsing on our cluster.

###### Multiline Modifier in Detail

https://docs.fluentbit.io/manual/pipeline/filters/multiline-stacktrace

* Allows combining multiline outputs into single line outputs or just working with multiline outputs in general.

###### Nest Modifier in Detail

https://docs.fluentbit.io/manual/pipeline/filters/nest

* Basically you can operate with or on nested data.
###### Rewrite Tag Modifier in Detail

https://docs.fluentbit.io/manual/pipeline/filters/rewrite-tag

* Tags are what makes routing possible. You can re-write tags in certain scenarios with conditions.
###### Standard Output in Detail

You can have the standard output data folow through the filter plugin, which can be useful for debugging.

https://docs.fluentbit.io/manual/pipeline/filters/standard-output
###### Checkist in Detail

https://docs.fluentbit.io/manual/pipeline/filters/checklist

The following plugin looks up if a value in a specified list exists and then allows the addition of a record to indicate if found. 

###### Tensorflow Light Detail

https://docs.fluentbit.io/manual/pipeline/filters/tensorflow

#### Buffer

#### Routing


#### Outputs

https://docs.fluentbit.io/manual/pipeline/outputs

There are many different output options, a few:

* NewRelic - https://docs.fluentbit.io/manual/pipeline/outputs/new-relic
* HTTP 
* Forward
* Standard Output
* Syslog
* Amazon Kinesis

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
docker run --name testlogger --log-driver=fluentd -it ubuntu
```
#### About FluentD Log Driver

* What is this FluentD log driver doing?

More info about it can be found [here](https://docs.docker.com/config/containers/logging/fluentd/).

> The fluentd logging driver sends container logs to the Fluentd collector as structured log data. Then, users can use any of the various output plugins of Fluentd to write these logs to various destinations. In addition to the log message itself, the fluentd log driver sends the following metadata in the structured log message:

* container_name
* container_id
* source
* log

Different options, shown in the documentation can be specified upon running the container. So for example:

```
docker run --log-driver=fluentd --log-opt fluentd-address=fluentdhost:24224
```

The options include:

* fluentd-address - specified as a fluentdhost:port, tcp://port, or path
* tag - by default, the first 12 characters of the docker id
* labels, labels-regex, env, env-regex ... are used to tag logs, which is talked about in docker documentation under [customizing log driver output](https://docs.docker.com/config/containers/logging/log_tags/)
* fluentd-async - connects to fluentd in the background. Messages are buffered until the connection is established. Defaults to false.
* fluentd-buffer-limit - set the numebr of events buffered on the memory.
* fluentd-retry-wait - how long to wait before retries.
* fluentd-max-retries
* fluentd-sub-second-precision

##### Adding FluentD Log Driver to Running Nginx Container

We have coverd [nginx](/about-nginx/nginx.md).

We can use the Dockerfile, build an image called, "nginxtest" with content and then run an nginx container on a vistable port, while forwarding the logs via the log-driver using fluentd, also allowing us to bash in to the image:


```
docker build -t nginxtest .

docker run --name some-nginx -d -p 8080:80 --log-driver=fluentd -it nginxtest
```
If need be we can exec into this container with:

```
docker exec -it some-nginx /bin/bash
```

However in this scenario, sending, echo "hello world" to stdout will not be picked up on the fluentbit listener side, only the web visits.

### Configuring FluentBit on Docker Container

In order to Configure FluentBit, meaning configuring the parser, a debug Container Image must be used, which contains a linux-like environment, allowing one to exec/bash into the container while running.

The default fluentbit images are, "Distroless," while the "Debug" images use alpine linux. [Here](https://hub.docker.com/layers/fluent/fluent-bit/1.8-debug/images/sha256-00249780660f7eeb0e1c12d5dadaed9a511fdc7314b210f45ca3c74529104dcd?context=explore) is an example.

```
fluent/fluent-bit:1.8-debug
```

This follows the above section on, [attempting to exec into a container](https://github.com/pwdelbloomboard/devopstools/blob/main/about-fluentbit/fluentbit.md#attempting-to-exec-into-container).

#### fluent-bit.conf

[fluent-bit.conf](/about-fluentbit/fluent-bit.conf)

* This appears to be the, "root" configuration file which sets both the input and output of the stream as mentioned within the fluentbit documentation on input and output.

#### parsers.conf

[parsers.conf](/about-fluentbit/parsers.conf)

* this appears to be a long default list of common types of logs to parse, including:

* apache
* apache2
* apache_error
* nginx
* json
* docker
* docker-daemon
* syslog-rfc5424
* syslog-rfc3164-local
* syslog-rfc3164
* mongodb
* envoy
* cri
* kube-custom

#### plugins.conf


#### Installing Tools to Edit Conf Files

This docker image is based upon debian, so we can do the following commands to install editor files and other possibly needed tools:

```
apt-get update
apt-get install curl
apt-get install wget
apt-get install vim
apt-get install nano
```
#### Editing parsers.conf to Filter Pre-Set Phrases



## Fluentbit on Kubernetes

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