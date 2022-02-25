# nslookup

* nslookup is simply nameserver lookup which looks up the name servers, for example:

```
$ nslookup patdel.com
Server:		192.168.0.1
Address:	192.168.0.1#53

Non-authoritative answer:
Name:	patdel.com
Address: 13.227.39.48
Name:	patdel.com
Address: 13.227.39.23
Name:	patdel.com
Address: 13.227.39.118
Name:	patdel.com
Address: 13.227.39.91
```

This runs a nameserver lookup on patdel.com, my personal domain, and displays the addresses, which are actually the addresses configured by Namecheap, the domain provider. 