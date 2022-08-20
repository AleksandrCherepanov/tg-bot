package user

import (
	"sync"
	"tg-bot/internal/list"
)

type User struct {
	Id          int64
	Name        string
	CurrentList *list.TaskList
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

	return user.Id
}
