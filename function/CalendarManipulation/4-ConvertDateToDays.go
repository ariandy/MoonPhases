package CalendarManipulation

func ConvertDateToDays(day byte, month byte, year uint16) uint16 {
	return uint16(day) + ConvertPreviousMonthToDays(month,year)
}