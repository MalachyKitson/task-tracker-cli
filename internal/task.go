package internal

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func AddTask(description string) {
	tasks, _ := LoadTasks()
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	SaveTasks(tasks)
	fmt.Println("Task added successfully (ID:", newTask.ID, ")")
}

func ListTasks() {
	tasks, _ := LoadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("[%d] %s (%s)\n", task.ID, task.Description, task.Status)
	}
}

func UpdateTask(idStr string, newDescription string) {
	tasks, _ := LoadTasks()
	id := parseID(idStr)
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			SaveTasks(tasks)
			fmt.Println("Task updated successfully.")
			return
		}
	}
	fmt.Println("Task not Found.")
}

func DeleteTask(idStr string) {
	tasks, _ := LoadTasks()
	id := parseID(idStr)
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			SaveTasks(tasks)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func MarkInProgress(idStr string) {
	UpdateStatus(idStr, "in-progress")
}

func MarkDone(idStr string) {
	UpdateStatus(idStr, "done")
}

func UpdateStatus(idstr string, status string) {
	tasks, _ := LoadTasks()
	id := parseID(idstr)
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			SaveTasks(tasks)
			fmt.Println("Task marked as", status)
			return
		}
	}
	fmt.Println("Task not found.")
}
