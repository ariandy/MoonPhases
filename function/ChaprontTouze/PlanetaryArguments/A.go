package PlanetaryArguments

import (
	ct "moon_phase/function/ChaprontTouze"
)

func A(day byte, month byte, year uint16) [14]float64 {
	var k float64 = float64(ct.K(day,month,year))
	var T float64 = ct.T(day,month,year)
	var PlanetaryArgs = [14]float64{
		299.77 + 0.107408*k - 0.009173*T*T, //A1
		251.88 + 0.016321*k, //A2
		251.83 + 26.651886*k, //A3
		349.42 + 36.412478*k, //A4
		84.66 + 18.206239*k, //A5
		141.74 + 53.303771*k, //A6
		207.14 + 2.453732*k, //A7
		154.84 + 7.306860*k, //A8
		34.52 + 27.261239*k, //A9
		207.19 + 0.121824*k, //A10
		291.34 + 1.844379*k, //A11
		161.72 + 24.198154*k, //A12
		239.56 + 25.513099*k,
		331.55 + 3.592518*k }
	return PlanetaryArgs
}