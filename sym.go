package vec2r

import (
	"fmt"
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

// SymDot returns a string containing the unicode symbol for vector dot product
func SymDot() string {
	return fmt.Sprintf("\u22c5")
}

// SymCross returns a string containing the unicode symbol for vector cross product
func SymCross() string {
	return fmt.Sprintf("\u2a2f")
}
