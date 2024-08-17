package commands

import "fmt"

func printASCITitle() {
	title := ` _____            _      _____                   _
|_   _|__ _  ___ | | __ |_   _|_ __  __ _   ___ | | __ ___  _ __
  | | / _' |/ __|| |/ /   | | | '__|/ _' | / __|| |/ // _ \| '__|
  | || (_| |\__ \|   <    | | | |  | (_| || (__ |   <|  __/| |
  |_| \__,_||___/|_|\_\   |_| |_|   \__,_| \___||_|\_\\___||_|  `

	fmt.Println(title)
}

func HandleHelpCommand(args []string) {
	if len(args) > 0 {
		printCommandHelp(Command(args[0]))
	} else {
		printDefaultHelp()
	}
}

func printDefaultHelp() {
	printASCITitle()
	fmt.Print("\n\n")
	fmt.Println("Usage: \n todo <command> [args]")
	fmt.Print("\n")
	printAvailableCommands()
}

func printAvailableCommands() {
	fmt.Println("Available commands:")
	for _, command := range CommandsMap {
		fmt.Printf(" %-30s %-20s \n", command.Name, command.Description)
	}
}

func printCommandHelp(command Command) {
	if commandInfo, ok := CommandsMap[command]; ok {
		fmt.Printf("Command: \n \t<%s>\n", commandInfo.Name)
		fmt.Printf("Description: \n \t%s\n", commandInfo.Description)
		fmt.Printf("Usage: \n \t%s\n", commandInfo.Usage)
		fmt.Printf("Example: \n \t%s\n", commandInfo.Example)
	} else {
		fmt.Println("Invalid command provided")
		printDefaultHelp()
	}
}
