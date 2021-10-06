
```
 kl postgres-instance -c postgres-fix 
Running postgres init
whoami: postgres
total 120K
drwxrwx--- 1 postgres postgres 4.0K Sep 17 22:21 base
drwxrwx--- 1 postgres postgres 4.0K Sep 20 15:50 global
drwxrwx--- 1 postgres postgres 4.0K Sep 17 20:47 pg_commit_ts
drwxrwx--- 1 postgres postgres 4.0K Sep 17 20:47 pg_dynshmem
...
-rw-rwx--- 1 postgres postgres    3 Sep 17 20:47 PG_VERSION
...
pg_resetwal: error: could not open file "PG_VERSION" for reading: Operation not permitted
```

Postgres Docker File has:

```
# explicitly set user/group IDs
RUN set -eux; \
	groupadd -r postgres --gid=999; \
# https://salsa.debian.org/postgresql/postgresql-common/blob/997d842ee744687d99a2b2d95c1083a2615c79e8/debian/postgresql-common.postinst#L32-35
	useradd -r -g postgres --uid=999 --home-dir=/var/lib/postgresql --shell=/bin/bash postgres; \
```


From Askubuntu:

```
> who is the user 999? Is it just custom to him or a code used universally?

In Ubuntu and the Ubuntu family flavours, the numeric user ID in live sessions is 999. (The literal user ID is ubuntu, kubuntu, lubuntu ... xubuntu but in all these cases the numeric user ID is 999.)

When you boot from a USB drive, 'Try Ubuntu', you boot into a live session.

You can check with the command

grep 999 /etc/group
and you will find the user if you run a live session. Otherwise you will probably not find anything via that command.

The operating system used for the Docker tutorial is probably an installed system. And there seems to be a user ID with the numeric 999. It is possible to create such a user ID, but in an installed Ubuntu system there is no standard user with that numeric user ID.
```

https://stackoverflow.com/questions/55241474/why-docker-compose-creates-directories-files-with-usergroup-999999