### Go Cobra Site

https://cobra.dev/

### Using Go Cobra

* To test things out, we used the docker image posted [here within our own repo](/about-go/Dockerfile).

### Installing Cobra and Cobra-CLI

Within the Container built based upon the Dockerfile above, we installed Cobra with:

```
go get -u github.com/spf13/cobra@latest
```

The following can be added to the Dockerfile in order to build it originally with Cobra (alnog with Viper):

```
RUN go get -u github.com/spf13/cobra
RUN go get -u github.com/spf13/viper
```
There is also a seperate tool, called the cobra-cli.

```
go install github.com/spf13/cobra-cli@latest
```

### Following the Cobra User Guide

Go Cobra recommends setting up a project structure.

```
appName/
|
└─── cmd/
|   |
|   └─── add.go
|   └─── your.go
|   └─── commands.go
└─── main.go
```

* For our, "main.go" we used, "commandtest.go" and for one of our commands we used, "tryit.go"

The main.go file, or in our case, commandtest.go is bare, it just calls the location of the cmd folder.

```
package main

import (
  "/home/app/cmd"
)

func main() {
  cmd.Execute()
}
```

However when we tried the above, we got:

* "commandtest.go:4:3: unknown import path "/home/app/cmd": internal error: module loader did not resolve import"

From working with Go previously, we recall that Go has a number of built-in envioronmental variables which serve as the pointers to paths where various packages exist, essentially the, "Go Environment."

GoPath works much like, "Path," within bash, in that multiple directories can be added to it. So starting out, if you do:

```
# go env | grep GOPATH
GOPATH="/go"
```
Similarly:

```
# echo $GOPATH
/go
```

So how do we add on to GOPATH?  Well, we can set it directly, with:

```
GOPATH=go:/home/app
```
However, this doesn't work, and there is a prescribed way to use Gopath, in fact the error directly references us to, "go help gopath"

Reading through, "go help gopath" we see:

> On Unix, the value is a colon-separated string.

> Each directory listed in GOPATH must have a prescribed structure:

> The src directory holds source code. The path below src
determines the import path or executable name.

> The pkg directory holds installed package objects.
As in the Go tree, each target operating system and
architecture pair has its own subdirectory of pkg
(pkg/GOOS_GOARCH).

> The bin directory holds compiled commands.
Each command is named for its source directory, but only
the final element, not the entire path.

So in summary, there must be:

* pkg
* bin
* dir

So, we organize everything within our /home directory under a new folder called, "commandapp" and within this directory we put pkg, bin and dir.

So now our overall structure is:

```
commandApp
|
└─── src/
    |
    └─── cobraApp/
        |
        └─── cmd/
        |   |
        |   └─── tryit.go
        |
        └─── commandtest.go
```
All of the above being said, we may not need this at all, because what the, "GOPATH" appears to be focused on, is creating binaries and installed *packages* which are different than *modules*. Packages appear to be more based upon the system architecture that is being used to run a particular application, for example, linux_amd64 would be a possible package directory, which holds folders designating the target operating system.

Indeed, if we look at the original GOPATH folder structure, we see go/bin, go/pkg, go/src.

* Within bin, we see the installed cobra-cli, envars, and simple, which are all packages we created on this machine.
* Within pkg, we see pkg/mod, pkg/sumdb. Within pkg/mod we see: cache  github.com  golang.org  gopkg.in, in which are different binaries that seem to correspond to Cobra and other things we're not familiar with.
* src is empty.

So backing up, we should instead perhaps use Go Modules, to use a higher layer of abstraction and easier way of laying out an application.

> When using modules, GOPATH is no longer used for resolving imports.
However, it is still used to store downloaded source code (in GOPATH/pkg/mod)
and compiled commands (in GOPATH/bin).

So with modules, the code can be organized as follows:

```
crash/
|    
└─── bang/              // (go code in package bang)
|   |   
|   └─── b.go
|
└─── foo/                   // (go code in package foo)
    |
    └─── f.go
    |
    └─── bar/               // (go code in package bar)
    |   |
    |   └─── x.go
    |
    └─── internal/
    |           |
    |           └─── baz/           // (go code in package baz)
    |                   |
    |                   └─── z.go
    |
    └─── quux/              /// (go code in package main)
            |
            └─── y.go
```

> The code in z.go is imported as "foo/internal/baz", but that
import statement can only appear in source files in the subtree
rooted at foo.

So translating this over to our layout:

```

commandapp/
    |
    └─── cmd/
    |       |
    |       └─── tryit.go         // go code in package cmd
    |
    └─── test/
            |
            └─── commandtest.go   // go code in package test

```
* So basically we can import anything in either *.go file with as "commandApp/cmd/tryit.go" or "commandApp/test/commandtest.go"
* Since the d ocumentation mentions in the above structure that code can be imported as, "foo/internal/baz" and not, "/home/user/go/src/foo/internal/baz" then likely we don't have to use root paths.

Once that structure has been built, we can then restart our structure of the commandtest.go file to import tryit.go from cmd:

```
package main

import (
  "commandapp/cmd/"
)

func main() {
  cmd.Execute()
}
```

* When we try to run the above, we get the error:

```
package command-line-arguments is not a main package
commandtest.go:4:3: unknown import path "/commandapp/cmd": internal error: module loader did not resolve import
```

That being said, for Cobra, the actual recommended structure is a bit different with our, "main.go" app actually being in parallel to the cmd folder:

```
/home/commandapp/
|
└─── cmd/
|   |
|   └─── tryit.go
|
└─── commandtest.go
```

Following [this tutorial](https://go.dev/doc/tutorial/create-module) we have to create a module.  So within the /home/commandapp directory, we did:

```
# go mod init commandapp
go: creating new go.mod: module commandapp
go: to add module requirements and sums:
	go mod tidy
```

So after completing this step, we attempted to go through, "go run," but this resulted in an unknown command error, so instead tried go build and got:

```
# go build commandtest.go
cmd/tryit.go:1:1: expected 'package', found 'EOF'
```

This appears to be stating that within tryit.go we should find at least one package.  Note that we also now have a go.mod file parallel to our commandtest.go file.

So now that we're at this step, we can create a root command within our /home/commandapp/cmd folder. Let's rename tryit.go to root.go and then use:

```
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
```
After running this and then we return to our /home/commandapp folder and do:

```
go build commandtest.go
cmd/root.go:7:2: no required module provides package github.com/spf13/cobra; to add it:
	go get github.com/spf13/cobra
cmd/root.go:8:2: no required module provides package github.com/spf13/viper; to add it:
	go get github.com/spf13/viper
```
So of course we ran the two commands suggested above and finally we are able to run:

```
go build commandtest.go
# commandapp/cmd
cmd/root.go:42:21: undefined: addCmd
cmd/root.go:43:21: undefined: initCmd
```

So now we're getting somewhere, because we're actually getting undefined errors for, "addCmd" and "initCmd".

So rather than using this dummy example, we will use a [concrete Cobra example here](https://github.com/spf13/cobra/blob/master/user_guide.md#example).  

* However we need to modify the code to make it work, starting with "package cmd" at the top.
* After we change the package of this concrete Cobra example above to, "cmd" we get the error:

```
./commandtest.go:8:3: undefined: cmd.Execute
```

Based upon this, we're going to start over completely from a different approach.

### Cobra Tutorial from YouTube

* [Cobra YouTube Tutorial](https://www.youtube.com/watch?v=-tO7zSv80UY)

> from within the project folder, you can initialize a cobra project with the cobra init command:

```
cobra init --pkg-name github.com/example/commandapp
```

This leads to a, "command not found," error, since Go is not available on the command line. To enable Go on the command line, we do:

```
export PATH="~/go/bin:$PATH"
```
However, it's still not available after this, we still get a command not found. So, looking at what happens if we try to envoke the go/bin file directly:

```
~/go/bin/cobra init --pkg-name commandapp
bash: /root/go/bin/cobra: No such file or directory
```
As shown above, within /go/bin we don't have cobra, we have cobra-cli.

It's possible that cobra-cli has replaced, "cobra" on the command line. Here is [user documentation on cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md).

Based upon that user documentation, the command would instead be:

```
 cobra-cli init github.com/example/commandapp
Error: Please run `go mod init <MODNAME>` before `cobra-cli init`
```
However we can see that it wants us to start a mod file first - so we do that, and then run:

```
cobra-cli init cobracommandapp
```
Which, we had changed the actual name that we want to end up with to, "cobracommandapp" since we wanted to distinguish it from the root folder, "cobracommand."

So now looking at the containing folder, we have a completely pre-set application. Note that we don't have to use any kind of URL to define the module, we can simply just use a containing folder.  The, "cobracommandapp," may be used at different points in the application.

In theory, if we run our command, it should run successfully.  However, we get the following, once again:

```
 go run main.go
main.go:7:8: package commandapp/cmd is not in GOROOT (/usr/local/go/src/commandapp/cmd)
```
We should have expected to see the placeholder Cobra description rather than the, "not in GOROOT" error.

Another tutorial:

### Starting Another YouTube Tutorial - More Updated

[Another YouTube Tutorial](https://www.youtube.com/watch?v=Ll-s38JKWi8)

This tutorial above worked.  Here are the commands used to successfully build a go Cobra application with the aboe tutorial:

```
root@a8e4cb13d84c:/home# mkdir commandapp
root@a8e4cb13d84c:/home# cd commandapp/
root@a8e4cb13d84c:/home/commandapp# go mod init commandapp
go: creating new go.mod: module commandapp
root@a8e4cb13d84c:/home/commandapp# cat go.mod

module commandapp

go 1.17

root@a8e4cb13d84c:/home/commandapp# go get -u github.com/spf13/cobra
go get: added github.com/inconshreveable/mousetrap v1.0.0
go get: added github.com/spf13/cobra v1.4.0
go get: added github.com/spf13/pflag v1.0.5
root@a8e4cb13d84c:/home/commandapp# ls
go.mod	go.sum
root@a8e4cb13d84c:/home/commandapp# touch main.go
root@a8e4cb13d84c:/home/commandapp# nano main.go
root@a8e4cb13d84c:/home/commandapp# mkdir cmd
root@a8e4cb13d84c:/home/commandapp# ls
cmd  go.mod  go.sum  main.go
root@a8e4cb13d84c:/home/commandapp# nano main.go
root@a8e4cb13d84c:/home/commandapp# ls
cmd  go.mod  go.sum  main.go
root@a8e4cb13d84c:/home/commandapp# cd cmd
root@a8e4cb13d84c:/home/commandapp/cmd# touch root.go
root@a8e4cb13d84c:/home/commandapp/cmd# ls
root.go
root@a8e4cb13d84c:/home/commandapp/cmd# nano root.go
root@a8e4cb13d84c:/home/commandapp/cmd# cd ..
root@a8e4cb13d84c:/home/commandapp# go build main.go
# commandapp/cmd
cmd/root.go:4:2: imported and not used: "fmt"
cmd/root.go:5:2: imported and not used: "os"
root@a8e4cb13d84c:/home/commandapp# go run main.go
# commandapp/cmd
cmd/root.go:4:2: imported and not used: "fmt"
cmd/root.go:5:2: imported and not used: "os"
root@a8e4cb13d84c:/home/commandapp# ls
cmd  go.mod  go.sum  main.go
root@a8e4cb13d84c:/home/commandapp# nano cmd/root.go
root@a8e4cb13d84c:/home/commandapp# go run main.go
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

root@a8e4cb13d84c:/home/commandapp# nano cmd/root.go
root@a8e4cb13d84c:/home/commandapp# go run main.go
Thanks for testing out our Go Cobra application
```

Basically, our main.go file was the following:

```
package main

import (
  "commandapp/cmd"
)

func main() {
  cmd.Execute()
}
```

While our root.go, resting within /cmd was the following:

```
package cmd

import (
	// "fmt"
	// "os"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	// cfgFile     string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "Command CLI Application A generator.",
		Long: `Thanks for testing out our Go Cobra application`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
```

Basically, one of the keys was initializing the go mod prior to anything else once the folder was created, and then installing go cobra with, "go get ~" after the modfile was created. This allowed the modfile to capture cobra as a dependency.  The mod file is similar in concept to a requirements.txt or a package-lock.json file in node, or a gemfile in ruby on rails. The modfile appears to update automatically as you install dependencies on the fly while building the code, rather than declaratively specifying them within requirements.txt or under some other paradigm. So just because you may have installed a particular dependency on your machine, does not mean it will work in a particular Golang project.



