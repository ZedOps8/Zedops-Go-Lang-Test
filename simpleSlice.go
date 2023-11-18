//This code demonstrates the creation and iteration over a slice of integers.

package main

import "fmt"

func main() {
    // Create a slice of integers
    numbers := []int{1, 2, 3, 4, 5}

    // Iterate over the slice and print each element
    for _, num := range numbers {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
}
