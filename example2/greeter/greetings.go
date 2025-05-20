package greeter

import (
	"fmt"

	"github.com/kazan/go-interfaces/example2/colorgen"
	"github.com/kazan/go-interfaces/example2/writer"
)

// Greeter is a struct that holds a name and a writer.
type Greeter struct {
	name   string
	writer writer.Writer
	color  colorgen.Colorizer
}

// NewGreeter creates a new Greeter with the specified name and writer.
func NewGreeter(name string, w writer.Writer, c colorgen.Colorizer) *Greeter {
	return &Greeter{
		name:   name,
		writer: w,
		color:  c,
	}
}

// Greet writes a greeting message to the writer.
func (g *Greeter) Greet() {
	g.writer.Write(fmt.Sprintf("Hello, %s! Your color is %s\n", g.name, g.color.Color()))
}
