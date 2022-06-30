https://go.dev/ref/spec#Identifiers

* Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

https://go.dev/ref/spec#Predeclared_identifiers


> Bootstrapping

> Current implementations provide several built-in functions useful during bootstrapping. These functions are documented for completeness but are not guaranteed to stay in the language. They do not return a result.

> print      prints all arguments; formatting of arguments is implementation-specific
> println    like print but prints spaces between arguments and a newline at the end

> Implementation restriction: print and println need not accept arbitrary argument types, but printing of boolean, numeric, and string types must be supported.

