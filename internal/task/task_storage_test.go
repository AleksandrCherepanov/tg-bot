package task

import (
	"testing"
)

func TestCreateGetTask(t *testing.T) {
	tasks := map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}

	taskStorage := NewStorage()
	for text, done := range tasks {
		taskStorage.CreateTask(text, done)
	}

	for i := len(tasks) - 1; i > 0; i-- {
		task, err := taskStorage.GetTask(int64(i))
		if err != nil {
			t.Fatalf("Task create error. %s", err.Error())
		}

		if _, ok := tasks[task.Text]; !ok {
			t.Fatalf("Task create error. Actual: %v", task.Text)
		}

		if tasks[task.Text] != task.IsDone {
			t.Fatalf("Task create error. Expected: %v. Actual: %v", tasks[task.Text], task.IsDone)
		}
	}

	allTasks := taskStorage.GetAllTasks()
	if len(allTasks) != len(tasks) {
		t.Fatalf("Task create error. Expected: %v. Actual: %v", len(tasks), len(allTasks))
	}
}

func TestDeleteTask(t *testing.T) {
	tasks := []string{
		"test1",
		"test2",
		"test3",
	}

	taskStorage := NewStorage()
	for _, t := range tasks {
		taskStorage.CreateTask(t, false)
	}

	taskStorage.DeleteTask(1)
	taskStorage.DeleteTask(3)
	taskStorage.DeleteTask(4)

	_, err := taskStorage.GetTask(2)
	if err != nil {
		t.Fatalf("Task delete error. Expected: %v. Actual: %v", nil, err.Error())
	}

	_, err = taskStorage.GetTask(1)
	if err == nil {
		t.Fatalf("Task delete error. Expected: error. Actual: nil")
	}

	allTasks := taskStorage.GetAllTasks()
	if len(allTasks) != 1 {
		t.Fatalf("Task delete error. Expected: len == 1. Actual: %v", len(allTasks))
	}
}

func TestDeleteAllTask(t *testing.T) {
	tasks := []string{
		"test1",
		"test2",
		"test3",
	}

	taskStorage := NewStorage()
	for _, t := range tasks {
		taskStorage.CreateTask(t, true)
	}

	taskStorage.DeleteAllTasks()
	allTasks := taskStorage.GetAllTasks()
	if len(allTasks) != 0 {
		t.Fatalf("Task delete all error. Expected: 0. Actual: %v", len(allTasks))
	}
}
