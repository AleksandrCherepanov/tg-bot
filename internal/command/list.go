package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram/client"
	"github.com/AleksandrCherepanov/tg-bot/internal/template"
	"github.com/AleksandrCherepanov/tg-bot/internal/user"
)

type CommandList struct {
	chatId      int64
	userStorage *user.UserStorage
}

func NewCommandList(chatId int64) *CommandList {
	return &CommandList{
		chatId:      chatId,
		userStorage: user.GetUserStorage(),
	}
}

func (c *CommandList) throwTelegramError(text string) error {
	return client.NewTelegramResponse(c.chatId, text, true)
}

func (c *CommandList) Handle(command string, args []string) (interface{}, error) {
	switch command {
	case "/l":
		{
			return c.getUserLists(c.chatId)
		}
	case "/lc":
		{
			if len(args) < 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.createUserList(c.chatId, args)
		}
	case "/ld":
		{
			if len(args) != 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			listId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.deleteUserList(c.chatId, listId)
		}
	case "/ls":
		{
			if len(args) != 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			listId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.setCurrentList(c.chatId, listId)
		}
	case "/lda":
		{
			return c.deleteAllUserLists(c.chatId)
		}
	}

	return nil, client.NewTelegramResponse(c.chatId, `Can't parse command\.`, true)
}

func (c *CommandList) getUserLists(userId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	lists, err := c.userStorage.GetListAllByUser(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Can't get user's lists")
	}

	text, err := template.NewAllListTemplate(lists).GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(userId, text, true), nil
}

func (c *CommandList) createUserList(userId int64, listNameParts []string) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	listName := strings.Join(listNameParts, " ")
	_, err = c.userStorage.CreateUserList(user.Id, listName)
	if err != nil {
		return nil, c.throwTelegramError("Can't create user's list")
	}

	return c.getUserLists(user.Id)
}

func (c *CommandList) deleteUserList(userId int64, listId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User nor found. Use command `/start` to create an user")
	}

	err = c.userStorage.DeleteUserListById(user.Id, listId)
	if err != nil {
		return nil, c.throwTelegramError("Can't delete user's list")
	}

	return c.getUserLists(user.Id)
}

func (c *CommandList) deleteAllUserLists(userId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	err = c.userStorage.DeleteAllUserLists(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Can't delete user's lists")
	}

	return c.getUserLists(user.Id)
}

func (c *CommandList) setCurrentList(userId int64, listId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	_, err = c.userStorage.SetCurrentList(user.Id, listId)
	if err != nil {
		return nil, c.throwTelegramError(fmt.Sprintf("Can't set list %d as the current one", listId))
	}

	response := client.NewTelegramResponse(c.chatId, fmt.Sprintf("List %d has set as the current one", listId), false)
	response = *response.WithPinMessage()
	return response, nil
}
