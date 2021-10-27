# SSL

* SSL is a type of [Public Key Certificate](https://en.wikipedia.org/wiki/Public_key_certificate).

At its most basic...

> In cryptography, a public key certificate, also known as a digital certificate or identity certificate, is an electronic document used to prove the ownership of a public key. The certificate includes information about the key, information about the identity of its owner (called the subject), and the digital signature of an entity that has verified the certificate's contents (called the issuer). If the signature is valid, and the software examining the certificate trusts the issuer, then it can use that key to communicate securely with the certificate's subject.

SSL certificates specifically deal with servers and clients on the internet.

##

> In TLS (an updated replacement for SSL), a server is required to present a certificate as part of the initial connection setup. A client connecting to that server will perform the certification path validation algorithm:

> The subject of the certificate matches the hostname (i.e. domain name) to which the client is trying to connect;
The certificate is signed by a trusted certificate authority.

> A TLS server may be configured with a self-signed certificate. When that is the case, clients will generally be unable to verify the certificate, and will terminate the connection unless certificate checking is disabled.

As per the applications, SSL certificates can be classified into three types:[3]

Domain Validation SSL;
Organization Validation SSL;
Extended Validation SSL.

## Certificate Authority

> In the X.509 trust model, a certificate authority (CA) is responsible for signing certificates. These certificates act as an introduction between two parties, which means that a CA acts as a trusted third party. A CA processes requests from people or organizations requesting certificates (called subscribers), verifies the information, and potentially signs an end-entity certificate based on that information. To perform this role effectively, a CA needs to have one or more broadly trusted root certificates or intermediate certificates and the corresponding private keys. CAs may achieve this broad trust by having their root certificates included in popular software, or by obtaining a cross-signature from another CA delegating trust. 
## The Certs Directory /etc/ssl/certs

Within debian the certs directory is at, "/etc/ssl/certs"

> For system-wide use, OpenSSL should provide you /etc/ssl/certs and /etc/ssl/private. The latter of which will be restricted 700 to root:root.
> If you have an application that doesn’t perform initial privilege separation from root, then it might suit you to locate them somewhere local to the application with the relevantly restricted ownership and permissions.
> This seems to be historical OpenSSL convention, not formally standardized.

### Error Setting Certificate Verify Locations

```
curl: (77) error setting certificate verify locations:  CAfile: /etc/ssl/certs/ca-certificates.crt CApath: /etc/ssl/certs
```

Even though it's a curl error, the issue has nothing to do with the URL itself, it has to do with the certs.

###  ca-certificates package

* Short Answer:

> ca-certificates is a deb package that contains certificates provided by the Certificate Authorities.

* Long Answer:

> All digital certificates need to be updated, replaced and changed every now and then. This package holds the updated versions of the ca-certificates that are common to everyone. It simplifies the process of downloading certificates and importing them manually. When you install the ca-certificates package, you also get an updater. You can run it manually or add it to a cron job. You can find more information in the links below.

Further Reading:

http://manpages.ubuntu.com/manpages/xenial/man8/update-ca-certificates.8.html
https://launchpad.net/ubuntu/xenial/+package/ca-certificates
https://launchpad.net/ubuntu/+source/ca-certificates/+changelog

#### Detailed Error Diagnosis on ca-certificates package

```
curl: (77) error setting certificate verify locations:  CAfile: /etc/ssl/certs/ca-certificates.crt CApath: /etc/ssl/certs
```

First, what is "CApath" and "CAfile" - if we check out curl:

```
--capath <CA certificate directory>
       (SSL)  Tells  curl to use the specified certificate directory to verify the peer. The certificates
       must be in PEM format, and if curl  is  built  against  OpenSSL,  the  directory  must  have  been
       processed  using  the  c_rehash  utility  supplied with OpenSSL. Using --capath can allow OpenSSL-
       powered curl to make SSL-connections much more efficiently than using  --cacert  if  the  --cacert
       file contains many CA certificates.

       If this option is used several times, the last one will be used.
```
Looking at the file, "*.crt" by going into the /etc/ssl/certs folder:

```
# pwd
/etc/ssl/certs
# ls | grep crt
ca-certificates.crt
```

We can see that the ca-certificates.crt file does exist.

If we look at the file specified, it shows the following, which is essentially a long string of certificates:

```
head ca-certificates.crt
-----BEGIN CERTIFICATE-----
... (truncated)
```

But what is the ca-certificates.crt file?

> Certificates are added to the CA certificate database using the update-ca-certificates command. This is a shell script that scans the source certificate directories and adds any certificates found to the certificate bundle (/etc/ssl/certs/ca-certificates.crt) as well as creating a symlink in /etc/ssl/certs to the certificate. The ca-certificates.crt file is a concatenation of certificates, each in PEM format. The script doesn't convert any certificate formats, therefore it assumes that all certificates in the source folders are in PEM format with a .crt file extension.

Evidently there is a program, "update-ca-certificates" which does the updating.

Evidently we can run this within a debian VM:

```
# update-ca-certificates
Updating certificates in /etc/ssl/certs...
0 added, 0 removed; done.
Running hooks in /etc/ca-certificates/update.d...
done.
```
We can find this script on the filesystem with:

```
# find / -name update-ca-certificates
/var/lib/dpkg/triggers/update-ca-certificates
/usr/sbin/update-ca-certificates
```
Looking at various resources online, this appears to potentially be a missing package problem.

* [Debian Bullseye Update CA Certificates](https://manpages.debian.org/bullseye/ca-certificates/update-ca-certificates.8.en.html)

```
apt-get install ca-certificates
```
However from this we get, "ca-certificates is already the newest version (20210119)."

What is this package exactly? All digital certificates need to be updated, replaced and changed every now and then. This package holds the updated versions of the ca-certificates that are common to everyone.

So what is "ca-certificates.crt"?

It is basically a PEM formatted file which lists all of the certificates within it, which when "update-ca-certificates" is ran, it updates particular types of certificate stores into ca-certificates.crt. This file, "ca-certificates.crt" is just a long list of all of your trusted CA's concatenated together

Note that even if we attempt to curl Google within a Dockerfile, we get an error:

```
=> ERROR [5/11] RUN curl https://www.google.com/                                                                                                                     0.2s
------
 > [5/11] RUN curl https://www.google.com/:
#13 0.190   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
#13 0.190                                  Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
#13 0.217 curl: (77) error setting certificate verify locations:  CAfile: /etc/ssl/certs/ca-certificates.crt CApath: /etc/ssl/certs
```
Within our base.Dockerfile, if we attempt to install the, "update-ca-certificates" package just to see if it will help:

```
RUN apt-get update && apt-get install -y --no-install-recommends update-ca-certificates
```
We get an error:  

```
#14 2.465 Reading state information...
#14 2.544 E: Unable to locate package update-ca-certificates
------
executor failed running [/bin/bash -euo pipefail -c apt-get update && apt-get install -y --no-install-recommends update-ca-certificates]: exit code: 100
```
What if we just try running, "update-ca-certificates" assuming it's already installed?

```
 > [6/12] RUN /usr/sbin/update-ca-certificates --verbose:
#14 0.203 /bin/bash: line 1: /usr/sbin/update-ca-certificates: No such file or directory
```
Going back into our shell, we notice that we have not been running as bash, but rather as the regular /bin/sh terminal. So activating /bin/bash and running the command:

```
root@9619f95feb7b:/# curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-${SSM_ENV_VERSION} &&     chmod +x /usr/local/bin/ssm-env
curl: (22) The requested URL returned error: 404 Not Found
```
We see there is a 404 error.

So, switching shells and then running the command:

```
Dockerfile...

SHELL ["/bin/sh"]

ARG SSM_ENV_VERSION=0.1.4
RUN curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-${SSM_ENV_VERSION} && \
    chmod +x /usr/local/bin/ssm-env

Output...

 > [6/10] RUN curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-0.1.4 &&     chmod +x /usr/local/bin/ssm-env:
#14 0.190 /bin/sh: 0: cannot open curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-${SSM_ENV_VERSION} &&     chmod +x /usr/local/bin/ssm-env: No such file

```
Now we get the output that the file does not actually exist. So, we can add in, "RUN touch /usr/local/bin/ssm-env"

-Prior to adding this on the base.Dockerfile, we get "# find / -name ssm-env" leading to "not found"
-After adding this on base.Dockerfile we are able to find the file at /usr/local/bin/ssm-env

However, we still get the error:

```
> [7/12] RUN curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-0.1.4:
#14 0.185 /bin/sh: 0: cannot open curl -fsSLo /usr/local/bin/ssm-env https://bb-tools.s3-us-west-1.amazonaws.com/ssm-env-linux-${SSM_ENV_VERSION}: No such file
```
If we switch the shell back to bash we still get the previous error:

```
curl: (77) error setting certificate verify locations
```
So what seems to be happening (clue found on stackoverflow) is that curl expects the file to be in one location, but it's actually in another location.

Where the file actually is:

-/etc/ssl/certs/ca-certificates.crt

Where curl expects it:

-/etc/ssl/certs/ca-certificates.crt

... which is the same.

So part of one of the below Stackoverflow issues says that this is an issue with not having ca-certificates, (the package) installed.

Of course as discussed above, if we try to install it, we get a, "no installation candidates" message.

There is a [long discussion since 2017 on docker debian github](https://github.com/debuerreotype/docker-debian-artifacts/issues/15#issuecomment-590919835) going over the pro's and cons of adding ca-certificates, and someone pointing out that there is an alternative image which includes ca-certificates, curl, netbase, wget, and gnupg (with dirmngr):

* buildpack-deps:buster-curl

However another discussion on node specifically states that -slim will not even be shipped with curl (hence why we are installing curl in our base.Dockerfile).

Finally, running a combined command:

```
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
```
Makes the base.Dockerfile run.  However, it's important to understand whether we really need to pre-create the file for /user/local/bin/ssm-env 

If we comment this out, it effects nothing, so evidently this really was an update and ca-certificates issue.

Further, it is important to understand which version of ca-certificates we are installing.

Taking a look at our container after the above was performed, confusingly the command results in, 'not found':

```
# ca-certificates
command not found
```

However, installing ca-certificates gets the following:

```
# apt-get install ca-certificates
Reading package lists... Done
Building dependency tree... Done
Reading state information... Done
ca-certificates is already the newest version (20210119).
```
So even though the command shows as not found, the installation works and there is actually a version number.

Further confusingly, every reference to ca-certificates after installation is a directory:

```
# find / -name ca-certificates
/usr/local/share/ca-certificates
/usr/share/doc/ca-certificates
/usr/share/ca-certificates
/etc/ca-certificates

# head /usr/local/share/ca-certificates
head: error reading '/usr/local/share/ca-certificates': Is a directory

(empty directory)

# head /usr/share/doc/ca-certificates
head: error reading '/usr/share/doc/ca-certificates': Is a directory

ls: copyright examples

examples:

ca-certificates-local

ls: debian local

~

# head /usr/share/ca-certificates
head: error reading '/usr/share/ca-certificates': Is a directory
# head /etc/ca-certificates
head: error reading '/etc/ca-certificates': Is a directory

```
So it appears to be a cascading structure of folders, with a Mozilla freeware copyright included 

If we specify the version it does indeed work out once again, using the RUN command:

```
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates=20210119
```
So installing this together with all of the other packages results in success as well, our final form of the package installation looks like:

```

RUN apt-get update                                                                                              && \
    apt-get install -y --no-install-recommends                                                                     \
        fonts-lato=2.*                                                                                             \
        git=1:2.30.*                                                                                               \
        imagemagick=8:6.9.*                                                                                        \
        less=551-2                                                                                                 \
        librsvg2-bin=2.50.*                                                                                        \
        vim=2:8.2.*                                                                                                \
        ca-certificates=20210119                                                                                && \
    apt-get clean                                                                                               && \
    rm -rf /var/lib/apt/lists/*

```

## The SSM Environment /user/local/bin/ssm-env

# Resources

* [Serverfault - SSL Certificates](https://serverfault.com/questions/62496/ssl-certificate-location-on-unix-linux)
* [What Certificates Format is ca-certificates?](https://unix.stackexchange.com/questions/514136/what-certificate-format-does-usr-local-share-ca-certificates-accept)
* [Debian Bullseye Update CA Certificates](https://manpages.debian.org/bullseye/ca-certificates/update-ca-certificates.8.en.html)
* [Stackoverflow - where curl expects the path to be](https://stackoverflow.com/questions/3160909/how-do-i-deal-with-certificates-using-curl-while-trying-to-access-an-https-url/13400988#13400988)
* [long discussion since 2017 on docker debian github](https://github.com/debuerreotype/docker-debian-artifacts/issues/15#issuecomment-590919835)