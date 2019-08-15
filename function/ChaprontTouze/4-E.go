package ChaprontTouze

import "math"

func E(day byte, month byte, year uint16) float64 {
	var T float64 = T(day,month,year)
	var result float64 = 1 - 0.002516*T - 0.0000074*T*T

	var temp uint32 = uint32(result*100000000)
	var temp2 float64 = float64(temp)/10
	temp2 = math.Round(temp2)
	temp2 = temp2 / 10000000
	// var temp4 uint32 = uint32(math.Round(temp2)) 
	return temp2
}