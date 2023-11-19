//This code uses an if statement to determine if a person is an adult or a minor based on their age.

package main

import "fmt"

func main() {
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a minor.")
    }
}
