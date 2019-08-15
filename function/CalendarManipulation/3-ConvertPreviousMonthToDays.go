package CalendarManipulation

func ConvertPreviousMonthToDays(month byte, year uint16) uint16 {
	var init uint16 = 0
	var i byte
	for i = 0; i < month-1; i++ {
		init = init + uint16(ArrayDaysOfTheMonth(year)[i])
	}
	return init
}