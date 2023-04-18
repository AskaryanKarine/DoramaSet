package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"fmt"
	"gorm.io/gorm"
)

type ListRepo struct {
	db         *gorm.DB
	doramaRepo repository.IDoramaRepo
}

type listModel struct {
	ID          int
	NameCreator string
	NameList    string
	Type        string
	Description string
}

func NewListRepo(db *gorm.DB, DR repository.IDoramaRepo) *ListRepo {
	return &ListRepo{db, DR}
}

func (l *ListRepo) GetUserLists(username string) ([]model.List, error) {
	var (
		resDB []listModel
		res   []model.List
	)
	result := l.db.Table("dorama_set.list l").Select("l.*").
		Joins("join dorama_set.userlist ul on l.id = ul.id_list").
		Where("ul.username = ?", username).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if len(resDB) == 0 {
		return nil, nil
	}

	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        r.Type,
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (l *ListRepo) GetPublicLists() ([]model.List, error) {
	var (
		resDB []listModel
		res   []model.List
	)
	result := l.db.Table("dorama_set.list").Where("type = public").Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if len(resDB) == 0 {
		return nil, fmt.Errorf("db: %w", errors.ErrorDontExistsInDB)
	}
	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        r.Type,
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (l *ListRepo) GetListId(id int) (*model.List, error) {
	var (
		resDB listModel
		res   model.List
	)
	result := l.db.Table("dorama_set.list").Where("id = ?", id).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	dorama, err := l.doramaRepo.GetListByListId(resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListByList: %w", err)
	}
	res = model.List{
		Id:          resDB.ID,
		Name:        resDB.NameList,
		Description: resDB.Description,
		CreatorName: resDB.NameCreator,
		Type:        resDB.Type,
		Doramas:     dorama,
	}
	return &res, nil
}

func (l *ListRepo) CreateList(list model.List) (int, error) {
	m := listModel{
		NameCreator: list.CreatorName,
		NameList:    list.Name,
		Type:        list.Type,
		Description: list.Description,
	}
	result := l.db.Table("dorama_set.list").Create(&m)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	m1 := struct {
		Username string
		IdList   int
	}{list.CreatorName, m.ID}
	result = l.db.Table("dorama_set.userlist").Create(&m1)
	if result.Error != nil {
		return -1, fmt.Errorf("db: %w", result.Error)
	}
	return m.ID, nil
}

func (l *ListRepo) DelList(id int) error {
	result := l.db.Table("dorama_set.list").Where("id = ?", id).Delete(&listModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) AddToList(idL, idD int) error {
	m := struct {
		IdDorama, IdList int
	}{idD, idL}
	result := l.db.Table("dorama_set.listdorama").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) DelFromList(idL, idD int) error {
	result := l.db.Table("dorama_set.listdorama").Where("id_list = ? and id_dorama = ?", idL, idD).Delete(&struct {
		IdDorama, IdList int
	}{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) AddToFav(idL int, username string) error {
	m := struct {
		Username string
		IdList   int
	}{username, idL}
	result := l.db.Table("dorama_set.userlist").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) GetFavList(username string) ([]model.List, error) {
	var (
		res   []model.List
		resDB []listModel
	)
	result := l.db.Table("dorama_set.list l").Select("l.*").
		Joins("join dorama_set.list l on l.id = ul.id_list").
		Where("ul.username = ? and ul.username != l.name_creator", username).
		Find(&resDB)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        r.Type,
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}
