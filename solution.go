package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type myInt int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesTriangle myInt = 3
	SidesSquare   myInt = 4
	SidesCircle   myInt = 0
)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	var square float64
	switch sidesNum {
	case SidesCircle:
		square = math.Pi * math.Pow(sideLen, 2)
	case SidesTriangle:
		square = math.Sqrt(3) * math.Pow(sideLen, 2) / 4
	case SidesSquare:
		square = math.Pow(sideLen, 2)
	}
	return square
}
