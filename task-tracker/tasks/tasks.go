package tasks

import "fmt"

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "TO DO"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN PROGRESS"
	TASK_STATUS_DONE        TaskStatus = "DONE"
)

type Task struct {
	ID     int        `json:"id"`
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"`
}

var tasks = readTaskFromJSONFile()

func ListTasks() {
	fmt.Printf("%-5s %-50s %-15s\n", "ID", "Task", "Status")
	for _, task := range tasks {
		fmt.Println(task.ID, task.Task, task.Status)
	}
}
