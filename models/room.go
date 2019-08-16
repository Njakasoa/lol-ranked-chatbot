package models

// Room for the game
type Room struct {
	ID      int     `json:"ID"`
	Owner   User    `json:"owner"`
	Token   string  `json:"token"`
	Players Players `json:"players"`
}

// Rooms for the app
type Rooms []Room
