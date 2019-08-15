package ChaprontTouze

import (
	// "math"
	fm "moon_phase/function/FloatManipulation"
)

func M(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph float64 = 2.5534
	var beth float64 = 29.10535670*k
	var gimel float64 = 0.0000014*T*T
	var daleth float64 = 0.00000011*T*T*T
	var result float64 = aleph+beth-gimel-daleth

	return fm.FiveDecimal(result)
}