// Package to describe and compute triangles
package triangle

import "math"

// Create a Kind type to describe what kind of triangle we have
type Kind int

const (
    NaT Kind = iota // 0
    Equ             // 1
    Iso             // 2
    Sca             // 3
)

// Calculate the Kind of triangle based on its side lengths
func KindFromSides(a, b, c float64) Kind {
	
	if (!IsATriangle(a, b, c)) {
		return NaT
	} else if (IsScalene(a, b, c)) {
		return Sca
	} else if (IsEquilateral(a, b, c)) {
		return Equ
	}
	
	return Iso

}

// Determines if a shape is a triangle at all given its side lengths
func IsATriangle(a, b, c float64) bool {

	return 	SidesAreNumbers(a, b, c) && 
			SidesHaveLength(a, b, c) && 
			!OneOrMoreSidesAreInfinite(a, b, c) && 
			AnyTwoSidesSumToGreaterThanThird(a, b, c)
}

// Returns true if all sides are numbers
func SidesAreNumbers(a, b, c float64) bool {
	return !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c)
}

// Returns true if all sides have a length greater than zero
func SidesHaveLength(a, b, c float64) bool {
	return a > 0 && b > 0 && c > 0 
}

// Returns true if any side is (positive) infinite length
func OneOrMoreSidesAreInfinite(a, b, c float64) bool {
	return math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1)
}

// Returns true if any two sides sum to greater than the length of the third side
func AnyTwoSidesSumToGreaterThanThird(a, b, c float64) bool {
	return a + b >= c && b + c >= a && c + a >= b 
}

// Determines if a triangle is equilateral. Note that shape must have already
// been determined to be a triangle.
func IsEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

// Determines if a triangle is isoceles. Note that shape must have already been 
// determined to be a triangle and NOT equilateral.
func IsIsoceles(a, b, c float64) bool {
	return a == b || b == c || c == a
}

// Determines if a triangle is scalene. Note that shape must have already
// been determined to be a triangle.
func IsScalene(a, b, c float64) bool {
	return a != b && b != c && c != a
}




