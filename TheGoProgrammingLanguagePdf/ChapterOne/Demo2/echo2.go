// Annother implementation of the echo program
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

/* range function returns two values:
1. The index of the element in the slice or array.
2. The value of the element at that index.
In this case, we are using the range function to iterate over the os.Args slice, starting
from index 1 (to skip the program name). The first value returned by range is ignored
(using the underscore _), and we only use the second value, which is the argument itself.
This is a more idiomatic way to write the loop in Go, as it avoids the need to manually
manage the index and the separator string.
*/
