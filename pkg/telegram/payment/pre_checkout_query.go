package payment

import "tg-bot/pkg/telegram/user"

type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             user.User  `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int64      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId *string    `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}
