package main

import (
	"fmt"
	"os"

	"github-user-activity.dev/github"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Invalid number of arguments:")
		fmt.Println()
		fmt.Println("Usage: go run main.go <username>")
	}
	githubUserName := args[1]

	github.PrintUserActivity(githubUserName)
}
