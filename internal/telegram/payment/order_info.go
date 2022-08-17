package payment

type OrderInfo struct {
	Name            *string          `json:"name"`
	PhoneNumber     *string          `json:"phone_number"`
	Email           *string          `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}
