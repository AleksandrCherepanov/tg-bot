package telegram

import "tg-bot/internal/telegram/user"

type VideoChatParticipantsInvited struct {
	Users []user.User `json:"users"`
}
