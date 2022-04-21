# dig

The command dig is a tool for querying DNS nameservers for information about host addresses, mail exchanges, nameservers, and related information. This tool can be used from any Linux (Unix) or Macintosh OS X operating system. The most typical use of dig is to simply query a single host.

# Example

```
 dig google.com

; <<>> DiG 9.10.6 <<>> google.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 22748
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 4, ADDITIONAL: 5

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;google.com.			IN	A

;; ANSWER SECTION:
google.com.		99	IN	A	142.250.190.142

;; AUTHORITY SECTION:
google.com.		56644	IN	NS	ns1.google.com.
google.com.		56644	IN	NS	ns4.google.com.
google.com.		56644	IN	NS	ns2.google.com.
google.com.		56644	IN	NS	ns3.google.com.

;; ADDITIONAL SECTION:
ns1.google.com.		60700	IN	A	216.239.32.10
ns2.google.com.		205814	IN	A	216.239.34.10
ns3.google.com.		309579	IN	A	216.239.36.10
ns4.google.com.		187813	IN	A	216.239.38.10

;; Query time: 9 msec
;; SERVER: 192.168.0.1#53(192.168.0.1)
;; WHEN: Wed Dec 22 07:21:50 CST 2021
;; MSG SIZE  rcvd: 191

```


Basically, it gives a bit more detail than nslookup.

host and nslookup both give highly structured, line by line answers which are useful in scripting. Contrast this to, "dig" which gives much more probing and comprehensive answers, but then it has more of a human readable, manual output.

Host is designed to be a simplified version of nslookup.  Dig is designed to be a more complicated, human version of nslookup.