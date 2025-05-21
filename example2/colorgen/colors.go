package colorgen

import "fmt"

// ColorType is an enum for different color generation strategies.
type ColorType int

const (
	TotallyRandom ColorType = iota
	RedShades
	GreyTones
)

// Colorizer is an interface that defines a method for generating random colors.
type Colorizer interface {
	Color() string
	Name() string
}

// New returns a Colorizer based on the given ColorType.
func New(t ColorType) (Colorizer, error) {
	switch t {
	case RedShades:
		return nil, nil
	case TotallyRandom:
		return &RandomColorizer{}, nil
	default:
		return nil, fmt.Errorf("unimplemented color type: %d", t)
	}
}
