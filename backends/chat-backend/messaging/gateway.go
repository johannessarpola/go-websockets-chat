package messaging

import "github.com/johannessarpola/go-websockets-chat/models"

type Gateway interface {
	Run()
	Send(models.Message) error
	Poll() ([]models.Message, error)
}
