package ChaprontTouze

import (
	"math"
	cm "moon_phase/function/CalendarManipulation"
)

func K(day byte, month byte, year uint16) int16 {
	var k float64 = cm.ConvertDaysToYears(day, month, year)
	k = (k-2000)*12.3685
	var temp int16 = int16(math.Round(k))
	return temp
}