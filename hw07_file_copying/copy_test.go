package main

import (
	"fmt"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	testCases := []struct {
		name         string
		offset       int64
		limit        int64
		expectedFile string
	}{
		{
			"offset 0 and limit 0",
			0,
			0,
			"testdata/out_offset0_limit0.txt",
		},
		{
			"offset 0 and limit 10",
			0,
			10,
			"testdata/out_offset0_limit10.txt",
		},
		{
			"offset 0 and limit 1000",
			0,
			1000,
			"testdata/out_offset0_limit1000.txt",
		},
		{
			"offset 0 and limit 10000",
			0,
			10000,
			"testdata/out_offset0_limit10000.txt",
		},
		{
			"offset 100 and limit 1000",
			100,
			1000,
			"testdata/out_offset100_limit1000.txt",
		},
		{
			"offset 6000 and limit 1000",
			6000,
			1000,
			"testdata/out_offset6000_limit1000.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "out.*.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer os.Remove(tmpFile.Name())

			copyErr := Copy("testdata/input.txt", tmpFile.Name(), testCase.offset, testCase.limit)
			require.NoError(t, copyErr, "unexpected error when copying a file")

			result, err := os.ReadFile(tmpFile.Name())
			require.NotErrorIs(t, err, fs.ErrNotExist, "file does not exist")

			expectedResult, err := os.ReadFile(testCase.expectedFile)
			if err != nil {
				t.Errorf("Can't read expected file '%v'", testCase.expectedFile)
			}
			require.Equal(t, expectedResult, result, "contents of the files do not match")
		})
	}

	t.Run("/dev/urandom", func(t *testing.T) {
		copyErr := Copy("~/dev/urandom", "new", 0, 0)
		require.Error(t, ErrUnsupportedFile, copyErr)
	})
}
