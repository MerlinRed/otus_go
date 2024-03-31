package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "good",
				Name:   "good",
				Age:    18,
				Email:  "good@mail.ru",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "1234567890123456789012345678901234567",
				Name:   "good",
				Age:    18,
				Email:  "good@mail.ru",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrorLen,
				},
			},
		},
		{
			in: User{
				ID:     "123",
				Name:   "good",
				Age:    1,
				Email:  "good@mail.ru",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   ErrorMin,
				},
			},
		},
		{
			in: User{
				ID:     "123",
				Name:   "good",
				Age:    51,
				Email:  "good@mail.ru",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Age",
					Err:   ErrorMax,
				},
			},
		},
		{
			in: User{
				ID:     "123",
				Name:   "good",
				Age:    18,
				Email:  "good@mail",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Email",
					Err:   ErrorRegexp,
				},
			},
		},
		{
			in: User{
				ID:     "123",
				Name:   "good",
				Age:    18,
				Email:  "good@mail.ru",
				Role:   "bad",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Role",
					Err:   ErrorIn,
				},
			},
		},
		{
			in: User{
				ID:     "123",
				Name:   "good",
				Age:    18,
				Email:  "good@mail.ru",
				Role:   "admin",
				Phones: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
				meta:   []byte{},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Phones",
					Err:   ErrorLen,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			errs := Validate(tt.in)
			require.Equal(t, tt.expectedErr, errs)
		})
	}
}

func TestValidateOther(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: App{Version: "123456"},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   ErrorLen,
				},
			},
		},
		{
			in: Token{
				Header:    []byte{'1', '2', '3'},
				Payload:   []byte{'4', '5', '6'},
				Signature: []byte{'7', '8', '9'},
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				Body: "",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 300,
				Body: "",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   ErrorIn,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			errs := Validate(tt.in)
			require.Equal(t, tt.expectedErr, errs)
		})
	}
}
