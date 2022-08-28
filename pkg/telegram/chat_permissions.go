package telegram

type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages"`
	CanSendMediaMessages  *bool `json:"can_send_media_messages"`
	CanSendPolls          *bool `json:"can_send_polls"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews"`
	CanChangeInfo         *bool `json:"can_change_info"`
	CanInvite_users       *bool `json:"can_invite___users"`
	CanPinMessages        *bool `json:"can_pin_messages"`
}
