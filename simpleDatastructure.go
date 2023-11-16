
package main

import "fmt"

func main() {
    // Creating and using an array
    numbers := [3]int{1, 2, 3}
    fmt.Println("Array:", numbers)

    // Creating and using a slice
    colors := []string{"red", "green", "blue"}
    fmt.Println("Slice:", colors)

    // Creating and using a map
    person := map[string]string{"name": "Alice", "age": "30"}
    fmt.Println("Map:", person)
}
