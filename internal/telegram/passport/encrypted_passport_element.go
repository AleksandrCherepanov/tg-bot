package passport

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        *string        `json:"data"`
	PhoneNumber *string        `json:"phone_number"`
	Email       *string        `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   *PassportFile  `json:"front_side"`
	ReverseSide *PassportFile  `json:"reverse_side"`
	Selfie      *PassportFile  `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Hash        string         `json:"hash"`
}
