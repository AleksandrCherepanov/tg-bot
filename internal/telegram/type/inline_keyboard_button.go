package telegram

type InlineKeyboardButton struct {
	Text                         *string       `json:"text"`
	Url                          *string       `json:"url"`
	CallbackData                 *string       `json:"callback_data"`
	WebApp                       *WebAppInfo   `json:"web_app"`
	LoginUrl                     *LoginUrl     `json:"login_url"`
	SwitchInlineQuery            *string       `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat"`
	CallbackGame                 *CallbackGame `json:"callback_game"`
	Pay                          *bool         `json:"pay"`
}
