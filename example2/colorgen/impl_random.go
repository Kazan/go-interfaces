package colorgen

import (
	"fmt"
	"math/rand"
)

// RandomColorizer is a struct that implements the Colorizer interface.
type RandomColorizer struct{}

// Color generates a random color in hexadecimal format.
// It returns a string representing the color.
// It implements the Color method of the Colorizer interface.
func (c *RandomColorizer) Color() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Name returns the name of the colorizer.
// It implements the Name method of the Colorizer interface.
func (c *RandomColorizer) Name() string {
	return "RandomColorizer"
}
