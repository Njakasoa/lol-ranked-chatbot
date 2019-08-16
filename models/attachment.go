package models

// Attachment for messaging
type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

//Attachments for messaging
type Attachments []Attachment
