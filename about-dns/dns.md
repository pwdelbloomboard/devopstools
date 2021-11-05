# DNS

### How do you tell if a problem is caused by DNS?

* A lot of times folks can't tell if an issue is caused by DNS
* They may tag a problem as a server issue, "slow / down."

#### Don't Try to Interpret Browser Error Messages

* Browsers may mention troubles finding a site or server being down error, this may be misleading.
* Use the Command Line

Python:

```
import requests
r = requests.get('http://example.com')

```

* Search/Google the error message. The results may tell you if it means, "resolving DNS failed."

#### Use Dig

#### Check Against More than One DNS Server


#### spy on the DNS requests being made with tcpdump



# Resources

* https://en.wikipedia.org/wiki/1.1.1.1

* https://blog.cloudflare.com/refresh-stale-dns-records-on-1-1-1-1/

* https://jvns.ca/blog/2021/11/04/how-do-you-tell-if-a-problem-is-caused-by-dns/