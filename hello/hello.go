// Source: https://golang.org/doc/code.html
package main

import (
    "fmt"
    "go/stringutil"
)

func main() {
    var msg string = "Hello, world"
    fmt.Printf(msg + "\n")
    fmt.Printf(stringutil.Reverse(msg) + "\n")
}

