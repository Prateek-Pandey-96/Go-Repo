package storage

import "github.com/prateek96/paymentGateway/entities"

type IStorage interface {
	AddClient(client *entities.Client) error
	RemoveClient(client *entities.Client) error
	HasClient(client *entities.Client) (bool, error)
	AddPaymode(paymode entities.Mode) error
	RemovePaymode(paymode entities.Mode) error
	ListPaymodes() ([]entities.Mode, error)
}
