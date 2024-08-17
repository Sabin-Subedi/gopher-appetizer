package tasks

import (
	"fmt"
	"os"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "TO DO"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN PROGRESS"
	TASK_STATUS_DONE        TaskStatus = "DONE"
)

type Task struct {
	ID        int        `json:"id"`
	Task      string     `json:"task"`
	Status    TaskStatus `json:"status"`
	CreatedAt string     `json:"created_at"`
	UpdateAt  string     `json:"updated_at"`
}

var tasks = readTaskFromJSONFile()

func ListTasks() {
	fmt.Printf("%-5s %-50s %-15s\n", "ID", "Task", "Status")
	for _, task := range tasks {
		fmt.Printf("%-5d %-50s %-15s\n", task.ID, task.Task, task.Status)
	}
}

func AddTask(task string) {
	newTask := Task{
		ID:        tasks[len(tasks)-1].ID + 1,
		Task:      task,
		Status:    TASK_STATUS_TODO,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	tasks = append(tasks, newTask)

	writeTaskToJSONFile(tasks)
	fmt.Printf("Task added successfully. (ID: %d)\n", newTask.ID)
}

func FindTaskByID(taskID int) (*Task, int) {
	for index, task := range tasks {
		if task.ID == taskID {
			return &task, index
		}
	}

	return nil, -1
}

func MarkTaskAsDone(taskID int) {
	task, index := FindTaskByID(taskID)
	if task == nil {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	task.Status = TASK_STATUS_DONE
	tasks[index] = *task
	writeTaskToJSONFile(tasks)
	fmt.Println("Hooray 🎉! You completed the task.")
}

func MarkTaskAsInProgress(taskID int) {
	task, index := FindTaskByID(taskID)
	if task == nil {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	task.Status = TASK_STATUS_IN_PROGRESS
	tasks[index] = *task
	writeTaskToJSONFile(tasks)
	fmt.Printf("Task marked as in progress. (ID:%d)\n", task.ID)
}

func DeleteTask(taskID int) {
	task, index := FindTaskByID(taskID)
	if task == nil {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	writeTaskToJSONFile(tasks)
	fmt.Printf("Task deleted successfully. (ID:%d)\n", task.ID)
}

func UpdateTask(taskID int, newTask string) {
	task, index := FindTaskByID(taskID)
	if task == nil {
		fmt.Fprintln(os.Stderr, "Task not found.")
		return
	}

	task.Task = newTask
	task.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	tasks[index] = *task
	writeTaskToJSONFile(tasks)
	fmt.Printf("Task updated successfully. (ID:%d)\n", task.ID)
}
