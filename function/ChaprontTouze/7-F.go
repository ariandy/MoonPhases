package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
	am "moon_phase/function/AngleManipulation"
)

func F(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph, beth, gimel, daleth, he, result float64
	aleph = 160.7108
	beth = 390.67050284*k
	gimel = 0.0016118*T*T
	daleth = 0.00000227*T*T*T 
	he = 0.000000011*T*T*T*T
	result = aleph + beth - gimel - daleth + he

	result = fm.FourDecimals(result) 
	result = am.DegreeCorrection(result)
	return fm.FourDecimals(result)
}