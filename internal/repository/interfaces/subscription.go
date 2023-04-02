package interfaces

import "DoramaSet/internal/logic/model"

type ISubscriptionRepo interface {
	GetList() ([]model.Subscription, error)
	GetSubscription(id int) (model.Subscription, error)
}
