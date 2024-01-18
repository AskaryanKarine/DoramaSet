package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type DoramaRepo struct {
	db      *gorm.DB
	picRepo repository.IPictureRepo
	epRepo  repository.IEpisodeRepo
	revRepo repository.IReviewRepo
}

type doramaModel struct {
	ID          int
	Name        string
	Description string
	ReleaseYear int
	Status      string
	Genre       string
}

func NewDoramaRepo(db *gorm.DB, PR repository.IPictureRepo, ER repository.IEpisodeRepo, RR repository.IReviewRepo) *DoramaRepo {
	return &DoramaRepo{db, PR, ER, RR}
}

func (d *DoramaRepo) getDoramaModel(ctx context.Context, m doramaModel) (*model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo getDoramaModel")
	defer span.End()
	ep, err := d.epRepo.GetList(ctx, m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListEp: %w", err)
	}
	photo, err := d.picRepo.GetListDorama(ctx, m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListDorama: %w", err)
	}
	review, err := d.revRepo.GetAllReview(ctx, m.ID)
	if err != nil {
		return nil, fmt.Errorf("getAllReview: %w", err)
	}
	rate, cnt, err := d.revRepo.AggregateRate(ctx, m.ID)
	if err != nil {
		return nil, fmt.Errorf("aggregateRate: %w", err)
	}
	tmp := model.Dorama{
		Id:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Genre:       m.Genre,
		Status:      m.Status,
		ReleaseYear: m.ReleaseYear,
		Episodes:    ep,
		Posters:     photo,
		Reviews:     review,
		Rate:        rate,
		CntRate:     cnt,
	}
	return &tmp, nil
}

func (d *DoramaRepo) GetList(ctx context.Context) ([]model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, nil
	}

	for _, r := range resDB {
		tmp, err := d.getDoramaModel(ctx, r)
		if err != nil {
			return nil, fmt.Errorf("getDoramaModel: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetListName(ctx context.Context, name string) ([]model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListName")
	defer span.End()
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	likeStr := "%" + strings.TrimRight(name, "\r\n") + "%"
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Where("name like ?", likeStr).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		tmp, err := d.getDoramaModel(ctx, r)
		if err != nil {
			return nil, fmt.Errorf("getDoramaModel: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetDorama(ctx context.Context, id int) (*model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetDorama")
	defer span.End()
	var (
		resDB doramaModel
	)
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	tmp, err := d.getDoramaModel(ctx, resDB)
	if err != nil {
		return nil, fmt.Errorf("getDoramaModel: %w", err)
	}

	return tmp, nil
}

func (d *DoramaRepo) CreateDorama(ctx context.Context, dorama model.Dorama) (int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreateDorama")
	defer span.End()
	m := doramaModel{
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (d *DoramaRepo) UpdateDorama(ctx context.Context, dorama model.Dorama) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo UpdateDorama")
	defer span.End()
	m := doramaModel{
		ID:          dorama.Id,
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) DeleteDorama(ctx context.Context, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DeleteDorama")
	defer span.End()
	result := d.db.WithContext(ctx).Table("dorama_set.dorama").Where("id = ?", id).Delete(&doramaModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) AddStaff(ctx context.Context, idD, idS int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetList")
	defer span.End()
	m := struct {
		IdDorama, IdStaff int
	}{idD, idS}
	result := d.db.WithContext(ctx).Table("dorama_set.doramastaff").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) GetListByListId(ctx context.Context, idL int) ([]model.Dorama, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListByListId")
	defer span.End()
	var (
		resDB []struct {
			IdDorama int
			IdList   int
		}
		res []model.Dorama
	)
	result := d.db.WithContext(ctx).Table("dorama_set.listdorama").Where("id_list = ?", idL).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, r := range resDB {
		dorama, err := d.GetDorama(ctx, r.IdDorama)
		if err != nil {
			return nil, fmt.Errorf("getDorama: %w", err)
		}
		res = append(res, *dorama)
	}
	return res, nil
}
