package FloatManipulation

import (
	"math"
)

func FiveDecimal(x float64) float64 {
	var temp float64 = x*100000
	temp = math.Round(temp)
	temp = temp / 100000
	return temp
}