package user

import (
	"fmt"
	"sync"
	"tg-bot/internal/list"
	"tg-bot/internal/task"
)

type User struct {
	Id          int64
	Name        string
	CurrentList *list.TaskList
	listStorage *list.TaskListStorage
}

type UserStorage struct {
	sync.Mutex
	userList map[int64]User
}

func (userStorage *UserStorage) CreateUser(id int64, name string) int64 {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user := User{}
	user.Id = id
	user.Name = name
	user.CurrentList = nil
	user.listStorage = list.NewTaskListStorage()

	return user.Id
}

func (userStorage *UserStorage) GetUserById(id int64) (User, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, ok := userStorage.userList[id]
	if !ok {
		return User{}, fmt.Errorf("user with id=%d not found", id)
	}

	return user, nil
}

func (userStorage *UserStorage) CreateUserList(userId int64, listName string) (int64, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.GetUserById(userId)
	if err != nil {
		return -1, err
	}

	return user.listStorage.CreateList(listName), nil
}

func (userStorage *UserStorage) SetCurrentList(userId int64, listId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.GetUserById(userId)
	if err != nil {
		return list.TaskList{}, err
	}

	userList, err := user.listStorage.GetListById(listId)
	if err != nil {
		return list.TaskList{}, err
	}

	user.CurrentList = &userList
	return userList, nil
}

func (userStorage *UserStorage) GetCurrentList(userId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.GetUserById(userId)
	if err != nil {
		return list.TaskList{}, err
	}

	if user.CurrentList == nil {
		return list.TaskList{}, fmt.Errorf("chose a list and set it as a current one")
	}

	return *user.CurrentList, err
}

func (userStorage *UserStorage) GetListAllByUser(userId int64) ([]list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.GetUserById(userId)
	if err != nil {
		return make([]list.TaskList, 0), err
	}

	userLists := user.listStorage.GetListAll()
	return userLists, nil
}

func (userStorage *UserStorage) GetUserListById(userId int64, listId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.GetUserById(userId)
	if err != nil {
		return list.TaskList{}, err
	}

	userList, err := user.listStorage.GetListById(listId)
	if err != nil {
		return list.TaskList{}, err
	}

	return userList, nil
}

func (userStorage *UserStorage) CreateUserTask(userId int64, taskText string, isDone bool) (int64, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return -1, err
	}

	return currentList.TaskStorage.CreateTask(taskText, isDone), nil
}

func (userStorage *UserStorage) GetUserTaskAll(userId int64) ([]task.Task, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return make([]task.Task, 0), err
	}

	return currentList.TaskStorage.GetAllTasks(), nil
}

func (userStorage *UserStorage) GetUserTaskById(userId int64, taskId int64) (task.Task, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return task.Task{}, err
	}

	userTask, err := currentList.TaskStorage.GetTask(taskId)
	if err != nil {
		return task.Task{}, err
	}

	return userTask, nil
}

func (userStorage *UserStorage) DeleteUserTaskAll(userId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DeleteAllTasks()
	return nil
}

func (userStorage *UserStorage) DeleteUserTaskByTaskId(userId int64, taskId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DeleteTask(taskId)
	return nil
}

func (userStorage *UserStorage) DoneUserTask(userId int64, taskId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return err
	}

	err = currentList.TaskStorage.Done(taskId)
	if err != nil {
		return err
	}

	return nil
}

func (userStorage *UserStorage) DoneUserTaskAll(userId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DoneAll()
	return nil
}

func (userStorage *UserStorage) UndoneUserTaskAll(userId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.GetCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.UndoneAll()
	return nil
}
