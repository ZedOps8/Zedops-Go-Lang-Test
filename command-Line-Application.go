package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Usage: go-example <arg1> <arg2>")
        return
    }

    fmt.Println("Arguments:")
    for i, arg := range args {
        fmt.Printf("Arg %d: %s\n", i, arg)
    }
}
