package telegram

import "tg-bot/pkg/telegram/user"

type ChatInviteLink struct {
	InviteLink              string    `json:"invite_link"`
	Creator                 user.User `json:"creator"`
	CreatesJoinRequest      bool      `json:"creates_join_request"`
	IsPrimary               bool      `json:"is_primary"`
	IsRevoked               bool      `json:"is_revoked"`
	Name                    *string   `json:"name"`
	ExpireDate              *int64    `json:"expire_date"`
	MemberLimit             *int64    `json:"member_limit"`
	PendingJoinRequestCount *int64    `json:"pending_join_request_count"`
}
