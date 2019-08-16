package main

import (
	// "fmt"
	// "math"
	// "math/big"
	// cm "moon_phase/function/CalendarManipulation"
	// ct "moon_phase/function/ChaprontTouze"
	nt "moon_phase/function/NaiveTesting"
)

func main()  {
	var day byte = 15
	var month byte = 2
	var year uint16 = 1977
	nt.CalendarManipulationTest(day, month, year)
	nt.ChaprontTouzeTest(day, month, year)
	// fmt.Println("k : ",ct.K(day,month,year))
	// fmt.Println("T : ",ct.T(day,month,year))
	// fmt.Println("JDE : ",ct.JDE(day,month,year))
	// fmt.Println("E : ",ct.E(day,month,year))
	// fmt.Println("M : ",ct.M(day,month,year))
	// fmt.Println("Mm : ",ct.Mm(day,month,year))
	// fmt.Println("F : ",ct.F(day,month,year))
	// fmt.Println("Omega : ",ct.Omega(day,month,year))
}

var day byte = 15
var month byte = 2
var year uint16 = 1977

