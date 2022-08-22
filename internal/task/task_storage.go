package task

import (
	"fmt"
	"sync"
)

type Task struct {
	Id     int64
	Text   string
	IsDone bool
}

func New(id int64, text string, isDone bool) Task {
	return Task{
		Id:     id,
		Text:   text,
		IsDone: isDone,
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

func (taskStorage *TaskStorage) CreateTask(text string, isDone bool) int64 {
	taskStorage.Mutex.Lock()
	defer taskStorage.Mutex.Unlock()

	task := New(taskStorage.nextId, text, isDone)
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

func (taskStorage *TaskStorage) Done(id int64) error {
	task, err := taskStorage.GetTask(id)
	if err != nil {
		return err
	}

	task.IsDone = true
	taskStorage.taskList[id] = task
	return nil
}

func (taskStorage *TaskStorage) DoneAll() {
	for i, t := range taskStorage.taskList {
		t.IsDone = true
		taskStorage.taskList[i] = t
	}
}

func (taskStorage *TaskStorage) Undone(id int64) error {
	task, err := taskStorage.GetTask(id)
	if err != nil {
		return err
	}

	task.IsDone = false
	taskStorage.taskList[id] = task
	return nil
}

func (taskStorage *TaskStorage) UndoneAll() {
	for i, t := range taskStorage.taskList {
		t.IsDone = false
		taskStorage.taskList[i] = t
	}
}
