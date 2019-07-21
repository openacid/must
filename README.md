# must

`must` is a "design by contract" implementation in golang,
for addressing silent bug that is caused by unexpected input etc.

[![Travis-CI](https://api.travis-ci.org/openacid/must.svg?branch=master)](https://travis-ci.org/openacid/must)
[![GoDoc](https://godoc.org/github.com/openacid/must?status.svg)](http://godoc.org/github.com/openacid/must)
[![Report card](https://goreportcard.com/badge/github.com/openacid/must)](https://goreportcard.com/report/github.com/openacid/must)
[![GolangCI](https://golangci.com/badges/github.com/openacid/must.svg)](https://golangci.com/r/github.com/openacid/must)
[![Sourcegraph](https://sourcegraph.com/github.com/openacid/must/-/badge.svg)](https://sourcegraph.com/github.com/openacid/must?badge)
![stability-unstable](https://img.shields.io/badge/stability-unstable-yellow.svg)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
 

- [Usage](#usage)
- [API](#api)
- [Examples](#examples)
- [Install](#install)
- [Customize tags](#customize-tags)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Usage

To enable expectation check in test environment, use `go build|test -tags debug`.
To disable it for a release, just `go build`.

```go
package main

import (
	"fmt"
	"math/bits"

	"github.com/openacid/must"
)

var mustbe = must.Be

func rshift(a, b int) int {

	// "go build" emits a single No-op instruction.
	// "go build -tags debug" will call the function and to the checking.
	mustbe.OK(func() {
		mustbe.NotZero(b)
		mustbe.True(bits.TrailingZeros(uint(a)) > 2, "a must be multiple of 8")
	})

	return a >> uint(b)
}

func main() {
	fmt.Println(rshift(0xf, 1)) // panic at line 19 with "go run -tags debug"
}
```

With the above code:

**Enable check** with `go run -tags debug`.

It would panic because `a` does not satisfy the input
expectation:

```
panic:
        ...
        Error:          Should be true
        Messages:       a must be multiple of 8
```

To disassemble the binary we could see there is checking statement instruction:

```
> go build -tags debug -o bin-debug .
> go tool objdump -S bin-debug
>
TEXT main.rshift(SB) github.com/openacid/must/examples/rshift/composite/main.go
func rshift(a, b int) int {
  ...
        mustbe.OK(func() {
  0x12481fd             0f57c0                  XORPS X0, X0
  0x1248200             0f110424                MOVUPS X0, 0(SP)
  ...
        f()
  0x124822d             488b1c24                MOVQ 0(SP), BX
  ...
```

**Disable check** with `go run .`

It just silently ignores the expectation and print the result:

```
7
```

To disassemble the binary we could see there is only a `NOPL` instruction:

```
> go build -o bin-release .
> go tool objdump -S bin-release

TEXT main.rshift(SB) github.com/openacid/must/examples/rshift/composite/main.go
        mustbe.OK(func() {
  0x1246030             90                      NOPL
  0x1246031             488b4c2410              MOVQ 0x10(SP), CX
        return a >> uint(b)
  0x1246036             4883f940                CMPQ $0x40, CX
  ...
  0x1246050             c3                      RET
```


# API

`must` uses the popular [testify](github.com/stretchr/testify) as underlying
asserting engine, thus there is a corresponding function defined for every
`testify` assertion function.
E.g.:

```
must.Be.True(v bool) --> testify/assert.True(t TestingT, v bool)
```


# Examples

[rshift-simple-check](examples/rshift/simple)

[rshift-composite-check](examples/rshift/composite)


# Install

```
go get -u github.com/openacid/must
```


# Customize tags

`must` provides with a default tag "debug" to enable expectation check.
It is also very easy to define your own tags.
To do this, create two files in one of your package, such as `mymust_debug.go` and `mymust_release.go`, like following:

`mymust_debug.go`:

```go
// +build mydebug

package your_package
import "github.com/openacid/must/enabled"
var mymustBe = enabled.Be
```

`mymust_release.go`:

```go
// +build !mydebug

package your_package
import "github.com/openacid/must/disabled"
var mymustBe = disabled.Be
```

And replace `must.Be` with `mymustBe` in your source codes.
Then your could enable it just by:

```
go build -tags mydebug
```

See more: [go-build-constraints](https://golang.org/pkg/go/build/#hdr-Build_Constraints)