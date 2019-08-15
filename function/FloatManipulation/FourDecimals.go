package FloatManipulation

import (
	"math"
)

func FourDecimals(x float64) float64 {
	var temp float64 = x*10000
	temp = math.Round(temp)
	temp = temp / 10000
	return temp
}