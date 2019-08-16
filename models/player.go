package models

// Player joins Room
type Player struct {
	ID   int  `json:"id"`
	User User `json:"user"`
	Role Role `json:"role"`
}

// Players of the Room
type Players []Player
