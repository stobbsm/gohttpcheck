package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// Accepts 1 argument, the http address to check

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "expecting 1 argument only, in the form of http://<HOST>:<PORT>/{ENDPOINT}")
		os.Exit(1)
	}

	endpoint := os.Args[1]
	// Take no longer then 500ms to perform the check
	c := http.Client{
		Timeout: time.Millisecond * 500,
	}

	r, err := c.Head(endpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	r.Body.Close()
	fmt.Fprintf(os.Stderr, "status: %d", r.StatusCode)
}
