## Official Haskell Docker Image

https://hub.docker.com/_/haskell

* Haskell slim/buster:

* [here](https://hub.docker.com/layers/library/haskell/slim-buster/images/sha256-d93790604bf5f3d90ab2d79886d10de5682889299128d0b5ec288bd8a40f9c1d?context=explore)


### Writing a Haskell Web Server

https://dev.to/parambirs/how-to-write-a-haskell-web-servicefrom-scratch---part-3-5en6

### Haskell Playground

https://tryhaskell.org/


### Getting Started with Haskell Docker

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

### Some Examples

```
ghci> [ x*2 | x<-[ 1..10 ] ]
[2,4,6,8,10,12,14,16,18,20]
```

* We can run a haskell pogram with:

```
main :: IO ()
main = do 
   let var1 = 10 
   let var2 = 6 
   putStrLn "The Addition of the two numbers is:" 
   print( var1 + var2 )
```
* Run it:

```
>> runhaskell Additionoperator.hs
The Addition of the two numbers is:
16
```

### Haskell Cabal

* [Haskell Cabal](https://www.haskell.org/cabal/)

https://www.haskell.org/cabal/download.html

* First install ghcup with:

[ghcup](https://www.haskell.org/ghcup/)

```
curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh
```

* If we run the above, we will get:

```
Welcome to Haskell!

This script can download and install the following binaries:
  * ghcup - The Haskell toolchain installer
  * ghc   - The Glasgow Haskell Compiler
  * cabal - The Cabal build tool for managing Haskell software
  * stack - A cross-platform program for developing Haskell projects (similar to cabal)
  * hls   - (optional) A language server for developers to integrate with their editor/IDE

ghcup installs only into the following directory,
which can be removed anytime:
  /root/.ghcup
```
* There are certain requirements:

```
System requirements
  Please ensure the following distro packages are installed before continuing (you can exit ghcup and return at any time): build-essential curl libffi-dev libffi6 libgmp-dev libgmp10 libncurses-dev libncurses5 libtinfo5
```
* Note that since this takes a while, if it's important to use cabal, it would be better to actually create an image and run the above commands with a -y default to install the above stack rather than just use a haskell image.

* Once the above is installed, note that:

```
In order to run ghc and cabal, you need to adjust your PATH variable.
To do so, you may want to run 'source /root/.ghcup/env' in your current terminal
session as well as your shell configuration (e.g. ~/.bashrc).
```
* So basically, we have to add cabal to our path.  This can be done with:

```
root@d23b4b52e8c6:~# echo "source /root/.ghcup/env" >> ~/.bashrc
root@d23b4b52e8c6:~# cat ~/.bashrc
# ~/.bashrc: executed by bash(1) for non-login shells.

# Note: PS1 and umask are already set in /etc/profile. You should not
# need this unless you want different defaults for root.
# PS1='${debian_chroot:+($debian_chroot)}\h:\w\$ '
# umask 022

# You may uncomment the following lines if you want `ls' to be colorized:
# export LS_OPTIONS='--color=auto'
# eval "`dircolors`"
# alias ls='ls $LS_OPTIONS'
# alias ll='ls $LS_OPTIONS -l'
# alias l='ls $LS_OPTIONS -lA'
#
# Some more alias to avoid making mistakes:
# alias rm='rm -i'
# alias cp='cp -i'
# alias mv='mv -i'
source /root/.ghcup/env
```
* But of course we also need to run this as ```source /root/.ghcup/env```
* Once that is added we can do:

```
root@d23b4b52e8c6:/home/volumebindmount/scotty-webserver# cabal init

Guessing dependencies...

Generating LICENSE...
Warning: unknown license type, you must put a copy in LICENSE yourself.
Generating CHANGELOG.md...
Generating app/Main.hs...
Generating scotty-webserver.cabal...

Warning: no synopsis given. You should edit the .cabal file and add one.
You may want to edit the .cabal file and add a Description field.
```
* So basically, we need to edit the .cabal file, meaning scott-webserver.cabal.
* Looking at the files we have:

```
root@d23b4b52e8c6:/home/volumebindmount/scotty-webserver# ls
CHANGELOG.md  app  scotty-webserver.cabal
```
* Looking in to the actual app directory, we see a helloworld file:

```
root@d23b4b52e8c6:/home/volumebindmount/scotty-webserver/app# cat Main.hs
module Main where

main :: IO ()
main = putStrLn "Hello, Haskell!"
```
* We can replace the helloworld app with a webserver app:

```
module Main where

{-# LANGUAGE OverloadedStrings #-}
import Web.Scotty
import Network.HTTP.Types

main = scotty 3000 $ do
  get "/" $ do                         -- handle GET request on "/" URL
    text "This was a GET request!"     -- send 'text/plain' response
  delete "/" $ do
    html "This was a DELETE request!"  -- send 'text/html' response
  post "/" $ do
    text "This was a POST request!"
  put "/" $ do
    text "This was a PUT request!"
```
* We then have to update our build-depends file:

```
    -- LANGUAGE extensions used by modules in this package.
    -- other-extensions:
    build-depends:    base ^>=4.16.4.0
                , scotty
                , http-types
    hs-source-dirs:   app
    default-language: Haskell2010
```

# References

* [How to Write Haskell Web Service from Scratch](https://dev.to/parambirs/how-to-write-a-haskell-web-servicefrom-scratch---part-3-5en6)
* [Install ghcup](https://www.haskell.org/ghcup/)