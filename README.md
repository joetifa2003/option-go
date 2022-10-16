# Option

A generic optional type library for golang like the rust option enum

[![GoReportCard example](https://goreportcard.com/badge/github.com/joetifa2003/option-go)](https://goreportcard.com/report/github.com/joetifa2003/option-go)
[![GoDoc reference example](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/joetifa2003/option-go)

- [Option](#option)
- [Installation](#installation)
- [Create an optional](#create-an-optional)
  - [With a value (Some)](#with-a-value-some)
  - [Without a value (None)](#without-a-value-none)
- [Unwrap an optional](#unwrap-an-optional)
  - [With a value (Some)](#with-a-value-some-1)
  - [Without a value (None)](#without-a-value-none-1)
- [Set an optional](#set-an-optional)
  - [Some](#some)
  - [None](#none)

# Installation

```
go get github.com/joetifa2003/option-go
```

# Create an optional

## With a value (Some)

```go
optional := option.NewSome(10)
fmt.Println(optional.IsSome())   // true
fmt.Println(optional.OrElse(14)) // 10 because it has a value of 10
```

## Without a value (None)

```go
optional := option.NewNone[int]()
fmt.Println(optional.IsNone())   // true
fmt.Println(optional.OrElse(14)) // 14 because it has no value
```

# Unwrap an optional

## With a value (Some)

```go
optional := option.NewSome("John Doe")
value, isSome := optional.Unwrap()
fmt.Println(value) // John Doe
fmt.Println(isSome) // true
```

## Without a value (None)

```go
optional := option.NewNone[string]()
value, isSome := optional.Unwrap()
fmt.Println(value) // "" zero value for type string
fmt.Println(isSome) // false
```

# Set an optional

## Some

```go
optional := option.NewNone[string]()
value, isSome := optional.Unwrap()
fmt.Println(value) // "" zero value for type string
fmt.Println(isSome) // false

optional.SetSome("John Doe")

value, isSome = optional.Unwrap()
fmt.Println(value) // John Doe
fmt.Println(isSome) // true
```

## None

```go
optional := option.NewSome("John Doe")
value, isSome := optional.Unwrap()
fmt.Println(value) // John Doe
fmt.Println(isSome) // true

optional.SetNone()

value, isSome = optional.Unwrap()
fmt.Println(value) // "" zero value for type string
fmt.Println(isSome) // false
```
