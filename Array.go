//This code demonstrates basic operations with arrays, including element access and modification.

package main

import "fmt"

func main() {
    // Create an array of integers
    numbers := [5]int{1, 2, 3, 4, 5}

    // Access and print elements of the array
    fmt.Printf("Element at index 2: %d\n", numbers[2])

    // Modify an element and print the updated array
    numbers[3] = 10
    fmt.Printf("Updated array: %v\n", numbers)
}
