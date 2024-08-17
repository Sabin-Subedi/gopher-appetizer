package commands

import (
	"fmt"
	"os"

	"task-tracker.dev/tasks"
)

type Command string

type CommandInfo struct {
	Name        Command
	Description string
	minArgs     int
	maxArgs     int
	Usage       string
	Example     string
}

const (
	COMMAND_HELP             Command = "help"
	ADD_COMMAND              Command = "add"
	DEL_COMMAND              Command = "delete"
	LIST_COMMAND             Command = "list"
	MARK_DONE_COMMAND        Command = "mark_done"
	MARK_IN_PROGRESS_COMMAND Command = "mark_in_progress"
)

func (c *CommandInfo) handleCommand(args []string) {
	if len(args) < c.minArgs || len(args) > c.maxArgs {
		fmt.Fprintf(os.Stderr, "Invalid number of arguments. Usage: %s\n", c.Usage)
		return
	}

	switch c.Name {
	case LIST_COMMAND:
		tasks.ListTasks()

	}
}

var CommandsMap = map[Command]CommandInfo{
	COMMAND_HELP: {
		Name:        COMMAND_HELP,
		Description: "Prints this help message.",
		minArgs:     0,
		maxArgs:     1,
		Usage:       "todo help",
		Example:     "todo help",
	},
	ADD_COMMAND: {
		Name:        ADD_COMMAND,
		Description: "Adds a new task.",
		minArgs:     1,
		maxArgs:     1,
		Usage:       "todo add <task>",
		Example:     "todo add 'Buy milk'",
	},
	DEL_COMMAND: {
		Name:        DEL_COMMAND,
		Description: "Deletes a task.",
		minArgs:     1,
		maxArgs:     1,
		Usage:       "todo delete <task_id>",
		Example:     "todo delete 1",
	},
	LIST_COMMAND: {
		Name:        LIST_COMMAND,
		Description: "Lists all tasks.",
		minArgs:     0,
		maxArgs:     0,
		Usage:       "todo list",
		Example:     "todo list",
	},
	MARK_DONE_COMMAND: {
		Name:        MARK_DONE_COMMAND,
		Description: "Marks a task as done.",
		minArgs:     1,
		Usage:       "todo mark_done <task_id>",
		Example:     "todo mark_done 1",
	},
	MARK_IN_PROGRESS_COMMAND: {
		Name:        MARK_IN_PROGRESS_COMMAND,
		Description: "Marks a task as in progress.",
		maxArgs:     1,
		Usage:       "todo mark_in_progress <task_id>",
		Example:     "todo mark_in_progress 1",
	},
}

func HandleCommand(c Command, args []string) {
	commandInfo, ok := CommandsMap[c]

	if !ok {
		fmt.Println("Invalid command provided")

		printDefaultHelp()
		return
	}

	commandInfo.handleCommand(args)
}
