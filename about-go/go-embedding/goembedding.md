# Go Embedding

## //go:embed

* //go:embed directive allows you to include contents of arbitrary files and directories in your go application.
* To test out using //go:embed, we do a simple file which goes through and reads from version.txt to grab version information and prints it out:

```
package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	Version string = strings.TrimSpace(version)
	//go:embed version.txt
	version string
)

func main() {
	fmt.Printf("Version %q\n", Version)
}
```
* The results of running this build are:

```
go run version.go
Version "0.0.1"
```

* However, the there was a linting error within VSCode:

```
gopls requires a module at the root of your workspace.
You can work with multiple modules by opening each one as a workspace folder.
Improvements to this workflow will be coming soon, and you can learn more here:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go
```

* Allegedly the way to fix this is to add the following to the VSCode Settings.Json:

```
    "gopls": {
        "ui.completion.usePlaceholders": true,
    }
```

* However, when we initialize a go module within the same folder as the version.go program we're working with, that clears the above linting error and replaces it with:

```
could not import embed (cannot find package "embed" in any of 
	/Users/.../.asdf/installs/golang/1.15.8/go/src/embed (from $GOROOT)
	/Users/.../.asdf/installs/golang/1.15.8/packages/src/embed (from $GOPATH))compiler
```
* This is in part because we may not have embed installed within our local machine, vs. on our remote machine on which go is running, so we can reasonably ignore this.
* That being said, the actual package we're trying to install is [go-embed](https://github.com/golang/go/blob/master/src/embed/embed.go) which can be found at that link.

* Evidently, we're not using the proper version of Go. The default version of Go running on my machine is 1.15.
* That being said, there is another version of Go being used in another program, which automatically updates the version of go to 1.17 within the go.mod file. How is this posisble?

```
module module.repo.com/whatever

go 1.17
```
More clues maye be found within the [Go Modules Documentation](https://go.dev/doc/modules/gomod-ref), which mentions:

> Each Go module is defined by a go.mod file that describes the module’s properties, including its dependencies on other modules and on versions of Go.

* What is the exact GO module being specifed by our machine?

```
$ go mod tidy
example.com/m imports
	embed: package embed is not in GOROOT (/Users/patrick.delaneybloomboard.com/.asdf/installs/golang/1.15.8/go/src/embed)
```
* What this seems to be saying is that GOROOT is set to /gloang/1.15.8 rather than /1.17.

Looking at the [Go Documentation for modules](https://go.dev/doc/modules/gomod-ref): 

> The module path must uniquely identify your module. For most modules, the path is a URL where the go command can find the code (or a redirect to the code). For modules that won’t ever be downloaded directly, the module path can be just some name you control that will ensure uniqueness. The prefix example/ is also reserved for use in examples like these.

* Within our docker container, the default go version is ```go version go1.17.5 linux/amd64``` whereas within our local machine, the go version is ```go version go1.15.8 darwin/amd64``` - however, the go.mod should be able to select the go version used to run the code.  Since our module path is just, "example.com/" it seems to just default back to our default go version.

* We may be able to reconcile this with a .tool-versions file which includes the golang distribution version:

``` .tool-versions
golang 1.17.5
```
After adding this version, which conforms to the version we're using within our Docker container, and then running, "go tidy," - we get:

```
$ go tidy
No preset version installed for command go
Please install a version by running one of the following:

asdf install golang 1.17.5

or add one of the following versions in your config file at /Users/patrick.delaneybloomboard.com/Projects/devopstools/about-go/go-docker/volumebindmount/goembedtest/.tool-versions
golang 1.15.8
golang 1.17.6
```
So on our local machine we can then run: ```asdf install golang 1.17.5``` which of course goes through and installs the proper version of go.  While the proper versioning does indeed get cleaned up, allowing the "go run" command to run, spitting out the following information:

```
$ go run version.go
Version "0.0.1"
```

* We still get a linting error on VSCode, however everything else seems to work, and we're using the same version of Go on the local machine as on the Docker Container, due to the .tool-versions file.

## //go:embed Conditional

* If we want to include version information conditionally, we can accomplish this with providing one main go file:

```
package main

import (
	"fmt"
	"strings"
)

var Version string = strings.TrimSpace(version)

func main() {
	fmt.Printf("Version %q\n", Version)
}
```

Which requires an input flag to run, with seperate go files:

``` version_dev.go
//go:build !prod
// +build !prod

package main

var version string = "dev"
```
and...
``` version_prod.go
//go:build prod
// +build prod

package main

import (
	_ "embed"
)

//go:embed version.txt
var version string
```





## Resources

* [How to Use Go Embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)
* [Request to Include Conditional Embedding](https://github.com/golang/go/issues/44484)