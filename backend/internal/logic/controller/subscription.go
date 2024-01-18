package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type SubscriptionController struct {
	repo  repository.ISubscriptionRepo
	urepo repository.IUserRepo
	pc    controller.IPointsController
	uc    controller.IUserController
	log   *logrus.Logger
}

func NewSubscriptionController(SR repository.ISubscriptionRepo, UR repository.IUserRepo,
	pc controller.IPointsController, uc controller.IUserController, log *logrus.Logger) *SubscriptionController {
	return &SubscriptionController{
		repo:  SR,
		urepo: UR,
		pc:    pc,
		uc:    uc,
		log:   log,
	}
}

func (s *SubscriptionController) GetAll(ctx context.Context) ([]model.Subscription, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetAll")
	defer span.End()
	res, err := s.repo.GetList(ctx)
	if err != nil {
		s.log.Warnf("get all subs err %s", err)
		return nil, fmt.Errorf("GetStaffList: %w", err)
	}
	s.log.Infof("got all subs")
	return res, nil
}

func (s *SubscriptionController) GetInfo(ctx context.Context, id int) (*model.Subscription, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetInfo")
	defer span.End()
	res, err := s.repo.GetSubscription(ctx, id)
	if err != nil {
		s.log.Warnf("get info sub err %s, value %d", err, id)
		return nil, fmt.Errorf("GetSubscription: %w", err)
	}
	s.log.Infof("got info one sub value %d", id)
	return res, nil
}

func (s *SubscriptionController) SubscribeUser(ctx context.Context, token string, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL SubscribeUser")
	defer span.End()
	user, err := s.uc.AuthByToken(ctx, token)
	if err != nil {
		s.log.Warnf("subscribe user auth err %s, token %s, value %d", err, token, id)
		return fmt.Errorf("authToken: %w", err)
	}

	sub, err := s.repo.GetSubscription(ctx, id)
	if err != nil {
		s.log.Warnf("subscribe user err %s, user %s, value %d", err, user.Username, id)
		return fmt.Errorf("getSubscription: %w", err)
	}

	err = s.pc.PurgePoint(ctx, user, sub.Cost)
	if err != nil {
		s.log.Warnf("subscribe user err %s, user %s, value %d", err, user.Username, id)
		return fmt.Errorf("purgePoint: %w", err)
	}

	user.Sub = sub
	user.LastSubscribe = time.Now()

	err = s.urepo.UpdateUser(ctx, *user)
	if err != nil {
		s.log.Warnf("subscribe user err %s, user %s, value %d", err, user.Username, id)
		return fmt.Errorf("updateUser: %w", err)
	}
	s.log.Infof("subscribe user %s, id sub %d", user.Username, id)
	return nil
}

func (s *SubscriptionController) UnsubscribeUser(ctx context.Context, token string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL UnsubscribeUser")
	defer span.End()
	user, err := s.uc.AuthByToken(ctx, token)
	if err != nil {
		s.log.Warnf("subscribe user auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}

	sub, err := s.repo.GetSubscriptionByPrice(ctx, 0)
	if err != nil {
		s.log.Warnf("unsubscribe user err %s, user %s", err, user.Username)
		return fmt.Errorf("getSubscriptionByPrice: %w", err)
	}
	user.Sub = sub

	err = s.urepo.UpdateUser(ctx, *user)
	if err != nil {
		s.log.Warnf("unsubscribe user err %s, user %s", err, user.Username)
		return fmt.Errorf("updateUser: %w", err)
	}
	s.log.Infof("unsubsribe user %s", user.Username)
	return nil
}

func (s *SubscriptionController) UpdateSubscribe(ctx context.Context, token string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL UpdateSubscribe")
	defer span.End()
	user, err := s.uc.AuthByToken(ctx, token)
	if err != nil {
		s.log.Warnf("update active user auth err %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	updateDate := user.LastSubscribe.Add(user.Sub.Duration)
	if !eqDate(ctx, time.Now(), updateDate) {
		return nil
	}

	err = s.SubscribeUser(ctx, token, user.Sub.Id)
	if err != nil {
		err := s.UnsubscribeUser(ctx, token)
		if err != nil {
			s.log.Warnf("update subscribe unsubsribe user %s error %s", user.Username, err)
			return fmt.Errorf("unsubscribeUser: %w", err)
		}
		return nil
	}

	s.log.Infof("update subsribe user %s", user.Username)
	return nil
}
