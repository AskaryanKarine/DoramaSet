package controller

import (
	"DoramaSet/internal/interfaces/controller"
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DoramaController struct {
	repo    repository.IDoramaRepo
	revRepo repository.IReviewRepo
	uc      controller.IUserController
	log     *logrus.Logger
}

func NewDoramaController(DRepo repository.IDoramaRepo, RRepo repository.IReviewRepo,
	uc controller.IUserController, log *logrus.Logger) *DoramaController {
	return &DoramaController{repo: DRepo, uc: uc, log: log, revRepo: RRepo}
}

func (d *DoramaController) GetAllDorama(ctx context.Context) ([]model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetAllDorama")
	defer span.End()
	res, err := d.repo.GetList(ctx)
	if err != nil {
		d.log.Warnf("get all dorama, get list err %s", err)
		return nil, fmt.Errorf("getList: %w", err)
	}
	d.log.Infof("got all dorama")
	return res, nil
}

func (d *DoramaController) GetDoramaByName(ctx context.Context, name string) ([]model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetDoramaByName")
	defer span.End()
	res, err := d.repo.GetListName(ctx, name)
	if err != nil {
		d.log.Warnf("get dorama by name, get list name error: %s, value: %s", err, name)
		return nil, fmt.Errorf("getListName: %w", err)
	}
	d.log.Infof("got dorama by name, value %s", name)
	return res, nil
}

func (d *DoramaController) GetDoramaById(ctx context.Context, id int) (*model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL GetDoramaById")
	defer span.End()
	res, err := d.repo.GetDorama(ctx, id)
	if err != nil {
		d.log.Warnf("get dorama by id error: %s, value: %d", err, id)
		return nil, fmt.Errorf("getDorama: %w", err)
	}
	d.log.Infof("got dorama by id, value %d", id)
	return res, nil
}

func (d *DoramaController) CreateDorama(ctx context.Context, token string, record *model.Dorama) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL CreateDorama")
	defer span.End()
	user, err := d.uc.AuthByToken(ctx, token)
	if err != nil {
		d.log.Warnf("create dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}

	if !user.IsAdmin {
		d.log.Warnf("create dorama, access err, username %s", user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	id, err := d.repo.CreateDorama(ctx, *record)
	if err != nil {
		d.log.Warnf("create dorama, err %s, username %s, value %v", err, user.Username, record)
		return fmt.Errorf("createDorama: %w", err)
	}
	record.Id = id
	d.log.Infof("created dorama, username %s, record %v", user.Username, record)
	return nil
}

func (d *DoramaController) UpdateDorama(ctx context.Context, token string, record model.Dorama) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL UpdateDorama")
	defer span.End()
	user, err := d.uc.AuthByToken(ctx, token)
	if err != nil {
		d.log.Warnf("update dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !user.IsAdmin {
		d.log.Warnf("update dorama, access err: %s username %s", err, user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}

	err = d.repo.UpdateDorama(ctx, record)
	if err != nil {
		d.log.Warnf("update dorama, update err: %s username %s, record %v", err, user.Username, record)
		return fmt.Errorf("updateDorama: %w", err)
	}
	d.log.Infof("updated dorama, username %s, record %v", user.Username, record)
	return nil
}

func (d *DoramaController) AddStaffToDorama(ctx context.Context, token string, idD, idS int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL AddStaffToDorama")
	defer span.End()
	user, err := d.uc.AuthByToken(ctx, token)
	if err != nil {
		d.log.Warnf("update dorama, auth err: %s, token %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	if !user.IsAdmin {
		d.log.Warnf("update dorama, access err: %s username %s", err, user.Username)
		return fmt.Errorf("%w", errors.ErrorAdminAccess)
	}
	err = d.repo.AddStaff(ctx, idD, idS)
	if err != nil {
		d.log.Warnf("add staff to dorama, add err: %s values %d %d", err, idD, idS)
		return fmt.Errorf("addStaff: %w", err)
	}
	d.log.Infof("added staff to dorama, value %d %d", idD, idS)
	return nil
}

func (d *DoramaController) AddReview(ctx context.Context, token string, idD int, review *model.Review) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL AddReview")
	defer span.End()
	user, err := d.uc.AuthByToken(ctx, token)
	if err != nil {
		d.log.Warnf("add review, auth err: %s, token: %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	review.Username = user.Username
	err = d.revRepo.CreateReview(ctx, idD, review)
	if err != nil {
		d.log.Warnf("add review, create err: %s, value %d %v", err, idD, review)
		return fmt.Errorf("createReview: %w", err)
	}
	d.log.Infof("added new review into dorama %d by user %s", idD, user.Username)
	return nil
}

func (d *DoramaController) DeleteReview(ctx context.Context, token string, idD int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "BL DeleteReview")
	defer span.End()
	user, err := d.uc.AuthByToken(ctx, token)
	if err != nil {
		d.log.Warnf("delete review, auth err: %s, token: %s", err, token)
		return fmt.Errorf("authToken: %w", err)
	}
	err = d.revRepo.DeleteReview(ctx, user.Username, idD)
	if err != nil {
		d.log.Warnf("delete review, delete err: %s, user %s, id %d", err, user.Username, idD)
		return fmt.Errorf("deleteReview: %w", err)
	}
	d.log.Infof("deleted review from dorama %d by user %s", idD, user.Username)
	return nil
}
