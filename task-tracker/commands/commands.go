package commands

import (
	"fmt"
	"os"
	"strconv"

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
	MARK_DONE_COMMAND        Command = "mark-done"
	MARK_IN_PROGRESS_COMMAND Command = "mark-in-progress"
	UPDATE_COMMAND           Command = "update"
)

func (c *CommandInfo) handleCommand(args []string) {
	if len(args) < c.minArgs || len(args) > c.maxArgs {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments.")
		fmt.Fprintf(os.Stderr, "Usage: \n \t%s\n", c.Usage)
		fmt.Fprintf(os.Stderr, "Example: \n \t%s\n", c.Example)
		return
	}

	switch c.Name {
	case LIST_COMMAND:
		tasks.ListTasks()
	case ADD_COMMAND:
		tasks.AddTask(args[0])
	case UPDATE_COMMAND:
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			printCommandHelp(UPDATE_COMMAND)
			return
		}
		tasks.UpdateTask(taskID, args[1])
	case MARK_DONE_COMMAND:
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			return
		}
		tasks.MarkTaskAsDone(taskID)
	case MARK_IN_PROGRESS_COMMAND:
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			return
		}
		tasks.MarkTaskAsInProgress(taskID)
	case DEL_COMMAND:
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid task ID.")
			return
		}
		tasks.DeleteTask(taskID)
	default:
		fmt.Fprintln(os.Stderr, "Command not implemented yet.")
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
	UPDATE_COMMAND: {
		Name:        UPDATE_COMMAND,
		Description: "Updates a task.",
		minArgs:     2,
		maxArgs:     2,
		Usage:       "todo update <task_id> <new_task>",
		Example:     "todo update 1 'Buy milk'",
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
