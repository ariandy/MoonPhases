package NaiveTesting

import (
	"fmt"
	cm "moon_phase/function/CalendarManipulation"
)
func CalendarManipulationTest(day byte, month byte, year uint16) {
	fmt.Println()
	fmt.Println("---Calendar Manipulation Testing Purpose---")
	fmt.Println()
	fmt.Println("IsLeapYear : ", cm.IsLeapYear(year))
	fmt.Println("ArrayDaysOfThemonth : ", cm.ArrayDaysOfTheMonth(year))
	fmt.Println("ConvertPreviousMonthToDays : ",cm.ConvertPreviousMonthToDays(month,year))
	fmt.Println("ConvertDateToDays : ", cm.ConvertDateToDays(day, month, year))
	fmt.Println("ConvertDateToYears : ", cm.ConvertDaysToYears(day, month, year))
	fmt.Println()
}