// Package that determines if a year is a leap year
package leap

// Given a year, returns true if is a leap year and false if it is not.
func IsLeapYear(year int) bool {

	var isDivisibleByFour bool = year % 4 == 0
	var isDivisibleByOneHundred bool = year % 100 == 0
	var isDivisibleByFourHundred bool = year % 400 == 0

	var isLeapYear bool = isDivisibleByFour && (!isDivisibleByOneHundred || isDivisibleByFourHundred)
	
	return isLeapYear
}
