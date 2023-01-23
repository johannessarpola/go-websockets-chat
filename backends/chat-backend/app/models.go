package app

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{
		Id:   uuid.New().String(),
		Name: name,
	}
}

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
