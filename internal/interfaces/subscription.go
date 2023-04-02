package interfaces

import "DoramaSet/internal/logic/model"

type ISubscriptionController interface {
	GetAll() ([]model.Subscription, error)
	GetInfo(id int) (model.Subscription, error)
	SubscribeUser(id int, username string) error
	UnsubscribeUser(username string) error
}
