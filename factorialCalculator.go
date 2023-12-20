//This code calculates the factorial of a given number (in this case, 5) using a recursive function.

package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    num := 5
    result := factorial(num)
    fmt.Printf("Factorial of %d is %d\n", num, result)
}
