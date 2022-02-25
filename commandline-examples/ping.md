# ping

ping is really simple, just ping a given IP or domain and see what the response time is.

The most famous one is:

```
ping -c 5 8.8.8.8
```

Which pings the all-powerful Google servers and gives you estimates on response times, like so:

```
$ ping -c 5 8.8.8.8
PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: icmp_seq=0 ttl=117 time=21.512 ms
64 bytes from 8.8.8.8: icmp_seq=1 ttl=117 time=14.517 ms
64 bytes from 8.8.8.8: icmp_seq=2 ttl=117 time=21.291 ms
64 bytes from 8.8.8.8: icmp_seq=3 ttl=117 time=21.591 ms
64 bytes from 8.8.8.8: icmp_seq=4 ttl=117 time=14.717 ms

--- 8.8.8.8 ping statistics ---
5 packets transmitted, 5 packets received, 0.0% packet loss
round-trip min/avg/max/stddev = 14.517/18.726/21.591/3.357 ms
```