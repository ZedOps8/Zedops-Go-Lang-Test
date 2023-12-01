//This code demonstrates the creation of 
//a simple struct to represent a person and prints out their name and age.

package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{Name: "Alice", Age: 30}
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
