package postgres

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	"DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
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

func (l *ListRepo) GetUserLists(ctx context.Context, username string) ([]model.List, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetUserLists")
	defer span.End()
	var (
		resDB []listModel
		res   []model.List
	)
	result := l.db.WithContext(ctx).Table("dorama_set.list l").Select("l.*").
		Joins("join dorama_set.userlist ul on l.id = ul.id_list").
		Where("ul.username = ? and ul.username = l.name_creator", username).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	if len(resDB) == 0 {
		return nil, nil
	}

	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        constant.ListType[r.Type],
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (l *ListRepo) GetPublicLists(ctx context.Context) ([]model.List, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetPublicLists")
	defer span.End()
	var (
		resDB []listModel
		res   []model.List
	)
	listType, _ := constant.GetTypeList(constant.PublicList)
	result := l.db.WithContext(ctx).Table("dorama_set.list l").Where("l.type = ?", listType).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        constant.ListType[r.Type],
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}

func (l *ListRepo) GetListId(ctx context.Context, id int) (*model.List, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetListId")
	defer span.End()
	var (
		resDB listModel
		res   model.List
	)
	result := l.db.WithContext(ctx).Table("dorama_set.list").Where("id = ?", id).Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	dorama, err := l.doramaRepo.GetListByListId(ctx, resDB.ID)
	if err != nil {
		return nil, fmt.Errorf("getListByList: %w", err)
	}
	res = model.List{
		Id:          resDB.ID,
		Name:        resDB.NameList,
		Description: resDB.Description,
		CreatorName: resDB.NameCreator,
		Type:        constant.ListType[resDB.Type],
		Doramas:     dorama,
	}
	return &res, nil
}

func (l *ListRepo) CreateList(ctx context.Context, list model.List) (int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreateList")
	defer span.End()
	key, err := constant.GetTypeList(list.Type)
	if err != nil {
		return -1, fmt.Errorf("getTypeList: %w", err)
	}
	m := listModel{
		NameCreator: list.CreatorName,
		NameList:    list.Name,
		Type:        key,
		Description: list.Description,
	}
	result := l.db.WithContext(ctx).Table("dorama_set.list").Create(&m)
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

func (l *ListRepo) DelList(ctx context.Context, id int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DelList")
	defer span.End()
	result := l.db.WithContext(ctx).Table("dorama_set.list").Where("id = ?", id).Delete(&listModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) AddToList(ctx context.Context, idL, idD int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo AddToList")
	defer span.End()
	m := struct {
		IdDorama, IdList int
	}{idD, idL}
	result := l.db.WithContext(ctx).Table("dorama_set.listdorama").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) DelFromList(ctx context.Context, idL, idD int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DelFromList")
	defer span.End()
	result := l.db.WithContext(ctx).Table("dorama_set.listdorama").Where("id_list = ? and id_dorama = ?", idL, idD).Delete(&struct {
		IdDorama, IdList int
	}{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) AddToFav(ctx context.Context, idL int, username string) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo AddToFav")
	defer span.End()
	m := struct {
		Username string
		IdList   int
	}{username, idL}
	list, err := l.GetListId(ctx, idL)
	if err != nil {
		return fmt.Errorf("degListId: %w", err)
	}
	if list.Type != constant.PublicList {
		return fmt.Errorf("db: %w", errors.ErrorPublic)
	}
	result := l.db.WithContext(ctx).Table("dorama_set.userlist").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (l *ListRepo) GetFavList(ctx context.Context, username string) ([]model.List, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetFavList")
	defer span.End()
	var (
		res   []model.List
		resDB []listModel
	)
	result := l.db.Table("dorama_set.list l").Select("l.*").
		Joins("join dorama_set.userlist ul on l.id = ul.id_list").
		Where("ul.username = ? and ul.username != l.name_creator", username).
		Find(&resDB)

	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	for _, r := range resDB {
		dorama, err := l.doramaRepo.GetListByListId(ctx, r.ID)
		if err != nil {
			return nil, fmt.Errorf("getListByList: %w", err)
		}
		tmp := model.List{
			Id:          r.ID,
			Name:        r.NameList,
			Description: r.Description,
			CreatorName: r.NameCreator,
			Type:        constant.ListType[r.Type],
			Doramas:     dorama,
		}
		res = append(res, tmp)
	}
	return res, nil
}
