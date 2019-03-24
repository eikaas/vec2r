# vec2r
Simple 2D vector library for silly games, ~~self-driving trucks~~, ~~calculating rocket trajectories~~, etc.

## Usage
```go
// Create vectors
v := vec2r.New(1.2, 3.4)
u := vec2r.New(5.6, 7.8)

// Vector atithmetic
v.Add(u)
v.Subtract(u)
v.Multiply(u)
v.Divide(u)

// Scalar operations
v.MultiplyScalar(3.1415)
v.DivideScalar(4.2)

// Magnitude
v.Magnitude()

// u is a normalized vector with the same direction as vector v
u := v.Normalize()

// Dot Product
s := vec2r.DotProduct(v, u)

// "Cross" Product
s := vec2r.CrossProduct(v, u)

// Get angle between vectors v & u
degrees := vec2r.Angle(v, u) * math.Pi/180

// Arrow notation
fmt.Println(vec2r.SymArrow('u'))

// Arrow notation, enclodes in vertical magnitude bars
fmt.Println(vec2r.SymMagnitude('u'))

// Cross and Dot operators
fmt.Println(vec2r.SymDot())
fmt.Println(vec2r.SymCross())
```
