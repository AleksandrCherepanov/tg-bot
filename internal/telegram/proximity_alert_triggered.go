package telegram

import "tg-bot/internal/telegram/user"

type ProximityAlertTriggered struct {
	Traveler user.User `json:"traveler"`
	Watcher  user.User `json:"watcher"`
	Distance int64     `json:"distance"`
}
