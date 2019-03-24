package vec2r_test

import (
	"math"
	"testing"

	"github.com/matryer/is"

	"github.com/eikaas/vec2r"
)

func TestNew(t *testing.T) {
	is := is.New(t)

	a := vec2r.New(3.14, 1.161)
	is.True(a.X == 3.14)  // should equal 3.14
	is.True(a.Y == 1.161) // should equal 1.161

	b := vec2r.New(-1.0, -2.5)
	is.True(b.X == -1.0) // should equal -1.0
	is.True(b.Y == -2.5) // should equal -2.5
}

func TestMagnitude(t *testing.T) {
	is := is.New(t)
	a := vec2r.Vec2D{12.0, 24.0}
	b := vec2r.Vec2D{55.0, 120.0}

	is.Equal(math.Round(a.Magnitude()*100)/100, 26.83)  // should be equal
	is.Equal(math.Round(b.Magnitude()*100)/100, 132.00) // should be equal
}

func TestEquals(t *testing.T) {
	is := is.New(t)
	a := vec2r.Vec2D{3.141, 5.926}
	b := vec2r.Vec2D{3.141, 5.926}
	c := vec2r.Vec2D{3.141, -5.926}
	is.True(a.Equals(b))  // should be true
	is.True(b.Equals(a))  // should be true
	is.True(!b.Equals(c)) // should be false
	is.True(!a.Equals(c)) // should be false
	is.True(!c.Equals(a)) // should be false
	is.True(!c.Equals(b)) // should be false
}

func TestNormalize(t *testing.T) {
	is := is.New(t)
	v := vec2r.Vec2D{1.0, 2.0}
	u := vec2r.Vec2D{3.0, 4.0}

	nv := v.Normalize()
	nu := u.Normalize()

	is.Equal(math.Round(nv.X*100)/100, 0.45)
	is.Equal(math.Round(nv.Y*100)/100, 0.89)

	is.Equal(math.Round(nu.X*100)/100, 0.60)
	is.Equal(math.Round(nu.Y*100)/100, 0.80)
}

func TestAdd(t *testing.T) {
	is := is.New(t)

	c := vec2r.Vec2D{5.0, 10.0}
	c.Add(vec2r.Vec2D{1.0, 2.0})
	is.Equal(c.X, 6.0)  // should equal 6.0
	is.Equal(c.Y, 12.0) // should equal 12.0

	a := vec2r.Vec2D{15.0, 25.0}
	a.Add(vec2r.Vec2D{15.0, 25.0})
	is.Equal(a.X, 30.0) // should equal 30.0
	is.Equal(a.Y, 50.0) // should equal 50.0

	b := vec2r.Vec2D{-10.0, -1.0}
	b.Add(vec2r.Vec2D{10.0, 1.0})
	is.Equal(b.X, 0.0) // should equal 0.0
	is.Equal(b.Y, 0.0) // should equal 0.0
}

func TestSubtract(t *testing.T) {
	a := vec2r.Vec2D{15.0, 25.0}

	is := is.New(t)

	a.Subtract(vec2r.Vec2D{15.0, 25.0})
	is.Equal(a.X, 0.0) // should equal 0.0
	is.Equal(a.Y, 0.0) // should equal 0.0

}

func TestMultiply(t *testing.T) {
	is := is.New(t)

	a := vec2r.Vec2D{2.0, 4.0}
	a.Multiply(vec2r.Vec2D{2.0, 4.0})
	is.Equal(a.X, 4.0)  // should equal 4.0
	is.Equal(a.Y, 16.0) // should equal 16.0

	b := vec2r.Vec2D{-2.0, -4.0}
	b.Multiply(vec2r.Vec2D{2, 3})
	is.Equal(b.X, -4.0)  // should equal -4.0
	is.Equal(b.Y, -12.0) // should equal -12.0
}

func TestDivide(t *testing.T) {
	is := is.New(t)

	a := vec2r.Vec2D{10, 20}
	a.Divide(vec2r.Vec2D{2, 2})
	is.Equal(a.X, 5.0)  // should equal 5.0
	is.Equal(a.Y, 10.0) // should equal 10.0

	b := vec2r.Vec2D{-20, 20}
	b.Divide(vec2r.Vec2D{2, 2})
	is.Equal(b.X, -10.0) // should equal
	is.Equal(b.Y, 10.0)  // should equal
}

func TestMultiplyScalar(t *testing.T) {
	is := is.New(t)

	a := vec2r.Vec2D{10, 20}
	a.MultiplyScalar(1.5)
	is.Equal(a.X, 15.0) // should equal 15.0
	is.Equal(a.Y, 30.0) // should equal 30.0

	b := vec2r.Vec2D{10, 20}
	b.MultiplyScalar(-1.5)
	is.Equal(b.X, -15.0) // should equal -15.0
	is.Equal(b.Y, -30.0) // should equal -30.0
}

func TestDivideScalar(t *testing.T) {
	is := is.New(t)
	a := vec2r.Vec2D{1, 2}
	a.DivideScalar(2)
	is.Equal(a.X, 0.5) // should equal 0.5
	is.Equal(a.Y, 1.0) // should equal 1.0

	b := vec2r.Vec2D{60.0, 30.0}
	b.DivideScalar(3)
	is.Equal(b.X, 20.0) // should equal 20.0
	is.Equal(b.Y, 10.0) // should equal 10.0
}

func TestDotProduct(t *testing.T) {
	is := is.New(t)
	a := vec2r.Vec2D{1, 2}
	b := vec2r.Vec2D{4, 8}
	is.Equal(vec2r.DotProduct(a, b), 20.0) // should be 20

	c := vec2r.Vec2D{31, 42}
	d := vec2r.Vec2D{42, 56}
	is.Equal(vec2r.DotProduct(c, d), 3654.0) // should be 3654

	e := vec2r.Vec2D{-6.2, 6.0}
	f := vec2r.Vec2D{-2.3, -1.3}
	dotp := math.Round(vec2r.DotProduct(e, f)*100) / 100
	is.Equal(dotp, 6.46)
}

func TestAngle(t *testing.T) {
	is := is.New(t)
	u := vec2r.Vec2D{1, 2}
	v := vec2r.Vec2D{0, 3}
	ang := math.Round(vec2r.Angle(u, v)*100) / 100
	is.Equal(ang, 0.46)
}

func TestString(t *testing.T) {
	is := is.New(t)
	is.Equal(vec2r.New(3.1415, 14.15).String(), "(3.14, 14.15)")           // string should be '(3.14, 14.15)'
	is.Equal(vec2r.New(0, 0).String(), "(0, 0)")                           // string should be '(0.0, 0.0)'
	is.Equal(vec2r.New(999.9999, 999.9999).String(), "(1000.00, 1000.00)") // string should be '(1000.00, 1000.00)'
	is.Equal(vec2r.New(15.0, 7.0).String(), "(15, 7)")                     // string should be '(15, 7)'
}
