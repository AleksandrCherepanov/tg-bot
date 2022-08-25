package command

import (
	"log"
	"tg-bot/internal/telegram"
)

type CommandStart struct {
}

func NewCommandStart() *CommandStart {
	return &CommandStart{}
}

func (commandStart *CommandStart) Handle(update *telegram.Update) (interface{}, error) {
	log.Println("START COMMAND IS CALLED")
	return "{}", nil
}
