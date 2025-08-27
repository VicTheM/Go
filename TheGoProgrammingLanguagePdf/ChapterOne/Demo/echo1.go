// An inefficient inplementation of bash echo command
package main

import (
	"fmt"
	"os"
)

func main() {
	var strr, seperator string // Declaration (i below was initialized on declaration)
	for i := 1; i < len(os.Args); i++ {
		strr += seperator + os.Args[i]
		seperator = " "
	}
	fmt.Println(strr)
}
