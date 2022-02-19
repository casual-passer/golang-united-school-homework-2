package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

type Sides int

const SidesTriangle Sides = 3
const SidesSquare Sides = 4
const SidesCircle Sides = 0

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	case SidesTriangle:
		return math.Sqrt(3) * sideLen * sideLen / 4
	default:
		return -1
	}
}
