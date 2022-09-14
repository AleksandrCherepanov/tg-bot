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

var userStorage *UserStorage

func GetUserStorage() *UserStorage {
	if userStorage == nil {
		userStorage = &UserStorage{}
		userStorage.userList = make(map[int64]User)
	}

	return userStorage
}

func (userStorage *UserStorage) CreateUser(id int64, name string) User {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user := User{}
	user.Id = id
	user.Name = name
	user.CurrentList = nil
	user.listStorage = list.NewTaskListStorage()
	userStorage.userList[id] = user

	return user
}

func (userStorage *UserStorage) getUserById(id int64) (User, error) {
	user, ok := userStorage.userList[id]
	if !ok {
		return User{}, fmt.Errorf("user with id=%d not found", id)
	}

	return user, nil
}

func (userStorage *UserStorage) GetUserById(id int64) (User, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	return userStorage.getUserById(id)
}

func (userStorage *UserStorage) CreateUserList(userId int64, listName string) (int64, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
	if err != nil {
		return -1, err
	}

	return user.listStorage.CreateList(listName), nil
}

func (userStorage *UserStorage) SetCurrentList(userId int64, listId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
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

func (userStorage *UserStorage) getCurrentList(userId int64) (list.TaskList, error) {
	user, err := userStorage.getUserById(userId)
	if err != nil {
		return list.TaskList{}, err
	}

	if user.CurrentList == nil {
		return list.TaskList{}, fmt.Errorf("chose a list and set it as a current one")
	}

	return *user.CurrentList, err
}

func (userStorage *UserStorage) GetCurrentList(userId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	return userStorage.getCurrentList(userId)
}

func (userStorage *UserStorage) GetListAllByUser(userId int64) ([]list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
	if err != nil {
		return make([]list.TaskList, 0), err
	}

	userLists := user.listStorage.GetListAll()
	return userLists, nil
}

func (userStorage *UserStorage) GetUserListById(userId int64, listId int64) (list.TaskList, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
	if err != nil {
		return list.TaskList{}, err
	}

	userList, err := user.listStorage.GetListById(listId)
	if err != nil {
		return list.TaskList{}, err
	}

	return userList, nil
}

func (userStorage *UserStorage) DeleteUserListById(userId int64, listId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
	if err != nil {
		return err
	}

	user.listStorage.DeleteListById(listId)
	return nil
}

func (userStorage *UserStorage) DeleteAllUserLists(userId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	user, err := userStorage.getUserById(userId)
	if err != nil {
		return err
	}

	user.listStorage.DeleteListAll()
	return nil
}

func (userStorage *UserStorage) CreateUserTask(userId int64, taskText string, isDone bool) (int64, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return -1, err
	}

	return currentList.TaskStorage.CreateTask(taskText, isDone), nil
}

func (userStorage *UserStorage) GetUserTaskAll(userId int64) ([]task.Task, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return make([]task.Task, 0), err
	}

	return currentList.TaskStorage.GetAllTasks(), nil
}

func (userStorage *UserStorage) GetUserTaskById(userId int64, taskId int64) (task.Task, error) {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
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

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DeleteAllTasks()
	return nil
}

func (userStorage *UserStorage) DeleteUserTaskByTaskId(userId int64, taskId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DeleteTask(taskId)
	return nil
}

func (userStorage *UserStorage) DoneUserTask(userId int64, taskId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
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

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.DoneAll()
	return nil
}

func (userStorage *UserStorage) UndoneUserTaskAll(userId int64) error {
	userStorage.Mutex.Lock()
	defer userStorage.Mutex.Unlock()

	currentList, err := userStorage.getCurrentList(userId)
	if err != nil {
		return err
	}

	currentList.TaskStorage.UndoneAll()
	return nil
}

func (userStorage *UserStorage) Exists(userId int64) bool {
	_, ok := userStorage.userList[userId]
	return ok
}
