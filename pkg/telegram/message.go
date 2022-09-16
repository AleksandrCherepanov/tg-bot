package telegram

import (
	"fmt"
	"tg-bot/pkg/telegram/passport"
	"tg-bot/pkg/telegram/payment"
	"tg-bot/pkg/telegram/user"
)

type Message struct {
	MessageId                     int64                          `json:"message_id"`
	From                          *user.User                     `json:"from"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	Date                          int64                          `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *user.User                     `json:"forward_from"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat"`
	ForwardFromMessageId          *int64                         `json:"forward_from_message_id"`
	ForwardSignature              *string                        `json:"forward_signature"`
	ForwardSenderName             *string                        `json:"forward_sender_name"`
	ForwardDate                   *int64                         `json:"forward_date"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward"`
	ReplyToMessage                *Message                       `json:"reply_to_message"`
	ViaBot                        *user.User                     `json:"via_bot"`
	EditDate                      *int64                         `json:"edit_date"`
	HasProtectedContent           *bool                          `json:"has_protected_content"`
	MediaGroupId                  *string                        `json:"media_group_id"`
	AuthorSignature               *string                        `json:"author_signature"`
	Text                          *string                        `json:"text"`
	Entities                      []MessageEntity                `json:"entities"`
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	Document                      *Document                      `json:"document"`
	Photo                         []PhotoSize                    `json:"photo"`
	Sticker                       *Sticker                       `json:"sticker"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	Caption                       *string                        `json:"caption"`
	CaptionEntities               []MessageEntity                `json:"caption_entities"`
	Contact                       *Contact                       `json:"contact"`
	Dice                          *Dice                          `json:"dice"`
	Game                          *Game                          `json:"game"`
	Poll                          *Poll                          `json:"poll"`
	Venue                         *Venue                         `json:"venue"`
	Location                      *Location                      `json:"location"`
	NewChatMembers                []user.User                    `json:"new_chat_members"`
	LeftChatMember                *user.User                     `json:"left_chat_member"`
	NewChatTitle                  *string                        `json:"new_chat_title"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo"`
	GroupChatCreated              *bool                          `json:"group_chat_created"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatId               *int64                         `json:"migrate_to_chat_id"`
	MigrateFromChatId             *int64                         `json:"migrate_from_chat_id"`
	PinnedMessage                 *Message                       `json:"pinned_message"`
	Invoice                       *payment.Invoice               `json:"invoice"`
	SuccessfulPayment             *payment.SuccessfulPayment     `json:"successful_payment"`
	ConnectedWebsite              *string                        `json:"connected_website"`
	PassportData                  *passport.PassportData         `json:"passport_data"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
}

func (m *Message) GetChatId() (int64, error) {
	if m.Chat == nil {
		return -1, fmt.Errorf("Chat is not defined")
	}

	return m.Chat.Id, nil
}
