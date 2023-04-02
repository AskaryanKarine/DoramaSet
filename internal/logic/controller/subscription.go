package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	repository "DoramaSet/internal/repository/interfaces"
)

type SubscriptionController struct {
	repo  repository.ISubscriptionRepo
	urepo repository.IUserRepo
	pc    interfaces.IPointsController
}

func (s *SubscriptionController) GetAll() ([]model.Subscription, error) {
	return s.repo.GetList()
}

func (s *SubscriptionController) GetInfo(id int) (model.Subscription, error) {
	return s.repo.GetSubscription(id)
}

func (s *SubscriptionController) SubscribeUser(id int, username string) error {
	user, err := s.urepo.GetUser(username)
	if err != nil {
		return err
	}

	sub, err := s.repo.GetSubscription(id)
	if err != nil {
		return err
	}

	if user.IsUsingPoints {
		return s.pc.PurgePoint(username, sub.Cost)
	}

	// тут что-то с оплатой?

	user.Sub = sub

	return s.urepo.UpdateUser(user)
}

func (s *SubscriptionController) UnsubscribeUser(username string) error {
	return nil
}
