package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("less than 2 arguments passed")
	}

	env, err := ReadDir(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	RunCmd(args[2:], env)
}
