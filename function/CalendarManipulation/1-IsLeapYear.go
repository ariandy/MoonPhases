package CalendarManipulation

func IsLeapYear(year uint16) bool {
	return (((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)) 
}