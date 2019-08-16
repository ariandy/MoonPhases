package NaiveTesting

import (
	"fmt"
	ct "moon_phase/function/ChaprontTouze"
)

func ChaprontTouzeTest(day byte, month byte, year uint16) {
	fmt.Println()
	fmt.Println("---Chapront-Touze Function Testing Purpose---")
	fmt.Println()
	fmt.Println("k : ",ct.K(day,month,year))
	fmt.Println("T : ",ct.T(day,month,year))
	fmt.Println("JDE : ",ct.JDE(day,month,year))
	fmt.Println("E : ",ct.E(day,month,year))
	fmt.Println("M : ",ct.M(day,month,year))
	fmt.Println("Mm : ",ct.Mm(day,month,year))
	fmt.Println("F : ",ct.F(day,month,year))
	fmt.Println("Omega : ",ct.Omega(day,month,year))
}