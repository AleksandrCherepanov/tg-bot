package command

import (
	"strconv"
	"strings"
	"tg-bot/internal/template"
	"tg-bot/internal/user"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandList struct {
	chatId      int64
	message     *telegram.Message
	userStorage *user.UserStorage
}

func NewCommandList(chatId int64, message *telegram.Message) *CommandList {
	return &CommandList{
		chatId:      chatId,
		message:     message,
		userStorage: user.GetUserStorage(),
	}
}

func (c *CommandList) Handle(command string, args []string) (interface{}, error) {
	switch command {
	case "/l":
		{
			return c.getUserLists(c.chatId, c.message.Chat.GetName())
		}
	case "/lc":
		{
			if len(args) < 1 {
				return client.NewTelegramResponse(c.chatId, `Invalid command arguments\.`, true), nil
			}
			return c.createUserList(c.chatId, c.message.Chat.GetName(), args)
		}
	case "/ld":
		{
			if len(args) != 1 {
				return client.NewTelegramResponse(c.chatId, `Invalid command arguments\.`, true), nil
			}
			listId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return client.NewTelegramResponse(c.chatId, `Invalid command arguments\.`, true), nil
			}
			return c.deleteUserList(c.chatId, c.message.Chat.GetName(), listId)
		}
	case "/lda":
		{
			return c.deleteAllUserLists(c.chatId, c.message.Chat.GetName())
		}
	}

	return client.NewTelegramResponse(c.chatId, `Can't parse command\.`, true), nil
}

func (c *CommandList) getUserLists(userId int64, name string) (interface{}, error) {
	user, notFound := c.userStorage.GetUserById(userId)
	if notFound != nil {
		user = c.userStorage.CreateUser(userId, name)
	}

	lists, err := c.userStorage.GetListAllByUser(user.Id)
	if err != nil {
		return nil, err
	}

	text, err := template.NewAllListTemplate(lists).GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(userId, text, true), nil
}

func (c *CommandList) createUserList(userId int64, userName string, listNameParts []string) (interface{}, error) {
	user, notFound := c.userStorage.GetUserById(userId)
	if notFound != nil {
		user = c.userStorage.CreateUser(userId, userName)
	}

	listName := strings.Join(listNameParts, " ")
	_, err := c.userStorage.CreateUserList(user.Id, listName)
	if err != nil {
		return nil, err
	}

	return c.getUserLists(user.Id, user.Name)
}

func (c *CommandList) deleteUserList(userId int64, userName string, listId int64) (interface{}, error) {
	user, notFound := c.userStorage.GetUserById(userId)
	if notFound != nil {
		user = c.userStorage.CreateUser(userId, userName)
	}

	err := c.userStorage.DeleteUserListById(user.Id, listId)
	if err != nil {
		return nil, err
	}

	return c.getUserLists(user.Id, user.Name)
}

func (c *CommandList) deleteAllUserLists(userId int64, userName string) (interface{}, error) {
	user, notFound := c.userStorage.GetUserById(userId)
	if notFound != nil {
		user = c.userStorage.CreateUser(userId, userName)
	}

	err := c.userStorage.DeleteAllUserLists(user.Id)
	if err != nil {
		return nil, err
	}

	return c.getUserLists(user.Id, user.Name)
}
