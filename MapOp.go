//This code demonstrates basic operations with maps, including access and adding new key-value pairs.


package main

import "fmt"

func main() {
    // Create a map of string to int
    ages := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }
    // Access and print values in the map
    fmt.Printf("Alice's age: %d\n", ages["Alice"])

    // Add a new entry to the map and print the updated map
    ages["Charlie"] = 35
    fmt.Printf("Updated map: %v\n", ages)
}

