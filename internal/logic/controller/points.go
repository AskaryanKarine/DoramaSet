package controller

import (
	"DoramaSet/internal/interfaces"
	"fmt"
	"time"
)

type PointsController struct {
	repo interfaces.IUserRepo
}

func checkYear(date time.Time) bool {
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

func (p *PointsController) EarnPointForLogin(username string) error {
	user, err := p.repo.GetUser(username)
	if err != nil {
		return fmt.Errorf("earnPointLogin: %w", err)
	}
	user.Points += 5

	if checkYear(user.RegData) {
		user.Points += 10
	}

	if time.Since(user.LastActive).Hours() > 4400.0 {
		user.Points += 50
	}

	err = p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("earnPointLogin: %w", err)
	}
	return nil
}

func (p *PointsController) PurgePoint(username string, point int) error {
	user, err := p.repo.GetUser(username)
	if err != nil {
		return fmt.Errorf("purgePoint: %w", err)
	}

	if user.Points < point {
		return fmt.Errorf("purgePoint: not enough point: you have %d, but need: %d", user.Points, point)
	}

	user.Points -= point

	err = p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("purgePoint: %w", err)
	}
	return nil
}

func (p *PointsController) EarnPoint(username string, point int) error {
	user, err := p.repo.GetUser(username)
	if err != nil {
		return fmt.Errorf("earnPoint: %w", err)
	}
	user.Points += point
	err = p.repo.UpdateUser(*user)
	if err != nil {
		return fmt.Errorf("earnPoint: %w", err)
	}
	return nil
}
