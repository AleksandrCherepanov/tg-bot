package user

import (
	"fmt"
	"sync"
	"tg-bot/internal/list"
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
