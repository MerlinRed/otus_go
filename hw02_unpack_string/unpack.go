package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	if len(packedString) == 0 {
		return packedString, nil
	}

	if len(packedString) < 2 {
		return packedString, nil
	}

	firstAlpha := packedString[0]
	lastAlpha := packedString[len(packedString)-1]

	if _, err := strconv.Atoi(string(firstAlpha)); err == nil {
		return "", ErrInvalidString

	}

	var builder strings.Builder
	var alpha string

	for index, value := range strings.Split(packedString, "") {
		if _, err := strconv.Atoi(value); err == nil {
			if string(lastAlpha) == value {
				if _, err := strconv.Atoi(string(packedString[index-1])); err == nil {
					return "", ErrInvalidString
				}
			}

			if string(firstAlpha) == value {
				if _, err := strconv.Atoi(string(packedString[index+1])); err == nil {
					return "", ErrInvalidString
				}
			}

			if _, err := strconv.Atoi(string(packedString[index+1])); err == nil {
				return "", ErrInvalidString
			}

			if _, err := strconv.Atoi(string(packedString[index-1])); err == nil {
				return "", ErrInvalidString
			}

		}

		if digit, err := strconv.Atoi(value); err == nil && digit != 0 {
			digit -= 1
			builder.WriteString(strings.Repeat(alpha, digit))
		} else {
			if num, e := strconv.Atoi(value); num == 0 && e == nil {
				continue
			}

			lastAplha := packedString[len(packedString)-1]
			if string(lastAplha) != value && string(lastAplha) != "0" {
				if beforeAlpha := packedString[index+1]; string(beforeAlpha) == "0" {
					continue
				}
			}

			alpha = value
			builder.WriteString(strings.Repeat(alpha, 1))

		}
	}

	fmt.Println(builder.String())

	return builder.String(), nil
}
