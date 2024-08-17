package main

import (
	"os"

	"task-tracker.dev/commands"
)

func main() {
	args := os.Args
	command := ""

	if len(args) > 1 {
		command = args[1]
	}

	commandArgs := []string{}

	if len(args) > 2 {
		commandArgs = args[2:]
	}

	commands.HandleCommand(commands.Command(command), commandArgs)

}
