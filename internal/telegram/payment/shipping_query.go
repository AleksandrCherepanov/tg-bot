package payment

import "tg-bot/internal/telegram/user"

type ShippingQuery struct {
	Id              string          `json:"id"`
	From            user.User       `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
