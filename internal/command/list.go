package command

import (
	"strconv"
	"tg-bot/internal/template"
	"tg-bot/internal/user"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandList struct {
	userStorage *user.UserStorage
}

func NewCommandList() *CommandList {
	return &CommandList{
		userStorage: user.GetUserStorage(),
	}
}

func (commandList *CommandList) Handle(
	update *telegram.Update,
	command string,
	args []string,
) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}

	switch command {
	case "/l":
		{
			return commandList.getUserLists(chatId, update.Message.Chat.GetName())
		}
	case "/lc":
		{
			if len(args) != 1 {
				return client.TelegramResponse(chatId, `Invalid command arguments\.`)
			}
			return commandList.createUserList(chatId, update.Message.Chat.GetName(), args[0])
		}
	case "/ld":
		{
			if len(args) != 1 {
				return client.TelegramResponse(chatId, `Invalid command arguments\.`)
			}
			listId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return client.TelegramResponse(chatId, `Invalid command arguments\.`)
			}
			return commandList.deleteUserList(chatId, update.Message.Chat.GetName(), listId)
		}
	case "/lda":
		{
			return commandList.deleteAllUserLists(chatId, update.Message.Chat.GetName())
		}
	}

	return client.TelegramResponse(chatId, `Can't parse command\.`)
}

func (commandList *CommandList) getUserLists(userId int64, name string) (interface{}, error) {
	user, notFound := commandList.userStorage.GetUserById(userId)
	if notFound != nil {
		user = commandList.userStorage.CreateUser(userId, name)
	}

	lists, err := commandList.userStorage.GetListAllByUser(user.Id)
	if err != nil {
		return nil, err
	}

	text, err := template.NewAllListTemplate(lists).GetText()
	return client.TelegramResponse(userId, text)
}

func (commandList *CommandList) createUserList(userId int64, userName string, listName string) (interface{}, error) {
	user, notFound := commandList.userStorage.GetUserById(userId)
	if notFound != nil {
		user = commandList.userStorage.CreateUser(userId, userName)
	}

	_, err := commandList.userStorage.CreateUserList(user.Id, listName)
	if err != nil {
		return nil, err
	}

	return commandList.getUserLists(user.Id, user.Name)
}

func (commandList *CommandList) deleteUserList(userId int64, userName string, listId int64) (interface{}, error) {
	user, notFound := commandList.userStorage.GetUserById(userId)
	if notFound != nil {
		user = commandList.userStorage.CreateUser(userId, userName)
	}

	err := commandList.userStorage.DeleteUserListById(user.Id, listId)
	if err != nil {
		return nil, err
	}

	return commandList.getUserLists(user.Id, user.Name)
}

func (commandList *CommandList) deleteAllUserLists(userId int64, userName string) (interface{}, error) {
	user, notFound := commandList.userStorage.GetUserById(userId)
	if notFound != nil {
		user = commandList.userStorage.CreateUser(userId, userName)
	}

	err := commandList.userStorage.DeleteAllUserLists(user.Id)
	if err != nil {
		return nil, err
	}

	return commandList.getUserLists(user.Id, user.Name)
}
