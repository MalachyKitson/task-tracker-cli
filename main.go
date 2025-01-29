package main

import (
	"fmt"
	"os"
	"task-tracker-cli/internal"
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
		internal.AddTask(os.Args[2])

	case "list":
		if len(os.Args) > 3 {
			fmt.Println("Usage: task-cli list or task-cli list <task status>")
			return
		}

		if len(os.Args) == 2 {
			internal.ListTasks("")
		} else if len(os.Args) == 3 {
			internal.ListTasks(os.Args[2])
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <new description>")
			return
		}
		internal.UpdateTask(os.Args[2], os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		internal.DeleteTask(os.Args[2])

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		internal.MarkInProgress(os.Args[2])

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		internal.MarkDone(os.Args[2])

	default:
		fmt.Println("Unknown command:", command)
	}
}
