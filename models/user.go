package models

// User of the app
type User struct {
	ID     int    `json:"id"`
	PsID   string `json:"psid"`
	Pseudo string `json:"pseudo"`
}

// Users of the app
type Users []User
