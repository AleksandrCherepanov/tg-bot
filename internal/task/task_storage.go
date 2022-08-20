package task

import (
	"fmt"
	"sync"
)

type Task struct {
	Id   int64
	Text string
}

func New(id int64, text string) Task {
	return Task{
		Id:   id,
		Text: text,
	}
}

type TaskStorage struct {
	sync.Mutex
	taskList map[int64]Task
	nextId   int64
}

func NewStorage() *TaskStorage {
	ts := &TaskStorage{}
	ts.taskList = make(map[int64]Task)
	ts.nextId = 1
	return ts
}

func (taskStorage *TaskStorage) CreateTask(text string) int64 {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	task := New(taskStorage.nextId, text)
	taskStorage.taskList[taskStorage.nextId] = task
	taskStorage.nextId++
	return task.Id
}

func (taskStorage *TaskStorage) GetTask(id int64) (Task, error) {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	task, ok := taskStorage.taskList[id]
	if ok {
		return task, nil
	}

	return Task{}, fmt.Errorf("task with id=%d not found", id)
}

func (taskStorage *TaskStorage) DeleteTask(id int64) {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	if _, ok := taskStorage.taskList[id]; ok {
		delete(taskStorage.taskList, id)
	}
}

func (taskStorage *TaskStorage) GetAllTasks() []Task {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	tasks := make([]Task, 0, 0)
	for _, task := range taskStorage.taskList {
		tasks = append(tasks, task)
	}

	return tasks
}

func (taskStorage *TaskStorage) DeleteAllTasks() {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	taskList := make(map[int64]Task)
	taskStorage.taskList = taskList
}
