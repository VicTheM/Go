// Fetches a list or urls sequentially
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		start2 := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		if b == nil {
			fmt.Printf("fetch: %s no data\n", url)
		}
		fmt.Printf("%.2fs elapsed\n", time.Since(start2).Seconds())
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
