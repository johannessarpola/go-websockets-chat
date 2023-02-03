package models

import (
	"github.com/google/uuid"
)

type Room struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewRoom(name string) *Room {
	return &Room{
		Id:   uuid.New().String(),
		Name: name,
	}
}
