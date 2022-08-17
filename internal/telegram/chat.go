package telegram

import "strings"

type Chat struct {
	Id                                 int64            `json:"id"`
	Type                               string           `json:"type"`
	Title                              *string          `json:"title"`
	Username                           *string          `json:"username"`
	FirstName                          *string          `json:"first_name"`
	LastName                           *string          `json:"last_name"`
	Photo                              *ChatPhoto       `json:"photo"`
	Bio                                *string          `json:"bio"`
	HasPrivateForwards                 *bool            `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages *bool            `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 *bool            `json:"join_to_send_messages"`
	JoinByRequest                      *bool            `json:"join_by_request"`
	Description                        *string          `json:"description"`
	InviteLink                         *string          `json:"invite_link"`
	PinnedMessage                      *Message         `json:"pinned_message"`
	Permissions                        *ChatPermissions `json:"permissions"`
	SlowModeDelay                      *int64           `json:"slow_mode_delay"`
	MessageAutoDeleteTime              *int64           `json:"message_auto_delete_time"`
	HasProtectedContent                *bool            `json:"has_protected_content"`
	StickerSetName                     *string          `json:"sticker_set_name"`
	CanSetStickerSet                   *bool            `json:"can_set_sticker_set"`
	LinkedChatId                       *int64           `json:"linked_chat_id"`
	Location                           *ChatLocation    `json:"location"`
}

func (c *Chat) GetName() string {
	nameParts := make([]string, 0, 0)
	if c.FirstName != nil {
		nameParts = append(nameParts, *c.FirstName)
	}

	if c.LastName != nil {
		nameParts = append(nameParts, *c.LastName)
	}
	
	if len(nameParts) == 0 && c.Username != nil {
		nameParts = append(nameParts, *c.Username)
	}

	if len(nameParts) == 0 {
		nameParts = append(nameParts, "anonymous user")
	}
	
	return strings.Join(nameParts, " ")
}
