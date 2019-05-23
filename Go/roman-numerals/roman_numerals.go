// Package romannumerals computes Roman Numerals given a "regular" (aka digital) number
package romannumerals

import (
	"errors"
	"strconv"
)

// Numerals is a struct that holds the relevant representations in an order of magnitude -
// the one, five, and ten characters for a single power of ten in Roman Numerals.
type Numerals struct {
	// Character equal to ONE unit
	one string
	// Character equal to FIVE units
	five string
	// Character equal to TEN units
	ten string
}

// Translator is a map of index type integer and value type Numerals. The index represents the
// order of magnitude to which the Numerals belong.
type Translator map[int]Numerals

const one = "I"
const five = "V"
const ten = "X"
const fifty = "L"
const hundred = "C"
const fivehundred = "D"
const thousand = "M"

// InitializeTranslator initializes and returns a translator object since I cannot have a constant map.
// In this case it is initialized to support numbers up to 3000 only.
func InitializeTranslator() (translator Translator) {
	translator = make(map[int]Numerals)
	translator[1] = InitializeNumerals(one, five, ten)
	translator[2] = InitializeNumerals(ten, fifty, hundred)
	translator[3] = InitializeNumerals(hundred, fivehundred, thousand)
	translator[4] = InitializeNumerals(thousand, "", "")

	return
}

// InitializeNumerals initializes and returns a Numerals struct for one power of ten
func InitializeNumerals(one string, five string, ten string) (numerals Numerals) {
	numerals = Numerals{
		one:  one,
		five: five,
		ten:  ten,
	}

	return
}

// ToRomanNumeral takes a digital number and translates it to its Roman Numeral equivalent.
func ToRomanNumeral(number int) (numeral string, err error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("can only translate numbers between 1 and 3000, inclusive")
	}

	var translator = InitializeTranslator()
	var slice = strconv.Itoa(number)

	for pos, char := range slice {
		numeral += DigitToRomanNumeral(char, len(slice)-pos, translator)
	}

	return
}

// DigitToRomanNumeral translates a single digit to its Roman Numeral equivalent.
func DigitToRomanNumeral(digit rune, orderOfMagnitude int, translator Translator) (numeral string) {
	var numerals = translator[orderOfMagnitude]

	switch digit {
	case '1':
		numeral = numerals.one
	case '2':
		numeral = numerals.one + numerals.one
	case '3':
		numeral = numerals.one + numerals.one + numerals.one
	case '4':
		numeral = numerals.one + numerals.five
	case '5':
		numeral = numerals.five
	case '6':
		numeral = numerals.five + numerals.one
	case '7':
		numeral = numerals.five + numerals.one + numerals.one
	case '8':
		numeral = numerals.five + numerals.one + numerals.one + numerals.one
	case '9':
		numeral = numerals.one + numerals.ten
	default:
		numeral = ""
	}

	return
}
