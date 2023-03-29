package interfaces

import "DoramaSet/internal/logic/model"

type ISubscriptionController interface {
	getAll() ([]model.Subscription, error)
	getInfo(record model.Subscription) (model.Subscription, error)
	subscribeUser(record model.Subscription, idUser int) error
	unsubscribeUser(record model.Subscription, idUser int) error
}

type ISubscriptionRepo interface {
}
