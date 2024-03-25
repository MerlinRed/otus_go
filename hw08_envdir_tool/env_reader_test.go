package main

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		envs, err := ReadDir("./testdata/env")

		require.Nil(t, err)
		require.Equal(t, Environment{
			"BAR":   EnvValue{"bar", false},
			"EMPTY": EnvValue{"", true},
			"FOO":   EnvValue{"   foo\nwith new line", false},
			"HELLO": EnvValue{"\"hello\"", false},
			"UNSET": EnvValue{"", true},
		}, envs)
	})

	t.Run("bad", func(t *testing.T) {
		_, err := ReadDir("./testdata/not_found")
		require.Error(t, &fs.PathError{}, err)
	})
}
