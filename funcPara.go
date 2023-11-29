//This code defines a function to add 
//two numbers and then calls the function with values to calculate the result.

package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(10, 5)
    fmt.Printf("Result: %d\n", result)
}
