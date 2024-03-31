package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrorLen    = errors.New("len error")
	ErrorMin    = errors.New("min error")
	ErrorMax    = errors.New("max error")
	ErrorRegexp = errors.New("regexp error")
	ErrorIn     = errors.New("in error")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	strResult := strings.Builder{}
	for _, err := range v {
		if err.Err != nil {
			strResult.WriteString(fmt.Sprintf("ValidationError: %s: %s\n", err.Field, err.Err))
		}
	}

	return strResult.String()
}

func Validate(v interface{}) error {
	if reflect.TypeOf(v).Kind() != reflect.Struct {
		return nil
	}
	structType := reflect.ValueOf(v)

	valErrors := ValidationErrors{}

	for i := 0; i < structType.NumField(); i++ {
		fieldContains := structType.Field(i)
		field := structType.Type().Field(i)
		tag := field.Tag.Get("validate")
		if tag == "" {
			continue
		}

		for _, value := range strings.Split(tag, "|") {
			err := validateField(value, fieldContains, field.Name)
			if err.Err != nil {
				valErrors = append(
					valErrors,
					err,
				)
			}
		}
	}
	if len(valErrors) != 0 {
		return valErrors
	}

	return nil
}

func validateField(value string, fieldContains reflect.Value, fieldName string) ValidationError {
	switch {
	case strings.HasPrefix(value, "len:"):
		expectedLen, _ := strconv.Atoi(strings.Split(value, ":")[1])
		if fieldContains.Len() > expectedLen {
			return ValidationError{Field: fieldName, Err: ErrorLen}
		}
	case strings.HasPrefix(value, "min:"):
		minValue, _ := strconv.Atoi(strings.Split(value, ":")[1])
		if fieldContains.Interface().(int) < minValue {
			return ValidationError{Field: fieldName, Err: ErrorMin}
		}
	case strings.HasPrefix(value, "max:"):
		maxValue, _ := strconv.Atoi(strings.Split(value, ":")[1])
		if fieldContains.Interface().(int) > maxValue {
			return ValidationError{Field: fieldName, Err: ErrorMax}
		}
	case strings.HasPrefix(value, "regexp:"):
		pattern := strings.Split(value, ":")[1]
		re := regexp.MustCompile(pattern)
		if !re.MatchString(fieldContains.String()) {
			return ValidationError{Field: fieldName, Err: ErrorRegexp}
		}
	case strings.HasPrefix(value, "in:"):
		validValues := strings.Split(strings.Split(value, ":")[1], ",")
		var preparedField string
		switch fieldContains.Kind() { //nolint:exhaustive
		case reflect.String:
			preparedField = fieldContains.String()
		case reflect.Int:
			preparedField = strconv.Itoa(int(fieldContains.Int()))
		}

		if !slices.Contains(validValues, preparedField) {
			return ValidationError{Field: fieldName, Err: ErrorIn}
		}
	}
	return ValidationError{}
}
