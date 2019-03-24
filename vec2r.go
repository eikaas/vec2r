// Package vec2r is a simple 2D vector library
package vec2r

import (
	"fmt"
	"math"
)

// Vec2D is a 2-dimensional euclidean vector
type Vec2D struct {
	X float64
	Y float64
}

// New creates a new vector
func New(x, y float64) *Vec2D {
	return &Vec2D{x, y}
}

// Copy returns a copy of the vector
func (v *Vec2D) Copy() *Vec2D {
	return &Vec2D{v.X, v.Y}
}

// Normalize returns a new, normalized version of the vector in the same direction with length 1
func (v *Vec2D) Normalize() *Vec2D {
	u := v.Copy()
	u.DivideScalar(v.Magnitude())
	return u
}

// Magnitude returns the magnitude of the vector
func (v *Vec2D) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Length returns the length of the vector
func (v *Vec2D) Length() float64 {
	return v.Magnitude()
}

// Equals returns true if the vector equals u
func (v *Vec2D) Equals(u Vec2D) bool {
	if v.X == u.X && v.Y == u.Y {
		return true
	}
	return false
}

// Add adds u to the vector
func (v *Vec2D) Add(u Vec2D) {
	v.X = v.X + u.X
	v.Y = v.Y + u.Y
}

// Subtract substracts u from the vector
func (v *Vec2D) Subtract(u Vec2D) {
	v.X -= u.X
	v.Y -= u.Y
}

// Multiply multiplies the vector with vector u
func (v *Vec2D) Multiply(u Vec2D) {
	v.X *= u.X
	v.Y *= u.Y
}

// Divide divides vector v by vector u
func (v *Vec2D) Divide(u Vec2D) {
	v.X /= u.X
	v.Y /= u.Y
}

// MultiplyScalar multiplies the vector by the scalar s
func (v *Vec2D) MultiplyScalar(s float64) {
	v.X *= s
	v.Y *= s
}

// DivideScalar divides the vector by the scalar s
func (v *Vec2D) DivideScalar(s float64) {
	v.X /= s
	v.Y /= s
}

// DotProduct calculates the dot-product of two vectors
func DotProduct(v, u Vec2D) float64 {
	return (u.X * v.X) + (u.Y * v.Y)
}

// CrossProduct calculates the (quasi) cross-product of two vectors
func CrossProduct(v, u Vec2D) float64 {
	return (u.X * v.X) - (u.Y * v.Y)
}

// Angle calculates angle between two vectors. Returned result is in radians
func Angle(v, u Vec2D) float64 {
	return math.Acos(DotProduct(u, v) / (u.Magnitude() * v.Magnitude()))
}

// String returns the string representation of the vector, rounded to two decimal places.
// If neither X or Y have a fraqtional part, they are cast to and printed as integers.
func (v *Vec2D) String() string {
	_, xfraq := math.Modf(v.X)
	_, yfraq := math.Modf(v.Y)
	if xfraq != 0.0 && yfraq != 0.0 {
		return fmt.Sprintf("(%.2f, %.2f)", v.X, v.Y)
	}
	return fmt.Sprintf("(%d, %d)", int(v.X), int(v.Y))
}
