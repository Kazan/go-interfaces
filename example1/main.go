package main

import (
	"os"

	"github.com/kazan/go-interfaces/example1/colorgen"
	"github.com/kazan/go-interfaces/example1/greeter"
	"github.com/kazan/go-interfaces/example1/writer"
)

func main() {
	// Create a new writer that writes to standard output
	w := writer.NewWriter(os.Stdout)

	// Create a new RandomColorizer that generates random colors
	randomColor, err := colorgen.New(colorgen.TotallyRandom)
	if err != nil {
		panic(err)
	}

	// Create a new greeter with the name "World" and the writer
	// and it will use the RandomColorizer to generate a random color
	greeterRandom := greeter.NewGreeter("Random", w, randomColor)

	// Call the Greet methods to write the greeting messages
	greeterRandom.Greet()
}
