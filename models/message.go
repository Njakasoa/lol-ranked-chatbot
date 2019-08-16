package models

// QuickReply for messaging
type QuickReply struct {
	Payload string `json:"payload,omitempty"`
}

// Message to analyse
type Message struct {
	MID         string       `json:"mid,omitempty"`
	Text        string       `json:"text,omitempty"`
	QuickReply  *QuickReply  `json:"quick_reply,omitempty"`
	Attachments *Attachments `json:"attachments,omitempty"`
	Attachment  *Attachment  `json:"attachment,omitempty"`
}
