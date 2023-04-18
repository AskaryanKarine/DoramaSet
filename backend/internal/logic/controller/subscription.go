package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"fmt"
	"time"
)

type SubscriptionController struct {
	repo  repository.ISubscriptionRepo
	urepo repository.IUserRepo
	pc    controller.IPointsController
	uc    controller.IUserController
}

func NewSubscriptionController(SR repository.ISubscriptionRepo, UR repository.IUserRepo,
	pc controller.IPointsController, uc controller.IUserController) *SubscriptionController {
	return &SubscriptionController{
		repo:  SR,
		urepo: UR,
		pc:    pc,
		uc:    uc,
	}
}

func (s *SubscriptionController) GetAll() ([]model.Subscription, error) {
	res, err := s.repo.GetList()
	if err != nil {
		return nil, fmt.Errorf("GetList: %w", err)
	}
	return res, nil
}

func (s *SubscriptionController) GetInfo(id int) (*model.Subscription, error) {
	res, err := s.repo.GetSubscription(id)
	if err != nil {
		return nil, fmt.Errorf("GetSubscription: %w", err)
	}
	return res, nil
}

func (s *SubscriptionController) SubscribeUser(token string, id int) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	sub, err := s.repo.GetSubscription(id)
	if err != nil {
		return fmt.Errorf("getSubscription: %w", err)
	}

	err = s.pc.PurgePoint(user, sub.Cost)
	if err != nil {
		return fmt.Errorf("purgePoint: %w", err)
	}

	user.Sub = sub
	user.LastSubscribe = time.Now()

	err = s.urepo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}
	return nil
}

func (s *SubscriptionController) UnsubscribeUser(token string) error {
	user, err := s.uc.AuthByToken(token)
	if err != nil {
		return fmt.Errorf("authToken: %w", err)
	}

	sub, err := s.repo.GetSubscriptionByPrice(0)
	if err != nil {
		return fmt.Errorf("getSubscriptionByPrice: %w", err)
	}
	user.Sub = sub

	err = s.urepo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}
	return nil
}
