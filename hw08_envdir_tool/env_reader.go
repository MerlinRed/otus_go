package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)

	for _, fileData := range files {
		file, err := os.Open(dir + "/" + fileData.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()

		env := string(bytes.ReplaceAll(scanner.Bytes(), []byte("\x00"), []byte("\n")))
		fileName := strings.Split(file.Name(), "/")[3]

		if env == "" || env == " " {
			envs[fileName] = EnvValue{Value: "", NeedRemove: true}
		} else {
			envs[fileName] = EnvValue{Value: env, NeedRemove: false}
		}
	}

	return envs, nil
}
