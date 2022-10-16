# Option

A generic optional type library for golang like the rust option enum

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
fmt.PrintLn(optional.isSome()) // true
fmt.PrintLn(optional.OrElse(14)) // 10 because it has a value of 10
```

## Without a value (None)

```go
optional := option.NewNone[Int]()
fmt.PrintLn(optional.isNone()) // true
fmt.PrintLn(optional.OrElse(14)) // 14 because it has no value
```

# Unwrap an optional

## With a value (Some)

```go
optional := option.NewSome("John Doe")
value, isSome := optional.Unwrap()
fmt.PrintLn(value) // John Doe
fmt.PrintLn(isSome) // true
```

## Without a value (None)

```go
optional := option.NewNone[string]()
value, isSome := optional.Unwrap()
fmt.PrintLn(value) // "" zero value for type string
fmt.PrintLn(isSome) // false
```

# Set an optional

## Some

```go
optional := option.NewNone[string]()
value, isSome := optional.Unwrap()
fmt.PrintLn(value) // "" zero value for type string
fmt.PrintLn(isSome) // false

optional.SetSome("John Doe")

value, isSome = optional.Unwrap()
fmt.PrintLn(value) // John Doe
fmt.PrintLn(isSome) // true
```

## None

```go
optional := option.NewSome("John Doe")
value, isSome := optional.Unwrap()
fmt.PrintLn(value) // John Doe
fmt.PrintLn(isSome) // true

optional.SetNone()

value, isSome = optional.Unwrap()
fmt.PrintLn(value) // "" zero value for type string
fmt.PrintLn(isSome) // false
```
