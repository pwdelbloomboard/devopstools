# About Python Maps

* Within functional programming, computations are done by taking input arguments, not modifying the input arguments or changing a program state in any way, and then output concrete values.
* Functions can in theory be developed in isolation, tested in isolation, and are very predictable because there are no state changes.
* Lists, arrays, other iterables represent the data along with a set offunctions that operate on the data and transform it. There are different techniques:

1. Mapping, applying a transformation function to an iterable to produce a new iterable.
2. Filtering applies a predicate or boolean function to generate a new iterable.
3. Reducing consists of applying a reduction function to an iterable to produce a single cumulative value.

* These are all fundamental parts of Python - map(), filter() and reduce().

### Maps Disambiguation, Map-Like Data Structure, vs. Python map()

* To clarify, the map() function in Python is a higher-order function that takes a function object and one or more iterables as its arguments, then applies the given function to each item in the iterable(s) and returns an iterator.
* However, there is also a, "map formatted dictionary," or perhaps better thought of, a "dictionary with a map-like structure," which is a data structure, a way of forming a dictionary to represent a collection of something, like the following:

```
operation_map = {
    'add': {
        'function': lambda x, y: x + y,
        'description': 'Adds two numbers together.',
    },
    'subtract': {
        'function': lambda x, y: x - y,
        'description': 'Subtracts the second number from the first number.',
    },
    'multiply': {
        'function': lambda x, y: x * y,
        'description': 'Multiplies two numbers together.',
    },
    'divide': {
        'function': lambda x, y: x / y,
        'description': 'Divides the first number by the second number.',
    },
}

```