# Host

> host command in Linux system is used for DNS (Domain Name System) lookup operations. In simple words, this command is used to find the IP address of a particular domain name or if you want to find out the domain name of a particular IP address the host command becomes handy. You can also find more specific details of a domain by specifying the corresponding option along with the domain name.

## Usage

```
$ host google.com
google.com has address 142.250.191.142
google.com has IPv6 address 2607:f8b0:4009:818::200e
google.com mail is handled by 30 alt2.aspmx.l.google.com.
google.com mail is handled by 10 aspmx.l.google.com.
google.com mail is handled by 40 alt3.aspmx.l.google.com.
google.com mail is handled by 50 alt4.aspmx.l.google.com.
google.com mail is handled by 20 alt1.aspmx.l.google.com.
```

Another example:

```
$ host patdel.com
patdel.com has address 13.227.39.23
patdel.com has address 13.227.39.118
patdel.com has address 13.227.39.91
patdel.com has address 13.227.39.48
```

Basically, it gives a bit more detail than nslookup.

host and nslookup both give highly structured, line by line answers which are useful in scripting. Contrast this to, "dig" which gives much more probing and comprehensive answers, but then it has more of a human readable, manual output.

Host is designed to be a simplified version of nslookup.  Dig is designed to be a more complicated, human version of nslookup.