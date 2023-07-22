# Elixir On Docker Example

## The Dockerfile


```
RUN mix local.hex --force && mix local.rebar --force 
```

* This line executes two commands to install Hex and Rebar inside the container. 
* Hex is a package manager that's necessary for fetching dependencies in Elixir projects, and Rebar is a build tool for Erlang, which Elixir is built on.
* [Hex Website](https://hex.pm/)

```
CMD ["/home/app/start"]
```

* This line sets the default command to be executed when a container is started from the Docker image. In this case, it's running the start script.
* Within start we have:

```
#!/bin/bash

# install package manager and update with hex and rebar
./hexrebar.sh

# setup a new app
./newapp.sh

# update dependencies
./update.sh

# compile
./compile.sh

# run the application
./runnohalt.sh
```

* Within hexrebar.sh, we have the following, which should have already completed within the Dockerfile:

```
#!/bin/bash

mix local.hex --force
mix local.rebar --force
```

* Wihin newapp.sh we have the following, which sets up a default app within Elixir:

```
#!/bin/bash

mix new my_app
```
* From here, a default project gets created within the directory my_app, which looks like this:

```
my_app/
├── _build/
│   └── dev/
│       └── lib/
│           └── my_app/
│               ├── ebin/
│               │   ├── my_app.app
│               │   └── Elixir.MyApp.beam
│               └── consolidated/
├── lib/
│   └── my_app.ex
├── mix.exs
├── README.md
└── test/
    ├── my_app_test.exs
    └── test_helper.exs
```

* After that point, there is an update with: `mix deps.get`
* Then finally, a `mix run --no-halt` is issued to run the app from within the directory.
* To recap the project directory that gets built:

* _build/: The directory where the compiled code artifacts are stored.
* lib/: The directory where your source code resides.
* mix.exs: The file that defines project configurations and dependencies.
* README.md: A markdown file for project description and instructions.
* test/: The directory where your test files reside.
* In the _build/ directory:
* dev/: Represents the development environment's build artifacts.
* lib/my_app/ebin/: Contains the compiled application bytecode.
* lib/my_app/consolidated/: Used for protocol consolidation.
* And in the lib/ directory:
* my_app.ex: The main file of your application.
* Finally, in the test/ directory:
* my_app_test.exs: An example test file.
* test_helper.exs: A file run before your tests; typically starts the test framework.

## Compile

* The above did not show the results of compiling the application with `mix compile` ... After this is added, we get several beam files added to our project folder:

```
/my_app/_build/dev/lib/my_app/consolidated$ ls
Elixir.Collectable.beam  Elixir.Enumerable.beam  Elixir.Hex.Solver.Constraint.beam  Elixir.IEx.Info.beam  Elixir.Inspect.beam  Elixir.List.Chars.beam  Elixir.String.Chars.beam
```
* Collectable: a protocol for data types that implement a collect/1 function, which allows elements to be collected into a structure.
* Enumerable: a protocol for data types that can be enumerated over (like lists, maps, etc.)
* IEx.Info: a protocol used in the IEx (Interactive Elixir) shell for printing information about data types.
* Inspect: a protocol that converts data types into a readable format, used when printing out values.
* List.Chars: a protocol to convert lists into strings.
* String.Chars: a protocol to convert data into strings.

When you see the Hex.Solver.Constraint module, it's related to the Hex package manager, which is used by Mix to resolve dependencies.

Seeing these .beam files in the consolidated directory means that the application has been compiled and these protocols have been consolidated. 
Protocol consolidation is a step in the Elixir build process that makes protocols faster at the expense of dynamic dispatch.

* After the files have been compiled, one should be able to use docker to interact with the Elixir shell to test functions with:

```
docker exec -it elixir-1p12-phoenix_container iex
```

Once one has activated the shell, one can interact with the hello function by:

```
MyApp.hello()
```

* Howver, this won't work right away, because the program has to be restarted properly after it has been compiled, in order to use iex.
* To do this, we bashed into the container we were working with with the bash.sh script, and then, after making sure we were in the same directory as 
`my_app.ex`, we ran `elixirc my_app.ex` which is the actual app itself.
* From here, after going into iex, we were able to call MyApp.hello() which showed the expected results of the function.

```
root@56afddc3a175:/home/app/my_app/lib# ls
my_app.ex
root@56afddc3a175:/home/app/my_app/lib# elixirc my_app.ex
root@56afddc3a175:/home/app/my_app/lib# iex
Erlang/OTP 24 [erts-12.3.2.13] [source] [64-bit] [smp:4:4] [ds:4:4:10] [async-threads:1] [jit]

Interactive Elixir (1.12.3) - press Ctrl+C to exit (type h() ENTER for help)
iex(1)> MyApp.hello()
:world
```

* So the important take-away here was that we had to run elixirc to get the application running from within the conainer, after the container had been running.
* For whatever reason, using the `mix` tool did not work where `elixirc` did.

* While elixirc can be used to compile individual Elixir files, for larger projects, you would typically use the mix tool, which can manage project dependencies, 
run tests, generate documentation, and more, as well as compiling the project. mix uses elixirc under the hood to perform the actual compilation.
* So it's possible that our compile.sh script would still work, but it has to be done in the correct location.

While elixirc can be used to compile individual Elixir files, for larger projects, you would typically use the mix tool, which can manage project dependencies, 
run tests, generate documentation, and more, as well as compiling the project. mix uses elixirc under the hood to perform the actual compilation.
