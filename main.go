package main

import (
	"fmt"
	// "math"
	// "math/big"
	// cm "moon_phase/function/CalendarManipulation"
	ct "moon_phase/function/ChaprontTouze"
	// nt "moon_phase/function/NaiveTesting"
)

func main()  {
	var day byte = 15
	var month byte = 2
	var year uint16 = 1977
	// nt.CalendarManipulationTest(day, month, year)
	fmt.Println(ct.K(day,month,year))
	fmt.Println(ct.T(day,month,year))
	fmt.Println(ct.JDE(day,month,year))
	fmt.Println(ct.E(day,month,year))
	fmt.Println(ct.M(day,month,year))
	fmt.Println(ct.Mm(day,month,year))
	fmt.Println(ct.F(day,month,year))
}

var day byte = 15
var month byte = 2
var year uint16 = 1977

