package romannumerals

import "strconv"

type Numerals struct {
	// Character equal to ONE unit
	one string
	// Character equal to FIVE units
	five string
	// Character equal to TEN units
	ten string
}

type Translator map[int]Numerals

const one = "I"
const five = "V"
const ten = "X"
const fifty = "L"
const hundred = "C"
const fivehundred = "D"
const thousand = "M"

func InitializeTranslator() (translator Translator) {
	translator[1] = InitializeNumerals(one, five, ten)
	translator[2] = InitializeNumerals(ten, fifty, hundred)
	translator[3] = InitializeNumerals(hundred, fivehundred, thousand)
	translator[4] = InitializeNumerals(thousand, "", "")

	return
}

func InitializeNumerals(one string, five string, ten string) (numerals Numerals) {
	numerals = Numerals{
		one:  one,
		five: five,
		ten:  ten,
	}

	return
}

func ToRomanNumeral(number int) (numeral string) {
	var translator = InitializeTranslator()
	var slice = strconv.Itoa(number)

	numeral += DigitToRomanNumeral(slice[0], len(slice)-0, translator)
	numeral += DigitToRomanNumeral(slice[1], len(slice)-1, translator)
	numeral += DigitToRomanNumeral(slice[2], len(slice)-2, translator)
	numeral += DigitToRomanNumeral(slice[3], len(slice)-3, translator)
	return
}

func DigitToRomanNumeral(digit byte, orderOfMagnitude int, translator Translator) (numeral string) {

	return
}
