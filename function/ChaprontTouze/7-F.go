package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
)

func F(day byte, month byte, year uint16) float64  {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph float64 = 160.7108
	var beth float64 = 390.67050284*k
	var gimel float64 = 0.0016118*T*T
	var daleth float64 = 0.00000227*T*T*T 
	var he float64 = 0.000000011*T*T*T*T
	var result float64 = aleph + beth - gimel - daleth + he

	return fm.FiveDecimal(result)
}