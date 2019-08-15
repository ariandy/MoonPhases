package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
)

func Mm(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph float64 = 201.5643
	var beth float64 = 385.81693528*k
	var gimel float64 = 0.0107582*T*T
	var daleth float64 = 0.00001238 *T*T*T
	var he float64 = 0.000000058 *T*T*T*T
	var result float64 = aleph+beth+gimel+daleth-he

	return fm.FiveDecimal(result)
}