package algorythms

import (
	"fmt"
	"strings"

	"github.com/Dario0117/simple-algorythms-in-go/localutils"
)

var equivalences = map[string]uint16{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var validRoman = "IVXLCDM"

// RunRomanNumeralParser ...
func RunRomanNumeralParser(args []string) {
	var result int
	var romanNumeral string
	var err bool
	if len(args) > 0 {
		result, romanNumeral, err = parseRoman(strings.ToUpper(args[0]))
	} else {
		romanNumeral = localutils.Prompt("Insert the roman numeral ")
		result, romanNumeral, err = parseRoman(strings.ToUpper(romanNumeral))
	}
	if !err {
		fmt.Printf("🎊 The convertion from %s to dec is: %d 🎊\n", romanNumeral, result)
	} else {
		fmt.Printf("❌ We can't convert %s to decimal\n", romanNumeral)
	}
}

func parseRoman(romanNumeral string) (int, string, bool) {
	var total int
	fives := "VLD"
	validForI := "VX"
	validForX := "LC"
	for idx, char := range romanNumeral {
		if !strings.Contains(validRoman, string(char)) {
			return 0, romanNumeral, true
		}
		actualValue := int(equivalences[string(char)])
		if idx > 0 {
			prevValue := int(equivalences[string(romanNumeral[idx-1])])
			if strings.Contains(fives, string(romanNumeral[idx-1])) {
				if actualValue > prevValue {
					return 0, romanNumeral, true
				}
			}
			if !strings.Contains(validForI, string(char)) {
				if prevValue == int(equivalences["I"]) {
					return 0, romanNumeral, true
				}
			}
			if !strings.Contains(validForX, string(char)) {
				if prevValue == int(equivalences["X"]) {
					return 0, romanNumeral, true
				}
			}
			if actualValue > prevValue {
				actualValue = actualValue - (prevValue * 2)
			}
			if idx > 1 {
				scndPrevValue := int(equivalences[string(romanNumeral[idx-2])])
				if prevValue == scndPrevValue {
					if prevValue != actualValue {
						return 0, romanNumeral, true
					}
				}
				if idx > 2 {
					thrdPrevValue := int(equivalences[string(romanNumeral[idx-3])])
					if scndPrevValue == thrdPrevValue || actualValue < scndPrevValue {
						return 0, romanNumeral, true
					}
				}
			}
		}
		total += actualValue
	}
	return total, romanNumeral, false
}
