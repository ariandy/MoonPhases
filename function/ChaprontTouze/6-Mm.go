package ChaprontTouze

import (
	fm "moon_phase/function/FloatManipulation"
	am "moon_phase/function/AngleManipulation"
)

func Mm(day byte, month byte, year uint16) float64 {
	var k float64 = float64(K(day,month,year))
	var T float64 = T(day,month,year)

	var aleph, beth, gimel, daleth, he, result float64
	aleph = 201.5643
	beth = 385.81693528*k
	gimel = 0.0107582*T*T
	daleth = 0.00001238 *T*T*T
	he = 0.000000058 *T*T*T*T
	result = aleph+beth+gimel+daleth-he

	result = fm.FourDecimals(result) 
	result = am.DegreeCorrection(result)
	return fm.FourDecimals(result)
}