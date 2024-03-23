package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
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
		file, err := os.Open(filepath.Join(dir, fileData.Name()))
		if err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		rowEnv := bytes.ReplaceAll(scanner.Bytes(), []byte("\x00"), []byte("\n"))
		env := string(bytes.Trim(rowEnv, "\t"))
		if env == "" || env == " " {
			envs[fileData.Name()] = EnvValue{Value: "", NeedRemove: true}
			continue
		}

		envs[fileData.Name()] = EnvValue{Value: env, NeedRemove: false}

		if err = file.Close(); err != nil {
			return nil, err
		}
	}

	return envs, nil
}
