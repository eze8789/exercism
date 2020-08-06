package leap

// IsLeapYear calculate if a year is leap. Video of what's a leap year http://www.youtube.com/watch?v=xX96xng7sAE.
func IsLeapYear(y int) bool {
	if y%4 == 0 {
		if y%100 == 0 {
			if y%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
