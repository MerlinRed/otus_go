package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func checkCornerCases(value string, packedString string, index int) (string, error) {
	firstAlpha := packedString[0]
	lastAlpha := packedString[len(packedString)-1]

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

	return value, nil
}

func Unpack(packedString string) (string, error) {
	if len(packedString) == 0 {
		return packedString, nil
	}

	if len(packedString) < 2 {
		return packedString, nil
	}

	var builder strings.Builder
	var alpha string

	for index, value := range strings.Split(packedString, "") {
		if _, err := strconv.Atoi(value); err == nil {
			if _, err := checkCornerCases(value, packedString, index); err != nil {
				return "", err
			}
		}

		if digit, err := strconv.Atoi(value); err == nil && digit != 0 {
			digit--
			builder.WriteString(strings.Repeat(alpha, digit))
			continue
		}

		if num, e := strconv.Atoi(value); num == 0 && e == nil {
			continue
		}

		lastAlpha := packedString[len(packedString)-1]
		if string(lastAlpha) != value && string(lastAlpha) != "0" {
			if beforeAlpha := packedString[index+1]; string(beforeAlpha) == "0" {
				continue
			}
		}

		alpha = value
		builder.WriteString(strings.Repeat(alpha, 1))
	}

	fmt.Println(builder.String())

	return builder.String(), nil
}
