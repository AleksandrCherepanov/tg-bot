package command

import (
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram/client"
	"github.com/AleksandrCherepanov/tg-bot/internal/template"
	"github.com/AleksandrCherepanov/tg-bot/internal/user"
)

type CommandTask struct {
	chatId      int64
	userStorage *user.UserStorage
}

func NewCommandTask(chatId int64) *CommandTask {
	return &CommandTask{
		chatId:      chatId,
		userStorage: user.GetUserStorage(),
	}
}

func (c *CommandTask) throwTelegramError(text string) error {
	return client.NewTelegramResponse(c.chatId, text, true)
}

func (c *CommandTask) Handle(command string, args []string) (interface{}, error) {
	switch command {
	case "/t":
		{
			return c.getUserTasks(c.chatId)
		}
	case "/tc":
		{
			if len(args) < 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.createUserTask(c.chatId, args)
		}
	case "/td":
		{
			if len(args) != 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			taskId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.deleteUserTask(c.chatId, taskId)
		}
	case "/tda":
		{
			return c.deleteAllUserTasks(c.chatId)
		}
	case "/tm":
		{
			if len(args) != 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			taskId, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.markUserTask(c.chatId, taskId)
		}
	case "/tma":
		{
			if len(args) != 1 {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			flag, err := strconv.ParseInt(args[0], 10, 0)
			if err != nil {
				return nil, c.throwTelegramError(`Invalid command arguments\.`)
			}
			return c.markAllUserTasks(c.chatId, int(flag))
		}
	}

	return nil, client.NewTelegramResponse(c.chatId, `Can't parse command\.`, true)
}

func (c *CommandTask) getUserTasks(userId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	tasks := currentList.TaskStorage.GetAllTasks()
	text, err := template.NewAllTaskTemplate(tasks).GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(userId, text, true), nil
}

func (c *CommandTask) createUserTask(userId int64, taskNameParts []string) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	taskName := strings.Join(taskNameParts, " ")
	_ = currentList.TaskStorage.CreateTask(taskName, false)
	if err != nil {
		return nil, c.throwTelegramError("Can't create user's task")
	}

	return c.getUserTasks(user.Id)
}

func (c *CommandTask) deleteUserTask(userId int64, taskId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User nor found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	currentList.TaskStorage.DeleteTask(taskId)
	if err != nil {
		return nil, c.throwTelegramError("Can't delete user's task")
	}

	return c.getUserTasks(user.Id)
}

func (c *CommandTask) deleteAllUserTasks(userId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	currentList.TaskStorage.DeleteAllTasks()
	if err != nil {
		return nil, c.throwTelegramError("Can't delete user's tasks")
	}

	return c.getUserTasks(user.Id)
}

func (c *CommandTask) markUserTask(userId int64, taskId int64) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	currentList.TaskStorage.Mark(taskId)
	if err != nil {
		return nil, c.throwTelegramError("Can't delete user's tasks")
	}

	return c.getUserTasks(user.Id)
}

func (c *CommandTask) markAllUserTasks(userId int64, flag int) (interface{}, error) {
	user, err := c.userStorage.GetUserById(userId)
	if err != nil {
		return nil, c.throwTelegramError("User not found. Use command `/start` to create an user")
	}

	currentList, err := c.userStorage.GetCurrentList(user.Id)
	if err != nil {
		return nil, c.throwTelegramError("Current list is not set. Use `/ls id` command to set current list")
	}

	if flag != 0 && flag != 1 {
		return nil, c.throwTelegramError("Flag must be 0 or 1")
	}

	if flag == 1 {
		currentList.TaskStorage.DoneAll()
	} else {
		currentList.TaskStorage.UndoneAll()
	}

	return c.getUserTasks(user.Id)
}
