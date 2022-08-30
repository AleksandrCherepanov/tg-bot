package list

import (
	"fmt"
	"sync"
	"tg-bot/internal/task"
)

type TaskList struct {
	Id          int64
	Name        string
	TaskStorage *task.TaskStorage
}

type TaskListStorage struct {
	sync.Mutex
	lists  map[int64]TaskList
	nextId int64
}

func NewTaskListStorage() *TaskListStorage {
	tls := &TaskListStorage{}
	tls.nextId = 1
	tls.lists = make(map[int64]TaskList)

	return tls
}

func (taskListStorage *TaskListStorage) CreateList(name string) int64 {
	taskListStorage.Mutex.Lock()
	defer taskListStorage.Mutex.Unlock()

	list := TaskList{}
	list.Id = taskListStorage.nextId
	list.Name = name
	list.TaskStorage = task.NewStorage()
	taskListStorage.lists[list.Id] = list
	taskListStorage.nextId++

	return list.Id
}

func (taskListStorage *TaskListStorage) GetListById(id int64) (TaskList, error) {
	taskListStorage.Mutex.Lock()
	defer taskListStorage.Mutex.Unlock()

	taskList, ok := taskListStorage.lists[id]
	if !ok {
		return TaskList{}, fmt.Errorf("task list with id=%d not found", id)
	}

	return taskList, nil
}

func (taskListStorage *TaskListStorage) GetListAll() []TaskList {
	taskListStorage.Mutex.Lock()
	defer taskListStorage.Mutex.Unlock()

	taskList := make([]TaskList, 0)
	for _, list := range taskListStorage.lists {
		taskList = append(taskList, list)
	}

	return taskList
}

func (taskListStorage *TaskListStorage) DeleteListById(id int64) {
	taskListStorage.Mutex.Lock()
	defer taskListStorage.Mutex.Unlock()

	_, ok := taskListStorage.lists[id]
	if ok {
		delete(taskListStorage.lists, id)
	}
}

func (taskListStorage *TaskListStorage) DeleteListAll() {
	taskListStorage.Mutex.Lock()
	defer taskListStorage.Mutex.Unlock()

	taskListStorage.lists = make(map[int64]TaskList)
}
