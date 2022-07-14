# Structs and Interfaces

## Interfaces

* [About Interfaces](https://go.dev/doc/effective_go#interfaces)

> Interfaces provide a way to specify the behavior of an object. If something can do **this** it can be used **here**.

> For example, Fprintf can generate output to anything with a Write method.
> String methods can implement custom printers.

## Quick Rundown on Maps, Structs, Interfaces, Slices, Arrays, Types

* Maps are like Dicts in Python
* Structs are like Objects in Python, where you can define fields.
* Interface 
* In Go you have Arrays and Slices. Slices can grow dynamically, just like a List in Python.
* Pretty much always use a Slice in Go, because with Arrays you have to define the Size.  E.g. Slices are dynamically re-sized.
* Types are a specialized Struct. You can define how to compare Types, or you can have specific functions which require a specific Type to operate on, which do not accept other Types as inputs.


## Pointers in Functions, *WhateverThingyArgs


## Further Expansion on Maps, Type Assertions, Interfaces, Interface Conversions, 

https://go.dev/doc/effective_go#maps

https://go.dev/ref/spec#Type_assertions

https://go.dev/doc/effective_go#interfaces

https://go.dev/doc/effective_go#interface_conversions 


