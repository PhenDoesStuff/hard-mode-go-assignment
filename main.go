package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1]) // os.Open returns a *File (f) and an Error, if there is one, (err)
	if err != nil {
		fmt.Println("Error:", err) // Error handling
		os.Exit(1)
	}
	io.Copy(os.Stdout, f) // File uses Reader and Writer, so we can use io.Copy
}
