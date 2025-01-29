# Task Tracker CLI
This is a solution to the [Task Tracker CLI](https://roadmap.sh/projects/task-tracker) from [roadmap.sh](https://roadmap.sh). It is a program used to manage tasks and works on backend development skills including working with the filesystem, handling user inputs, and building a simple CLI application.

## How to run this program
Clone this repository and build the project as a .exe file in the terminal
```bash
git clone https://github.com/MalachyKitson/task-tracker-cli.git
cd task-tracker-cli
go build -o task-cli.exe
```

The list of commands and their usage in the terminal is given below:
```bash
# Adding a new task
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1

# Marking a task as in progress or done
./task-cli mark-in-progress 1
./task-cli mark-done 1

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```
