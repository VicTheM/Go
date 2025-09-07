package main

import (
	"fmt"
)

func main() {
	// start := time.Now()
	// fmt.Println("Main function started")
	// go routine()

	// time.Sleep(1 * time.Second)
	for x := 0; x < 1000000; x++ {
		go fmt.Println("rotine iteration: ", x, "")
	}

	// secs := time.Since(start).Seconds()
	// fmt.Println("secs used: ", secs)
}

// func routine() {
// 	for x := 0; x < 10; x++ {
// 		fmt.Println("rotine iteration: ", x)
// 	}
// }
