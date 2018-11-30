package main

import (
	"nodeone/maggie/cmd"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Take all arguments, except the application name.
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "self" {
		cmd.Execute()
	} else {
		// Send everything to lando.
		var shellCmd *exec.Cmd

		if len(args) > 0 {
			shellCmd = exec.Command("lando", strings.Join(args, " "))
		} else {
			shellCmd = exec.Command("lando")
		}

		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr

		// Run lando.
		shellCmd.Run()
	}
}
