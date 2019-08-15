package ChaprontTouze

import "math"

func JDE(day byte, month byte, year uint16) float64 {
	var aleph, beth, gimel, daleth, he float64
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)
	var result float64
	aleph = 2451550.09766
	beth = 29.530588861*k 
	gimel = 0.00015437 * T*T
	daleth = 0.000000150 * T*T*T
	he = 0.00000000073 * T*T*T*T
	result = aleph + beth + gimel - daleth + he

	var convertedToUint64 uint64 = uint64(result*1000000)
	var convertedToFloat64 float64 = float64(convertedToUint64/10)
	var rounding float64 = math.Round(convertedToFloat64)
	var final float64 = rounding/100000
	return final
}