# go-errors

Representing errors in Go as data structures.

## TODO

View presentation [here](https://go-talks.appspot.com/github.com/alextanhongpin/go-errors/main.slide#1)


## Motivation

While it is more common in go (or any other languages) to represent errors as code, it is sometimes easier to just represent them as plain data structures (aka maps).

There are several benefits to this approach:
- it is easier to read and understand
- it is easier to serialize and deserialize
- it is easier to transform plain data structures to code than vice versa
- it is easier to modify the data structure than to modify the code
- plain data structures can be automated and tested
- the same implementation can be used in another language, making code-sharing possible
- data can be changed by just swapping/conditionally loading files (useful for translation)
- runtime-check is possible, so there's no compromise	in performance
- the representation of data structures is simpler - you actually type less code, and less code means less bugs


## What's your usecase?

- [x] to be able to categorize errors
- [x] to be able to translate errors
- [x] to be able to match errors by the error code/kind
- [x] to be able to include meaningful context (aka data associated with the error)
- [x] to be able to serialize errors to json (for logging)
