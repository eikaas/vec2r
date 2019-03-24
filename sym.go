package vec2r

import (
	"fmt"
)

const (
	// SymCross is the unicode symbol used for vector cross product
	SymCross = "\u2a2f"

	// SymDot is the unicode symbol for vector dot product
	SymDot = "\u22c5"
)

// SymArrow returns a string where the chosen rune, v, is written with an arrow above
func SymArrow(v rune) string {
	return fmt.Sprintf("%c\u20D7", v)
}

// SymMagnitude returns a string where the chosen run, v, is written with an arrow above it
// and enclosed in double vertical bars to symbolize its magnitude scalar identity
func SymMagnitude(v rune) string {
	return fmt.Sprintf("\u2016%c\u20d7\u2016", v)
}
