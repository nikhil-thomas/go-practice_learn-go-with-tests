package main

import (
	"fmt"
	"io"
	"os"
)

// Greet writes greeting into buffer
func Greet(w io.Writer, v string) {
	fmt.Fprintf(w, "Hello, %s", v)

}

func main() {
	Greet(os.Stdout, "Chris")
}
