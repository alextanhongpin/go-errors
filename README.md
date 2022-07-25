# go-errors

Representing errors in Go as data structures.

## TODO

View presentation [here](https://alextanhongpin.github.io/go-errors/#1)


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


## How to implement it

1. Define your error kinds
2. Define your error codes in the errors.toml
3. Import and init the errors
4. Call them
5. Map your errors to http error codes

## Which format to store errors (toml, json, or yaml?)

Personally TOML is the best, even though there's repetition of keys. They are easier to diff too line by line so you know which translations is missing, unlike nested json/yaml.


https://engineering.zalando.com/posts/2021/04/modeling-errors-in-graphql.html

# Thoughts

- separate domain errors from usecase errors, e.g createUser.passwordTooShort, password too short is from value object. This provides a *hint from where the error originates*
- will it lead to repetition? The same error may appear in two different usecase, albeit with the same message. But tying usecase to error does add a lot of clarity. One disadvantage is when the usecase changes.
- separate store errors from usecase errors. E.g user not found etc. Errors from repository should be handled
- catch all errors should be unknown, and not internal server error. all unknown errors needs to be handled.
- errors should be handled layer by layer. That is, we should not just propagate the erroe without khandling them. So errors in the repository layer should be handled in the usecase for clarity.
- another idea is to add tagging capability to errors, so that two similar error are treated distinctly. For example, user.emailExists can be returned in the usecase create user or update email. We can add the tag createUser and updateUserEmail so that both points to the same error, but is distinct. So errors can be declared as ErrCreateUserEmailExists and ErrUpdateUserEmailExists respectively.


