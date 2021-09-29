# Diagnosing Network Issues

## Running a traceroute

### What is a traceroute?

Traceroute is a command line utility for troubleshooting network connectivity issues. It’s like a map that shows you how data travels from your computer to the internet, and what stops it takes at each point.

Whenever you access a website on the internet, your data travels from your device to your router. Then that data is sent to other routers (or servers) across the world. Each time your data is sent to a new router, it is known as a hop. Running a traceroute lets you see all the hops your data makes and how long it takes between each hop. This way, you will know where your network issues are coming from.

Running a traceroute on your Mac can also tell you about your network latency. Latency is a measurement of how long it takes data to travel from your device to a website or server.

### Performing a Basic Traceroute

The domain, 8.8.8.8 is owned by Google and is a common domain to ping or do a traceroute to as a baseline. To do this simply enter in your terminal:

```
traceroute 8.8.8.8
```

From there you should see a map of the various hops that occur, for example:

```
traceroute to 8.8.8.8 (8.8.8.8), 64 hops max, 52 byte packets
 1  192.168.0.1 (192.168.0.1)  1.865 ms  2.008 ms  2.800 ms
 2  174-141-194-1.fttp.usinternet.com (174.141.194.1)  2.410 ms  2.296 ms  1.825 ms
 3  v96-usi-cr09-mpls.usinternet.com (207.153.5.33)  1.415 ms  1.453 ms  1.403 ms
 4  100ge-po5.usi-cr07-mpls.usinternet.com (207.153.0.44)  3.374 ms  1.449 ms  1.610 ms
 5  200ge-po2.usi-cr02-mpls.usinternet.com (207.153.0.12)  3.218 ms  1.539 ms  1.621 ms
 6  100ge-po12.usi-cr01-chi1.usinternet.com (206.55.180.38)  11.174 ms  9.458 ms  9.370 ms
 7  * * *
 8  108.170.243.174 (108.170.243.174)  10.640 ms  11.079 ms
    108.170.244.1 (108.170.244.1)  9.495 ms
 9  142.251.60.15 (142.251.60.15)  9.929 ms
    142.251.60.13 (142.251.60.13)  9.991 ms
    142.251.60.3 (142.251.60.3)  9.662 ms
10  dns.google (8.8.8.8)  9.806 ms  10.468 ms  11.465 ms
```

### Performing a Traceroute to a Specific Server

To perform a traceroute to a specific server, the IP address of the server must be known.