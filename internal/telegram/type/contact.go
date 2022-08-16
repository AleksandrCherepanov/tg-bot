package telegram

type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	UserId      *int64  `json:"user_id"`
	Vcard       *string `json:"vcard"`
}
