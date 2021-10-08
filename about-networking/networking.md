# Networking

## Listening to Ports

* The [lsof command](https://en.wikipedia.org/wiki/Lsof), "List Open Files," is for unix-like operating systems. Open files in the system include disk files, named pipes, network sockets and devices opened by all processes.

So while running a pod, we can run the command:

```
sudo lsof -i -P | grep LISTEN 
```
We run this in sudo because the MacOS may protect and not show certain processes if this is not run as a sudo command.  When we run this, we get the following:

```
remoted      86                          root    4u  IPv6 0x9f77ff03ae927e21      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49154 (LISTEN)
remoted      86                          root    5u  IPv6 0x9f77ff03ae928481      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49155 (LISTEN)
remoted      86                          root    6u  IPv6 0x9f77ff03ae928ae1      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49156 (LISTEN)
remoted      86                          root    7u  IPv6 0x9f77ff03ae927161      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49157 (LISTEN)
remoted      86                          root    8u  IPv6 0x9f77ff03af7c27c1      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49158 (LISTEN)
remoted      86                          root    9u  IPv6 0x9f77ff03af7c2e21      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:49159 (LISTEN)
Lens       4386                    local_user   40u  IPv4 0x9f77ff03cdd9f089      0t0    TCP localhost:51351 (LISTEN)
Lens       4386                    local_user   42u  IPv4 0x9f77ff03beed77e9      0t0    TCP localhost:59968 (LISTEN)
Lens       4386                    local_user   53u  IPv4 0x9f77ff03cdd9bdc1      0t0    TCP localhost:60010 (LISTEN)
Lens       4386                    local_user   54u  IPv4 0x9f77ff03be1af971      0t0    TCP localhost:59982 (LISTEN)
Lens       4386                    local_user   58u  IPv4 0x9f77ff03d9c9cf49      0t0    TCP localhost:56379 (LISTEN)
Lens       4386                    local_user   59u  IPv4 0x9f77ff03c75e3089      0t0    TCP localhost:60845 (LISTEN)
Lens       4386                    local_user   60u  IPv4 0x9f77ff03cdcd9f49      0t0    TCP localhost:64475 (LISTEN)
Lens       4386                    local_user   61u  IPv4 0x9f77ff03cdd9dc39      0t0    TCP localhost:55640 (LISTEN)
Lens       4386                    local_user   62u  IPv4 0x9f77ff03cdcdd211      0t0    TCP localhost:59965 (LISTEN)
Lens\x20H 42638                    local_user   30u  IPv6 0x9f77ff03bdd7de21      0t0    TCP *:62640 (LISTEN)
Lens\x20H 42638                    local_user   32u  IPv4 0x9f77ff03c75ddf49      0t0    TCP localhost:62641 (LISTEN)
kubectl   43078                    local_user    8u  IPv4 0x9f77ff03cdcec661      0t0    TCP localhost:63175 (LISTEN)
com.docke 49312                    local_user   54u  IPv6 0x9f77ff03cdbae161      0t0    TCP *:5000 (LISTEN)
com.docke 49312                    local_user   71u  IPv6 0x9f77ff03cdb6bae1      0t0    TCP *:80 (LISTEN)
com.docke 49312                    local_user   72u  IPv6 0x9f77ff03cdb6a161      0t0    TCP *:27017 (LISTEN)
com.docke 49312                    local_user   75u  IPv6 0x9f77ff03cdb6a7c1      0t0    TCP *:5432 (LISTEN)
com.docke 49312                    local_user   84u  IPv6 0x9f77ff03cdb6ae21      0t0    TCP *:58043 (LISTEN)
com.docke 49312                    local_user  618u  IPv6 0x9f77ff03af7c3ae1      0t0    TCP *:59213 (LISTEN)
kubectl   55140                    local_user   21u  IPv4 0x9f77ff03c73b7661      0t0    TCP localhost:3003 (LISTEN)
kubectl   55140                    local_user   24u  IPv6 0x9f77ff03cdbae7c1      0t0    TCP localhost:3003 (LISTEN)
dnsmasq   68556                        nobody    5u  IPv4 0x9f77ff03c7372211      0t0    TCP 192.168.0.103:53 (LISTEN)
dnsmasq   68556                        nobody    7u  IPv4 0x9f77ff03c7370dc1      0t0    TCP localhost:53 (LISTEN)
dnsmasq   68556                        nobody    9u  IPv6 0x9f77ff03be1a2ae1      0t0    TCP [fe80:f::36d2:21ce:71aa:3d8d]:53 (LISTEN)
dnsmasq   68556                        nobody   11u  IPv6 0x9f77ff03be435161      0t0    TCP [fe80:e::690b:8101:a070:9c36]:53 (LISTEN)
dnsmasq   68556                        nobody   13u  IPv6 0x9f77ff03be4357c1      0t0    TCP [fe80:8::e4eb:dcff:fe3a:81ed]:53 (LISTEN)
dnsmasq   68556                        nobody   15u  IPv6 0x9f77ff03be435e21      0t0    TCP [fe80:7::e4eb:dcff:fe3a:81ed]:53 (LISTEN)
dnsmasq   68556                        nobody   17u  IPv6 0x9f77ff03c04a1481      0t0    TCP [fe80:6::cfb:6dfe:9d68:2b]:53 (LISTEN)
dnsmasq   68556                        nobody   19u  IPv6 0x9f77ff03c04a1ae1      0t0    TCP [fe80:4::aede:48ff:fe00:1122]:53 (LISTEN)
dnsmasq   68556                        nobody   21u  IPv6 0x9f77ff03c04a0161      0t0    TCP [fe80:1::1]:53 (LISTEN)
dnsmasq   68556                        nobody   23u  IPv6 0x9f77ff03c04a07c1      0t0    TCP localhost:53 (LISTEN)
```
Note that kubectl is listing, "TCP localhost:3003 (LISTEN)" twice - in theory there shuold not be two processes using the same port.  This is because we were running, "port forwarding," at this time, and the process forwards from 127.0.0.1 (localhost in ipv4) as well as [::1] (localhost in ipv6) at the same time.

```
$ kubectl port-forward buysellguess-dep-6867c7cfdf-4nlp5 3003:3000
Forwarding from 127.0.0.1:3003 -> 3000
Forwarding from [::1]:3003 -> 3000
Handling connection for 3003
```

### Cleaned Up Listen to Ports Version

Adding additional flags, "4 -n" cleans up the results and only shows the ipv4 ports, which may be useful if not using ipv6 locally.

```
sudo lsof -i 4 -n -P | grep LISTEN
```