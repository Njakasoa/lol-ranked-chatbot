package models

// Role for each Player
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Roles for the app
type Roles []Role
