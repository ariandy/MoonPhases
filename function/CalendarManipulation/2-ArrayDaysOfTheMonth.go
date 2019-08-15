package CalendarManipulation

func ArrayDaysOfTheMonth(year uint16) [12]byte {
	var daysInMonth = [12]byte{31,28,31,30,31,30,31,31,30,31,30,31}
	if IsLeapYear(year) {
		daysInMonth[1] = 29
	}
	return daysInMonth
}