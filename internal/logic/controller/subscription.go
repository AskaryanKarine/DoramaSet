package controller

import (
	"DoramaSet/internal/interfaces"
	"DoramaSet/internal/logic/model"
	"fmt"
)

type SubscriptionController struct {
	repo  interfaces.ISubscriptionRepo
	urepo interfaces.IUserRepo
	pc    interfaces.IPointsController
	uc    interfaces.IUserController
}

func (s *SubscriptionController) GetAll() ([]model.Subscription, error) {
	res, err := s.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("getAll: %w", err)
	}
	return res, nil
}

func (s *SubscriptionController) GetInfo(id int) (*model.Subscription, error) {
	res, err := s.repo.GetSubscription(id)
	if err != nil {
		return nil, fmt.Errorf("GetInfo: %w", err)
	}
	return res, nil
}

func (s *SubscriptionController) SubscribeUser(id int, token string) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	sub, err := s.repo.GetSubscription(id)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	s.pc.PurgePoint(user.Username, sub.Cost)

	user.Sub = sub

	err = s.urepo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	return nil
}

func (s *SubscriptionController) UnsubscribeUser(token string) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	sub, err := s.repo.GetSubscriptionByPrice(0)
	if err != nil {
		return fmt.Errorf("unsubscribe: %w", err)
	}
	user.Sub = sub

	err = s.urepo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	return nil
}
