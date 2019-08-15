package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
)

func Omega(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph, beth, gimel, daleth, result float64
	aleph = 124.7746
	beth = 1.56375588*k
	gimel = 0.0020672*T*T 
	daleth = 0.00000215*T*T*T
	result = aleph - beth + gimel + daleth

	return fm.FourDecimals(result)
}