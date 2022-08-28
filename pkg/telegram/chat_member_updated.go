package telegram

import "tg-bot/pkg/telegram/user"

type ChatMemberUpdated struct {
	Chat          Chat           `json:"chat"`
	From          *user.User     `json:"from"`
	Date          int64          `json:"date"`
	OldChatMember interface{}    `json:"old_chat_member"`
	NewChatMember interface{}    `json:"new_chat_member"`
	InviteLink    ChatInviteLink `json:"invite_link"`
}
