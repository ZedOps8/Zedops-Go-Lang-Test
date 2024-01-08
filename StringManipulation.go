//This code demonstrates basic string manipulation by converting a text to uppercase and lowercase.

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "This is a sample text"
	uppercaseText := strings.ToUpper(text)
	lowercaseText := strings.ToLower(text)

	fmt.Println("Original Text:", text)
	fmt.Println("Uppercase Text:", uppercaseText)
	fmt.Println("Lowercase Text:", lowercaseText)
}
