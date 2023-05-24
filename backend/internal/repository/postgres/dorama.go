package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
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

func (d *DoramaRepo) getDoramaModel(m doramaModel) (*model.Dorama, error) {
	ep, err := d.epRepo.GetList(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListEp: %w", err)
	}
	photo, err := d.picRepo.GetListDorama(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListDorama: %w", err)
	}
	review, err := d.revRepo.GetAllReview(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getAllReview: %w", err)
	}
	rate, cnt, err := d.revRepo.AggregateRate(m.ID)
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

func (d *DoramaRepo) GetList() ([]model.Dorama, error) {
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	result := d.db.Table("dorama_set.dorama").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, nil
	}

	for _, r := range resDB {
		tmp, err := d.getDoramaModel(r)
		if err != nil {
			return nil, fmt.Errorf("getDoramaModel: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetListName(name string) ([]model.Dorama, error) {
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	likeStr := "%" + strings.TrimRight(name, "\r\n") + "%"
	result := d.db.Table("dorama_set.dorama").Where("name like ?", likeStr).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}

	for _, r := range resDB {
		tmp, err := d.getDoramaModel(r)
		if err != nil {
			return nil, fmt.Errorf("getDoramaModel: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetDorama(id int) (*model.Dorama, error) {
	var (
		resDB doramaModel
	)
	result := d.db.Table("dorama_set.dorama").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	tmp, err := d.getDoramaModel(resDB)
	if err != nil {
		return nil, fmt.Errorf("getDoramaModel: %w", err)
	}

	return tmp, nil
}

func (d *DoramaRepo) CreateDorama(dorama model.Dorama) (int, error) {
	m := doramaModel{
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.Table("dorama_set.dorama").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (d *DoramaRepo) UpdateDorama(dorama model.Dorama) error {
	m := doramaModel{
		ID:          dorama.Id,
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	result := d.db.Table("dorama_set.dorama").Save(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) DeleteDorama(id int) error {
	result := d.db.Table("dorama_set.dorama").Where("id = ?", id).Delete(&doramaModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) AddStaff(idD, idS int) error {
	m := struct {
		IdDorama, IdStaff int
	}{idD, idS}
	result := d.db.Table("dorama_set.doramastaff").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (d *DoramaRepo) GetListByListId(idL int) ([]model.Dorama, error) {
	var (
		resDB []struct {
			IdDorama int
			IdList   int
		}
		res []model.Dorama
	)
	result := d.db.Table("dorama_set.listdorama").Where("id_list = ?", idL).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, r := range resDB {
		dorama, err := d.GetDorama(r.IdDorama)
		if err != nil {
			return nil, fmt.Errorf("getDorama: %w", err)
		}
		res = append(res, *dorama)
	}
	return res, nil
}
