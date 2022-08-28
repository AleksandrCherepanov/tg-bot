package telegram

import "tg-bot/pkg/telegram/user"

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       user.User       `json:"from"`
	Date       int64           `json:"date"`
	Bio        *string         `json:"bio"`
	InviteLink *ChatInviteLink `json:"invite_link"`
}
