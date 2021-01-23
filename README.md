# script

A super package to provide syntactic sugars for writing scripts in Golang.

## Use with Gorun

  1. Go get to obtain the `gorun` binary: `go get gorun`
  1. Create an empty dir and initialize a new go module, `mkdir build`, `cd build`, `go mod init build`
  1. Inside the `build` dir, call `go get -u github.com/chanwit/script@master` to add it as a dependency
  1. Create `build.go`, as followed:
```go
/// 2>/dev/null ; gorun "$0" "$@" ; exit $?

// go.mod >>>
// module build
// go 1.15
// require (
//   github.com/chanwit/script v0.0.0-20210123134408-7360ef9587f5
// )
// <<< go.mod

package main

import (
	. "github.com/chanwit/script"
)

func main() {
	out := Var()
	System("echo hello").To(out)
	Printf("we are echoing: %s\n", out.String())

	list := Var()
	Exec("docker", "ps").To(list)
	Print(list.String())
}
```
  then you can `chmod +x build.go` and simply run it with `./build.go`

## Features

- Go Pipe wrapper to provide fluent interface
- Variable declaration for storing results from sub-processes
