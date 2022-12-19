## Official Haskell Docker Image

https://hub.docker.com/_/haskell

* Haskell slim/buster:

* [here](https://hub.docker.com/layers/library/haskell/slim-buster/images/sha256-d93790604bf5f3d90ab2d79886d10de5682889299128d0b5ec288bd8a40f9c1d?context=explore)




### Writing a Haskell Web Server

https://dev.to/parambirs/how-to-write-a-haskell-web-servicefrom-scratch---part-3-5en6

### Haskell Playground

https://tryhaskell.org/


# Getting Started with Haskell Docker

```
$ docker exec -ti haskellslimbustercontainer /bin/bash
root@aa0e987089b3:/home/volumebindmount#
root@aa0e987089b3:/home/volumebindmount# ghci
GHCi, version 9.4.3: https://www.haskell.org/ghc/  :? for help
ghci>
```

to quit:

```
ghci> :quit
Leaving GHCi.
```

to compile a file:

```
# ghc helloworld.hs
[1 of 2] Compiling Main             ( helloworld.hs, helloworld.o )
[2 of 2] Linking helloworld
root@aa0e987089b3:/home/volumebindmount# ls
Main.hs  helloworld  helloworld.hi  helloworld.hs  helloworld.o
```
to run the file:

```
# ./helloworld
Hello, World!
```

Alternatively, just running the file:

```
# runhaskell helloworld.hs
Hello, World!
```

Load scripts into the interpreter:

```
ghci> :load helloworld
```