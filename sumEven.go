//This code calculates and prints the sum of even numbers between 1 and 10.

package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            sum += i
        }
    }
    fmt.Printf("Sum of even numbers from 1 to 10: %d\n", sum)
}
