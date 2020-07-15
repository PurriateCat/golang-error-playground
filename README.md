![Go](https://github.com/PurriateCat/golang-error-playground/workflows/Go/badge.svg)
![Lint](https://github.com/PurriateCat/golang-error-playground/workflows/Lint%20Code%20Base/badge.svg?branch=master)

# golang-error-playground
Some small golang programs that shows how to use and handle errors.

## Content
File | Description
--- | --- 
`custom_struct_error.go` | Implement custom error using structs
`custom_http_error.go` | Example http error implementation
`errors_package.go` | Errors using the native errors package
`type_assertion.go` | Checking for error type by assertions

## Playing around
You can run every program by its own using `go run [FILE]`.

Examples:
```
$ go run custom_struct_error.go
$ go run custom_http_error.go
$ go run errors_package.go
$ go run type_assertion.go
```
