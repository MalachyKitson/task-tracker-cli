package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <task description>")
			return
		}
		AddTask(os.Args[2])

	case "list":
		ListTasks()

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <new description>")
			return
		}
		UpdateTask(os.Args[2], os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		DeleteTask(os.Args[2])

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		MarkInProgress(os.Args[2])

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		MarkDone(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}
