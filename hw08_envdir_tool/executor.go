package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(name)
			if err != nil {
				return 1
			}
			continue
		}
		err := os.Setenv(name, value.Value)
		if err != nil {
			return 1
		}
	}

	commandName := cmd[0]
	args := cmd[1:]

	command := exec.Command(commandName, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = os.Environ()

	err := command.Run()
	if err != nil {
		return 1
	}

	return command.ProcessState.ExitCode()
}
