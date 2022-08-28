package telegram

import "tg-bot/pkg/telegram/payment"

type Update struct {
	UpdateId           int64                     `json:"update_id"`
	Message            *Message                  `json:"message"`
	EditedMessage      *Message                  `json:"edited_message"`
	ChannelPost        *Message                  `json:"channel_post"`
	EditedChannelPost  *Message                  `json:"edited_channel_post"`
	InlineQuery        *InlineQuery              `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult       `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery            `json:"callback_query"`
	ShippingQuery      *payment.ShippingQuery    `json:"shipping_query"`
	PreCheckoutQuery   *payment.PreCheckoutQuery `json:"pre_checkout_query"`
	Poll               *Poll                     `json:"poll"`
	PollAnswer         *PollAnswer               `json:"poll_answer"`
	MyChatMember       *ChatMemberUpdated        `json:"my_chat_member"`
	ChatMember         *ChatMemberUpdated        `json:"chat_member"`
	ChatJoinRequest    *ChatJoinRequest          `json:"chat_join_request"`
}
