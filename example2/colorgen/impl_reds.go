package colorgen

import (
	"fmt"
	"math/rand"
)

// RedShadesColorizer is a struct that implements the Colorizer interface.
type RedShadesColorizer struct{}

// Color generates a random shade of red in hexadecimal format.
// It returns a string representing the color.
// It implements the Color method of the Colorizer interface.
func (c *RedShadesColorizer) Color() string {
	r := rand.Intn(256)
	g := 0
	b := 0
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// Name returns the name of the colorizer.
// It implements the Name method of the Colorizer interface.
func (c *RedShadesColorizer) Name() string {
	return "RedShadesColorizer"
}
