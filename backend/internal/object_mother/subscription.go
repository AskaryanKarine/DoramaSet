package object_mother

import (
	"DoramaSet/internal/logic/model"
	"math/rand"
	"time"
)

type SubscriptionMother struct{}

type subscriptionFunc func(p *model.Subscription)

func SubscriptionWithID(id int) subscriptionFunc {
	return func(p *model.Subscription) {
		p.Id = id
	}
}

func SubscriptionWithName(name string) subscriptionFunc {
	return func(p *model.Subscription) {
		p.Name = name
	}
}

func SubscriptionWithDescriptions(description string) subscriptionFunc {
	return func(p *model.Subscription) {
		p.Description = description
	}
}

func SubscriptionWithCost(cost int) subscriptionFunc {
	return func(p *model.Subscription) {
		p.Cost = cost
	}
}

func SubscriptionWithDuration(t time.Duration) subscriptionFunc {
	return func(p *model.Subscription) {
		p.Duration = t
	}
}

func SubscriptionWithAccessLvl(accessLvl int) subscriptionFunc {
	return func(p *model.Subscription) {
		p.AccessLvl = accessLvl
	}
}

func (e SubscriptionMother) GenerateSubscription(opts ...subscriptionFunc) *model.Subscription {
	p := &model.Subscription{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (e SubscriptionMother) GenerateRandomSubscription() *model.Subscription {
	return &model.Subscription{
		Id:          rand.Int(),
		Name:        randStringBytes(8),
		Description: randStringBytes(8),
		Cost:        rand.Intn(1000),
		Duration:    time.Duration(rand.Intn(100)),
		AccessLvl:   rand.Intn(3),
	}
}

func (e SubscriptionMother) GenerateRandomSubscriptionSlice(size int) []model.Subscription {
	r := make([]model.Subscription, 0)
	for i := 0; i < size; i++ {
		r = append(r, *e.GenerateRandomSubscription())
	}
	return r
}
