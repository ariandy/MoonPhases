package ChaprontTouze

import "math"

func T(day byte, month byte, year uint16) float64 {
	var T float64 = float64(K(day,month,year))/1236.85
	var rounding float64 = math.Round(T * 100000) / 100000
	return rounding
}