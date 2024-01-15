package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	packagedRune := []rune(packedString)
	var newString string

	for index, value := range packagedRune {
		if unicode.IsDigit(value) && index == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(value) && unicode.IsDigit(packagedRune[index-1]) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(value) {
			digit, _ := strconv.Atoi(string(value))

			if digit == 0 {
				newString = newString[:len(newString)-1]
				continue
			}

			digit--
			newString += strings.Repeat(string(packagedRune[index-1]), digit)
			continue
		}

		newString += string(value)
	}

	fmt.Println(newString)

	return newString, nil
}
