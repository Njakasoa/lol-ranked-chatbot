package models

// Entry for callback
type Entry struct {
	ID        string     `json:"id,omitempty"`
	Time      int        `json:"time,omitempty"`
	Messaging Messagings `json:"messaging,omitempty"`
}

// Entries for callback
type Entries []Entry

// Callback for endpoint
type Callback struct {
	Object string  `json:"object,omitempty"`
	Entry  Entries `json:"entry,omitempty"`
}
