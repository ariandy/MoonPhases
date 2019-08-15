package CalendarManipulation

import "math"

func ConvertDaysToYears(day byte, month byte, year uint16) float64 {
	var totalDays uint16 = 365
	if IsLeapYear(year) {
		totalDays = 366
	}
	var result float64 = float64(ConvertDateToDays(day,month,year))/float64(totalDays)
	var rounding float64 = math.Round(float64(int(result*1000))/10)/100+float64(year)
	return rounding
}