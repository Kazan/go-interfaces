package main

import (
	"os"

	"github.com/kazan/go-interfaces/example1/greeter"
	"github.com/kazan/go-interfaces/example1/writer"
)

func main() {
	// Create a new writer that writes to standard output
	w := writer.NewWriter(os.Stdout)

	// Create a new greeter with the name "World" and the writer
	greeter := greeter.NewGreeter("World", w)

	// Call the Greet method to write the greeting message
	greeter.Greet()
}
