//This code demonstrates a switch statement to determine the type of day based on a given string

package main

import "fmt"

func main() {
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("It's the start of the week.")
    case "Friday":
        fmt.Println("It's the end of the week.")
    default:
        fmt.Println("It's a regular day.")
    }
}
