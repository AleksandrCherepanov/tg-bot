package telegram

type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}
