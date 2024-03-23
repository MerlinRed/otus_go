package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("мало аргументов")
	}

	env, err := ReadDir(args[1])
	if err != nil {
		panic("не удалось прочесть файл")
	}

	RunCmd(args[2:], env)
}
