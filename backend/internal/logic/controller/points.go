package controller

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type PointsController struct {
	repo             repository.IUserRepo
	everyDayPoint    int
	everyYearPoint   int
	longNoLoginPoint int
	longNoLoginHours float64
	log              *logrus.Logger
}

func NewPointController(URepo repository.IUserRepo, dPoint, YPoint, lPoint int,
	lHours float64, log *logrus.Logger) *PointsController {
	return &PointsController{
		repo:             URepo,
		everyDayPoint:    dPoint,
		everyYearPoint:   YPoint,
		longNoLoginPoint: lPoint,
		longNoLoginHours: lHours,
		log:              log,
	}
}

func checkYear(ctx context.Context, date time.Time) bool {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL checkYear")
	defer span.End()
	today := time.Now()

	if date.Month() == time.February && date.Day() == 29 {
		date = date.Add(-time.Hour * 24)
	}

	if date.Day() != today.Day() {
		return false
	}

	if date.Month() != today.Month() {
		return false
	}

	return true
}

func (p *PointsController) EarnPointForLogin(ctx context.Context, user *model.User) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL EarnPointForLogin")
	defer span.End()
	user.Points += p.everyDayPoint

	if checkYear(ctx, user.RegData) {
		user.Points += p.everyYearPoint
	}

	if time.Since(user.LastActive).Hours() > p.longNoLoginHours {
		user.Points += p.longNoLoginPoint
	}

	err := p.repo.UpdateUser(ctx, *user)
	if err != nil {
		p.log.Warnf("earn point for login err %s username %s", err, user.Username)
		return fmt.Errorf("updateUser: %w", err)
	}
	p.log.Infof("earned point for login user %s", user.Username)
	return nil
}

func (p *PointsController) PurgePoint(ctx context.Context, user *model.User, point int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL PurgePoint")
	defer span.End()
	if user.Points < point {
		err := errors.BalanceError{
			Have: user.Points,
			Want: point,
		}
		p.log.Warnf("purge point err %s, username %s", err, user.Username)
		return fmt.Errorf("purgePoint: %w", err)
	}

	user.Points = user.Points - point

	err := p.repo.UpdateUser(ctx, *user)
	if err != nil {
		p.log.Warnf("purge point err %s, username %s", err, user.Username)
		return fmt.Errorf("updateUser: %w", err)
	}
	p.log.Infof("purged point user %s, value %d", user.Username, point)
	return nil
}

func (p *PointsController) EarnPoint(ctx context.Context, user *model.User, point int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL EarnPoint")
	defer span.End()
	user.Points += point
	err := p.repo.UpdateUser(ctx, *user)
	if err != nil {
		p.log.Warnf("earn point err %s, username %s, value %d", err, user.Username, point)
		return fmt.Errorf("updateUser: %w", err)
	}
	p.log.Infof("earned point user %s, value %d", user.Username, point)
	return nil
}
