package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
)

func Omega(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph float64 = 124.7746
	var beth float64 = 1.56375588*k
	var gimel float64 = 0.0020672*T*T 
	var daleth float64 = 0.00000215*T*T*T
	var result float64 = aleph - beth + gimel + daleth

	return fm.FiveDecimal(result)
}