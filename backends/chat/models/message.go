package models

import (
	"time"
)

type Message struct {
	User      *User     `json:"user"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func NewMessage(user *User, message string) *Message {
	return &Message{
		User:      user,
		Message:   message,
		Timestamp: time.Now(),
	}
}
