# tcpdump

https://www.tcpdump.org/

## Description

This is the home web site of tcpdump, a powerful command-line packet analyzer; and libpcap, a portable C/C++ library for network traffic capture.

This flexible, powerful command-line tool helps ease the pain of troubleshooting network issues.

## Usage Example

Using, "tcpdump -D" we can view all of the available interfaces on our computer.

```
$ tcpdump -D
1.en0 [Up, Running]
2.awdl0 [Up, Running]
3.llw0 [Up, Running]
4.utun0 [Up, Running]
5.utun1 [Up, Running]
6.en5 [Up, Running]
7.lo0 [Up, Running, Loopback]
8.bridge0 [Up, Running]
9.en1 [Up, Running]
10.en2 [Up, Running]
11.en3 [Up, Running]
12.en4 [Up, Running]
13.gif0 [none]
14.stf0 [none]
15.ap1 [none]
```

From [here](https://stackoverflow.com/questions/29958143/what-are-en0-en1-p2p-and-so-on-that-are-displayed-after-executing-ifconfig)


```
In arbitrary order of my familarity / widespread relevance:

lo0 is loopback.

en0 at one point "ethernet", now is WiFi (and I have no idea what extra en1 or en2 are used for).

fw0 is the FireWire network interface.

stf0 is an IPv6 to IPv4 tunnel interface to support the transition from IPv4 to the IPv6 standard.

gif0 is a more generic tunneling interface [46]-to-[46].

awdl0 is Apple Wireless Direct Link

p2p0 is related to AWDL features. Either as an old version, or virtual interface with different semantics than awdl.
```

So as an example, we can check out eth0:

```
sudo tcpdump --interface en0
```
Basically, tons and tons of data runs through this, particularly if using a lot of traffic such as a video application.


The format of what the packet analysis looks like in general is:

```
19:12:15.899845 IP 216.17.90.205.https > 192.168.0.103.53921: UDP, length 1250
```
* This basically shows a UDP packet source and destination and size.

* there may also be netbios-ns, which is local internet connectivity.
* 1e100.net is a Google-owned domain name used to identify the servers in our network.  This could come from YouTube or any Google owned property.
* etc. you can Google individual packets.

## Usage

$ tcpdump -h
tcpdump version tcpdump version 4.9.3 -- Apple version 100.100.2
libpcap version 1.9.1
LibreSSL 2.8.3
Usage: tcpdump [-aAbdDefhHIJKlLnNOpqStuUvxX#] [ -B size ] [ -c count ]
		[ -C file_size ] [ -E algo:secret ] [ -F file ] [ -G seconds ]
		[ -i interface ] [ -j tstamptype ] [ -M secret ] [ --number ]
		[ -Q in|out|inout ]
		[ -r file ] [ -s snaplen ] [ --time-stamp-precision precision ]
		[ --immediate-mode ] [ -T type ] [ --version ] [ -V file ]
		[ -w file ] [ -W filecount ] [ -y datalinktype ] [ -z postrotate-command ]
		[ -g ] [ -k ] [ -o ] [ -P ] [ -Q meta-data-expression]
		[ --apple-tzo offset] [--apple-truncate]
		[ -Z user ] [ expression ]

## 