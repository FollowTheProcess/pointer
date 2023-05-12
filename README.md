# Pointer

[![License](https://img.shields.io/github/license/FollowTheProcess/pointer)](https://github.com/FollowTheProcess/pointer)
[![Go Reference](https://pkg.go.dev/badge/github.com/FollowTheProcess/pointer.svg)](https://pkg.go.dev/github.com/FollowTheProcess/pointer)
[![Go Report Card](https://goreportcard.com/badge/github.com/FollowTheProcess/pointer)](https://goreportcard.com/report/github.com/FollowTheProcess/pointer)
[![GitHub](https://img.shields.io/github/v/release/FollowTheProcess/pointer?logo=github&sort=semver)](https://github.com/FollowTheProcess/pointer)
[![CI](https://github.com/FollowTheProcess/pointer/workflows/CI/badge.svg)](https://github.com/FollowTheProcess/pointer/actions?query=workflow%3ACI)
[![codecov](https://codecov.io/gh/FollowTheProcess/pointer/branch/main/graph/badge.svg)](https://codecov.io/gh/FollowTheProcess/pointer)

***Work fearlessly with pointers in Go***

## Project Description

`package pointer` is a *tiny*, simple and obvious library helping you to always work safely with pointers and avoid the dreaded `nil pointer dereference` 🚀

## Installation

To use this project in your code:

```shell
go get github.com/FollowTheProcess/pointer@latest
```

And then simply:

```go
import "github.com/FollowTheProcess/pointer"
```

## Quickstart

Using `pointer` is so easy, you already know how to do it!

```go
import (
    "fmt"

    "github.com/FollowTheProcess/pointer"
)

func main() {
    var s *string // nil

    // More code here, you've forgotten `s` is nil by now!
    // but pointer has you covered!
    str := pointer.Or(s, "hello")

    fmt.Println(str) // "hello"

    // We just want the zero value if it's nil
    fmt.Println(pointer.OrDefault(s)) // ""
}
```

`pointer.Or` will return the pointer it's passed in **only** if it's not nil! Otherwise you'll get back the value you provided.

If you just want the zero value instead, use `pointer.OrDefault`

### Credits

This package was created with [copier] and the [FollowTheProcess/go_copier] project template.

[copier]: https://copier.readthedocs.io/en/stable/
[FollowTheProcess/go_copier]: https://github.com/FollowTheProcess/go_copier
