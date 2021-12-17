# About Grok

Grok is a tool that combines multiple predefined regular expressions to match and split text and map the text segments to keys. Grok sits on top of Regex as a wrapper and is geared specifically toward parsing log files.

Grok is used in a similar way that Regex is used. So for example in Python you might have something along the lines of:

```
r'(?:[+-]?(?:[0-9]+))'
```
with the, r'' enclosing the actual Regex expression, and the expression itself (?:[+-]?(?:[0-9]+)) being a pattern-matching algorithm that pulls out integers, in Grok you would simply write:

```
%{INT}
```
With %{} being a sort of flag signifying the pattern to look for enclosed, and INT taking the place of the Regex pattern above. Grok can be incredibly useful to make code much more readable and workable, in the case of IPV6 Grok pattern, the Regex expression is 1268 characters, whereas in Grok it is simply: %{IPV6}.

Grok ships with about 120 predefined patterns for syslog logs, apache and other webserver logs, mysql logs, etc. It is easy to extend Grok with custom patterns.

The original source for Grok appears to be here at [Google Code](https://code.google.com/archive/p/semicomplete/wikis/Grok.wiki). It's not clear, but Grok may have been developed alongside Elasticsearch/Logstash/Kibana, previously known as the, "ELK Stack."

Logstash is now evidently maintained and held in a [Github organization organized by Elastic.co](https://github.com/logstash-plugins) as a series of plugins which extend its functionality to a wide variety of services.

* The [core logstash patterns are kept here](https://github.com/logstash-plugins/logstash-patterns-core).
* [NewRelic Specific Logstash Patterns are kept here](https://github.com/logstash-plugins/logstash-output-newrelic).
### Grok Tutorial

[Grok Tutorial](https://coralogix.com/blog/logstash-grok-tutorial-with-examples/)

#### Basics

* Grok works by combining text patterns into something that matches your logs.
* The syntax is:

```
%{SYNTAX:SEMANTIC}
```
Where:

* SYNTAX is the name of the pattern that will match your text. 
* For example, 3.44 will be matched by the NUMBER pattern
* 55.3.244.1 will be matched by the IP pattern.

As mentioned above, the default Grok was built by Logstash, and as such if you are using Grok on the command line, you have to utilize what is known as a Logstash configuration file, or just a configuration file, to tell Grok how to deal with incoming files that it is looking at.

The configuration file has three seperate sections:

* Input
* Filter
* Output

That being said, ther eare different types of configuration file formats specified by different applications, environments, programs or whatever location you might be using grok. For example, within Elastic's documentation, they mention the format to be along the lines of:

```
input {
        file {
                type => "/home/samplelog.txt"
        }
}
filter {
        grok {
                match => { "message" => "%{IP:client} %{WORD:method} %{URIPATHP>
        }
}
```
...which appears to be a json-like format which includes hashes. Elastic has extensive documentation for their [Grok filter plugin listed here](https://www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html).  Some of this documentation in terms of options may transfer over to other environments.

However, for the [Grok command-line tool on Debian](https://www.unix.com/man-page/debian/1/grok/), gives a different format for input, with the default format being:

```
	# --- Begin sample grok config
	# This is a comment. :)
	#
	# enable or disable debugging. Debug is set false by default.
	# the 'debug' setting is valid at every level.
	# debug values are copied down-scope unless overridden.
	debug: true

	# you can define multiple program blocks in a config file.
	# a program is just a collection of inputs (files, execs) and
	# matches (patterns and reactions),
	program {
	  debug: false

	  # file with no block. settings block is optional
	  file "/var/log/messages"

	  # file with a block
	  file "/var/log/secure" {
	    # follow means to follow a file like 'tail -F' but starts
	    # reading at the beginning of the file.  A file is followed
	    # through truncation, log rotation, and append.
	    follow: true
	  }

	  # execute a command, settings block is optional
	  exec "netstat -rn"

	  # exec with a block
	  exec "ping -c 1 www.google.com" {
	    # automatically rerun the exec if it exits, as soon as it exits.
	    # default is false
	    restart-on-exit: false

	    # minimum amount of time from one start to the next start, if we
	    # are restarting. Default is no minimum
	    minimum-restart-interval: 5

	    # run every N seconds, but only if the process has exited.
	    # default is not to rerun at all.
	    run-interval: 60

	    # default is to read process output only from stdout.
	    # set this to true to also read from stderr.
	    read-stderr: false
	  }

	  # You can have multiple match {} blocks in your config.
	  # They are applied, in order, against every line of input that
	  # comes from your exec and file instances in this program block.
	  match {
	    # match a pattern. This can be any regexp and can include %{foo}
	    # grok patterns
	    pattern: "some pattern to match"

	    # You can have multiple patterns here, any are valid for matching.
	    pattern: "another pattern to match"

	    # the default reaction is "%{@LINE}" which is the full line
	    # matched.	the reaction can be a special value of 'none' which
	    # means no reaction occurs, or it can be any string. The
	    # reaction is emitted to the shell if it is not none.
	    reaction: "%{@LINE}"

	    # the default shell is 'stdout' which means reactions are
	    # printed directly to standard output. Setting the shell to a
	    # command string will run that command and pipe reaction data to
	    # it.
	    #shell: stdout
	    shell: "/bin/sh"

	    # flush after every write to the shell.
	    # The default is not to flush.
	    flush: true

	    # break-if-match means do not attempt any further matches on
	    # this line.  the default is false.
	    break-if-match: true
	  }
	}
	# -- End config
```

##### Example - IP Request Log

With the following fictional IP Request Log:

```
55.3.244.1 GET /index.html 15824 0.043
```
The pattern for this would be:

```
%{IP:client} %{WORD:method} %{URIPATHPARAM:request} %{NUMBER:bytes} %{NUMBER:duration}
```
* Note that each %{ITEM} enclosed with curly parenthesis includes a SYNTAX:SEMANTIC and these are seperated by spaces.

Within that, a config_file can be created that has the following properties:

```
    input {
      file {
        path => "/home/samplelog.txt"
      }
    }
    filter {
      grok {
        match => { "message" => "%{IP:client} %{WORD:method} %{URIPATHPARAM:request} %{NUMBER:bytes} %{NUMBER:duration}" }
      }
    }
```
Note that the above pattern is designed for 

### Available Grok Patterns Documentation

* [Blog Post from NewRelic on Grok Parsing](https://newrelic.com/blog/how-to-relic/how-to-use-grok-log-parsing)

* The [core logstash patterns are kept here](https://github.com/logstash-plugins/logstash-patterns-core).
* [NewRelic Specific Logstash Patterns are kept here](https://github.com/logstash-plugins/logstash-output-newrelic).

A thourough way to look through all of the core patterns would be to download the relevant Github patterns and do a filtered search for logstash patterns with an awk search:

```
find ./logstash-patterns-core/patterns -type f -exec awk '{print $1}' {} \; | grep "^[^#\ ]" | sort
```
This is done in a [Stackoverflow answer on Grok Patterns here](https://stackoverflow.com/a/42711962).

Ultimately, NewRelic sites a thourough list of all Grok Patterns being listed at this [Grok Debugger Tool](https://grokdebug.herokuapp.com/patterns#).


### Using Automatic Grok Discovery Tool

There is an [Automatic Grok discovery tool](https://grokconstructor.appspot.com/do/automatic) online which may or may not be helpful.
### Grok Pattern Checker

One of the simplest ways to test out grok is to use the Grok Pattern Checker, available at this endpoint:

https://grokdebug.herokuapp.com/

With the above example, having the following as our input:

```
55.3.244.1 GET /index.html 15824 0.043
```
...and using the pattern-matching expression:

```
%{IP:client} %{WORD:method} %{URIPATHPARAM:request} %{NUMBER:bytes} %{NUMBER:duration}
```
You can get the following output (opting for singles format):

```
{
  "client": [
    "55.3.244.1"
  ],
  "IPV6": [
    null
  ],
  "IPV4": [
    "55.3.244.1"
  ],
  "method": [
    "GET"
  ],
  "request": [
    "/index.html"
  ],
  "URIPATH": [
    "/index.html"
  ],
  "URIPARAM": [
    null
  ],
  "bytes": [
    "15824"
  ],
  "BASE10NUM": [
    "15824",
    "0.043"
  ],
  "duration": [
    "0.043"
  ]
}
```

### Installing Grok on Debian

There may be a need to actually install Grok on Debian and to use it within the context of the command line. Within the command line on debian, do:

```
apt-get -y install grok
```
Using grok once it is installed:

```
root@a93522e30f3d:/# grok -h
Usage: grok [-d] <-f config_file>
       -d        (optional) Daemonize/background
       -f file   (required) Use specified config file
```
We may need a sample file to work with in order for grok to actually parse a log line.


#### PyGrok

https://github.com/garyelephant/pygrok

#### GoGrok

Written in GoLang

https://github.com/tsaikd/gogstash

#### Grok Exporter for Prometheus Monitoring

https://github.com/fstab/grok_exporter


# Resources

* [1] [Grok Tutorial](https://coralogix.com/blog/logstash-grok-tutorial-with-examples/)
* [2] [Elastic.co Grok Plugin Overview and Documentation](https://www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html)
* [3] [Grok Exporter for Prometheus Monitoring](https://github.com/fstab/grok_exporter)
* [4] [Debugging Logstash Configuration File](https://dzone.com/articles/how-to-debug-your-logstash-configuration-file)
* [5] [Debian Grok](https://www.unix.com/man-page/debian/1/grok/)
* [6] [Parsing in Grok - New Relic](https://newrelic.com/blog/how-to-relic/how-to-use-grok-log-parsing)