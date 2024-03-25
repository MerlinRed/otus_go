package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	cases := []struct {
		testName   string
		cmd        []string
		returnCode int
	}{
		{
			testName:   "good",
			cmd:        []string{"/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"},
			returnCode: 0,
		},
		{
			testName:   "bad",
			cmd:        []string{"bad command"},
			returnCode: 1,
		},
		{
			testName:   "not found command file",
			cmd:        []string{"/bin/bash", "./testdata/not_exists.sh", "arg1=1", "arg2=2"},
			returnCode: 1,
		},
	}

	env, err := ReadDir("./testdata/env")
	require.Nil(t, err)

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			ComArgVar := RunCmd(c.cmd, env)
			require.Equal(t, c.returnCode, ComArgVar)
		})
	}
}
