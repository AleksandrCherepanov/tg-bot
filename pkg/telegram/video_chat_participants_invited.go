package telegram

import "tg-bot/pkg/telegram/user"

type VideoChatParticipantsInvited struct {
	Users []user.User `json:"users"`
}
