# traceroute

* Used to trace routes, duh!  Trace your route through various routers to a domain or IP address.
## usage

Basically, you can run a traceroute which will show the results of your computer's attempting to access a domain and the various routers it goes through.

At the time this was written, the author's service provider was US Internet, so you can see the route go through various US Internet servers, then finally it goes to, "74.125.49.46" which is a Google router in Freemont, CA., from there it gets routed to several other Google routers and finally ends at 1e100.net which is the fundamental Google server address.

```
$ traceroute google.com
traceroute to google.com (142.250.191.142), 64 hops max, 52 byte packets
 1  192.168.0.1 (192.168.0.1)  1.607 ms  1.510 ms  3.014 ms
 2  174-141-194-1.fttp.usinternet.com (174.141.194.1)  2.889 ms  3.802 ms  2.204 ms
 3  v95-usi-cr09-mpls.usinternet.com (207.153.5.2)  1.574 ms  1.574 ms  2.099 ms
 4  200ge-po1.usi-cr08-mpls.usinternet.com (207.153.0.30)  1.659 ms  1.568 ms
    100ge-po5.usi-cr07-mpls.usinternet.com (207.153.0.44)  1.658 ms
 5  200ge-po7.usi-cr06-mpls.usinternet.com (207.153.0.18)  2.861 ms
    200ge-po1.usi-cr06-mpls.usinternet.com (207.153.0.16)  1.668 ms
    200ge-po7.usi-cr06-mpls.usinternet.com (207.153.0.18)  1.650 ms
 6  200ge-po4.usi-cr06-mtka.usinternet.com (207.153.0.8)  1.827 ms  1.944 ms  1.978 ms
 7  200ge-po5.usi-cr01-oma.usinternet.com (207.153.0.51)  9.369 ms  9.464 ms  9.345 ms
 8  200ge-po11.usi-cr02-oma.usinternet.com (207.153.0.25)  9.485 ms  9.311 ms  9.611 ms
 9  100ge-po3.usi-cr01-dvn.usinternet.com (207.153.0.46)  15.553 ms  15.505 ms  15.713 ms
10  100ge-po4.usi-cr02-chi.usinternet.com (207.153.0.49)  14.093 ms  14.168 ms  14.037 ms
11  74.125.49.46 (74.125.49.46)  14.125 ms  14.226 ms  14.078 ms
12  108.170.243.174 (108.170.243.174)  15.297 ms
    108.170.243.193 (108.170.243.193)  16.428 ms
    108.170.243.174 (108.170.243.174)  15.361 ms
13  142.251.60.7 (142.251.60.7)  14.309 ms
    142.251.60.5 (142.251.60.5)  14.316 ms
    142.251.60.7 (142.251.60.7)  14.588 ms
14  ord38s29-in-f14.1e100.net (142.250.191.142)  14.308 ms  14.361 ms  14.160 ms
```