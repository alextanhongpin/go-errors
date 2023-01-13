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

# Errors at different layers


Each layers has their own goal.

Repository
- the end goal is to map errors to sentinel errors, not custom errors yet
- also, provide assertions functions for other package to use, e.g IsNotFoundError rather than allowoing them to import the error to check
- for find one, map errors to not found
- for update and delete, optionally return the not found if no entries were updated or deleted
- for insert, return errors regarding constraints, such as uniqueness etc.
- the errors should be handled by the caller and converted to custom errors accordingly 
- on second thought, we can just return custom errors so that any caller dont have to map the errors. repository is a part of domain after all.
- all errors not converted to custom errors will become internal server error
- storage is for table specific, repository is domain specific. storage error is generic, hence return sentinel error. repository error is specific, hence domain errors.

usecase
- the end goal is to return clear errors per usecase
- domain errors are generic, e.g. UserExistsError could be used by different usecase. However, we might want to mark this error under CreateUserError.UserExists for the usecase CreateUser. This makes the CreateUser usecase explicit, since we can return User|CreateUserError 
- for find one, there is not found error, by condition 
- for create, there can be constraints during creation
- for update or delete, we coukd return not found error
- all is subject to business logic error, e.g. can create/update/delete/read
- ~we can skip repo errors since they are handled.~ actually usecase should handle the conversion. the top layers should always handle thr lower layer errors. This is mainly because go is bad at extending errors, until the errors.Join is merged into 1.20. 
- services should return boolean, then usecase can decide which errors to return
- if the service must return error, then the errors must be wrapped. It is preferable if thr service returns only one error, or also another error object 
- validation errors should not be part of this (required fields etc)
- multi errors is still a pain 

api
- api should map the domain errors to rest http errors
- api layer is also responsible for sending the errors to error management such as sentry
- api layer should also log the errors
- api layer can include additional metadatas such as url path and request id of the errors


# Thoughts

- separate domain errors from usecase errors, e.g createUser.passwordTooShort, password too short is from value object. This provides a *hint from where the error originates*
- will it lead to repetition? The same error may appear in two different usecase, albeit with the same message. But tying usecase to error does add a lot of clarity. One disadvantage is when the usecase changes.
- separate store errors from usecase errors. E.g user not found etc. Errors from repository should be handled
- catch all errors should be unknown, and not internal server error. all unknown errors needs to be handled.
- errors should be handled layer by layer. That is, we should not just propagate the erroe without khandling them. So errors in the repository layer should be handled in the usecase for clarity.
- another idea is to add tagging capability to errors, so that two similar error are treated distinctly. For example, user.emailExists can be returned in the usecase create user or update email. We can add the tag createUser and updateUserEmail so that both points to the same error, but is distinct. So errors can be declared as ErrCreateUserEmailExists and ErrUpdateUserEmailExists respectively. (update: tagging is good)
- how many errors should a function return? When working with a large function, there's a possibility of branching, and they often lead to different errors returned from the function. Having a single error per function is usually desirable, because it simplifies error mapping (when mapping from error to another). Currently it's not possible to create error unions with golang.
