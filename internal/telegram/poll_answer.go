package telegram

import "tg-bot/internal/telegram/user"

type PollAnswer struct {
	PollId    string    `json:"poll_id"`
	User      user.User `json:"user"`
	OptionIds []int64   `json:"option_ids"`
}
