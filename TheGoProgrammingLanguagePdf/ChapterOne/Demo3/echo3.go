// Efficient implementation of the echo program
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// strings.join takes a slice of strings and a separator string, and returns a single string
// with the elements of the slice joined together, separated by the separator string.
// In this case, we are joining the command-line arguments (excluding the program name) with
// a space as the separator. This is a more efficient and idiomatic way to concatenate strings
// in Go, as it avoids the need for manual string concatenation and the use of a
