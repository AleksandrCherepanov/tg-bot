package telegram

import "tg-bot/internal/telegram/user"

type MessageEntity struct {
	Type          string     `json:"type"`
	Offset        int64      `json:"offset"`
	Length        int64      `json:"length"`
	Url           *string    `json:"url"`
	User          *user.User `json:"user"`
	Language      *string    `json:"language"`
	CustomEmojiId *string    `json:"custom_emoji_id"`
}
