package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/constant"
	errors2 "DoramaSet/internal/logic/errors"
	"DoramaSet/internal/logic/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListRepo struct {
	db         *mongo.Database
	doramaRepo repository.IDoramaRepo
}

func NewListRepo(db *mongo.Database, DR repository.IDoramaRepo) *ListRepo {
	return &ListRepo{db, DR}
}

type listModel struct {
	ID          int    `bson:"id"`
	NameCreator string `bson:"name_creator"`
	NameList    string `bson:"name_list"`
	Type        string `bson:"type"`
	Description string `bson:"description"`
}

func (l *ListRepo) getListLogicModel(ctx context.Context, list listModel) (*model.List, error) {
	dorama, err := l.doramaRepo.GetListByListId(ctx, list.ID)
	if err != nil {
		return nil, fmt.Errorf("getDoramaListByListID: %w", err)
	}
	return &model.List{
		Id:          list.ID,
		Name:        list.NameList,
		Description: list.Description,
		CreatorName: list.NameCreator,
		Type:        constant.ListType[list.Type],
		Doramas:     dorama,
	}, nil
}

func (l *ListRepo) GetUserLists(ctx context.Context, username string) ([]model.List, error) {
	var (
		resDB []listModel
		res   []model.List
	)
	collection := l.db.Collection("list")
	filter := bson.D{{"name_creator", username}}
	opts := options.Find().SetSort(bson.D{{"id", 1}})
	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		tmp, err := l.getListLogicModel(ctx, r)
		if err != nil {
			return nil, err
		}
		res = append(res, *tmp)
	}
	return res, err
}

func (l *ListRepo) GetPublicLists(ctx context.Context) ([]model.List, error) {
	var (
		resDB []listModel
		res   []model.List
	)
	collection := l.db.Collection("list")
	listType, _ := constant.GetTypeList(constant.PublicList)
	filter := bson.D{{"type", listType}}
	opts := options.Find().SetSort(bson.D{{"id", 1}})
	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		tmp, err := l.getListLogicModel(ctx, r)
		if err != nil {
			return nil, err
		}
		res = append(res, *tmp)
	}
	return res, err
}

func (l *ListRepo) GetListId(ctx context.Context, id int) (*model.List, error) {
	var (
		resDB listModel
	)
	collection := l.db.Collection("list")
	filter := bson.D{{"id", id}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	res, err := l.getListLogicModel(ctx, resDB)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *ListRepo) CreateList(ctx context.Context, list model.List) (int, error) {
	var (
		maxIDEp listModel
	)
	collection := l.db.Collection("list")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	err := collection.FindOne(nil, filter, opts).Decode(&maxIDEp)
	if err != nil && err != mongo.ErrNoDocuments {
		return -1, fmt.Errorf("db: %w", err)
	}
	if err == mongo.ErrNoDocuments {
		maxIDEp.ID = 0
	}
	key, err := constant.GetTypeList(list.Type)
	if err != nil {
		return -1, fmt.Errorf("getTypeList: %w", err)
	}
	newList := listModel{
		ID:          maxIDEp.ID + 1,
		NameCreator: list.CreatorName,
		NameList:    list.Name,
		Type:        key,
		Description: list.Description,
	}
	_, err = collection.InsertOne(nil, newList)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	return newList.ID, nil
}

func (l *ListRepo) DelList(ctx context.Context, id int) error {
	collection := l.db.Collection("list")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (l *ListRepo) AddToList(ctx context.Context, idL, idD int) error {
	var resDB struct {
		Dorama []int `bson:"dorama"`
	}
	collection := l.db.Collection("list")
	helpFilter := bson.D{{"id", idL}, {"dorama", idD}}
	err := collection.FindOne(nil, helpFilter).Decode(&resDB)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	if len(resDB.Dorama) > 0 {
		return errors2.ErrorExistInDB
	}
	filter := bson.D{{"id", idL}}
	update := bson.D{{"$push", bson.D{{"dorama", idD}}}}
	_, err = collection.UpdateOne(nil, filter, update)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (l *ListRepo) DelFromList(ctx context.Context, idL, idD int) error {
	collection := l.db.Collection("list")
	filter := bson.D{{"id", idL}}
	update := bson.D{{"$pull", bson.D{{"dorama", idD}}}}
	_, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (l *ListRepo) AddToFav(ctx context.Context, idL int, username string) error {
	type query struct {
		Username string `bson:"username"`
		Episode  int    `bson:"list"`
	}
	var resDB query

	collection := l.db.Collection("_user_fav")
	helpFilter := bson.D{{"username", username}, {"list", idL}}
	err := collection.FindOne(nil, helpFilter).Decode(&resDB)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	if err == nil {
		return errors2.ErrorExistInDB
	}
	_, err = collection.InsertOne(nil, query{
		Username: username,
		Episode:  idL,
	})
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (l *ListRepo) GetFavList(ctx context.Context, username string) ([]model.List, error) {
	type query struct {
		Username string `bson:"username"`
		List     int    `bson:"list"`
	}
	var resDB []query
	var res []model.List

	collection := l.db.Collection("_user_fav")
	helpFilter := bson.D{{"username", username}}
	cur, err := collection.Find(nil, helpFilter)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		list, err := l.GetListId(ctx, r.List)
		if err != nil {
			return nil, err
		}
		res = append(res, *list)
	}

	return res, nil
}
