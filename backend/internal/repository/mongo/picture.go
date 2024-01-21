package mongo

import (
	"DoramaSet/internal/logic/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PictureRepo struct {
	db *mongo.Database
}

func NewPictureRepo(db *mongo.Database) *PictureRepo {
	return &PictureRepo{db}
}

type pictureModel struct {
	ID  int    `bson:"id"`
	URL string `bson:"url"`
}

func (p *PictureRepo) getPictureById(ctx context.Context, id int) (*model.Picture, error) {
	var tmp pictureModel
	collection := p.db.Collection("_picture")
	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&tmp)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	res := model.Picture{
		Id:  tmp.ID,
		URL: tmp.URL,
	}
	return &res, nil
}

func (p *PictureRepo) GetListDorama(ctx context.Context, idDorama int) ([]model.Picture, error) {
	var (
		res   []model.Picture
		resDB []struct {
			IDDorama  int `bson:"id_dorama"`
			IDPicture int `bson:"id_picture"`
		}
	)
	helpCollection := p.db.Collection("_dorama_picture")
	helpFilter := bson.D{{"id_dorama", idDorama}}
	helpCur, err := helpCollection.Find(context.TODO(), helpFilter)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = helpCur.All(context.TODO(), &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		tmp, err := p.getPictureById(ctx, r.IDPicture)
		if err != nil {
			return nil, fmt.Errorf("getPictureById: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (p *PictureRepo) GetListStaff(ctx context.Context, idStaff int) ([]model.Picture, error) {
	var (
		res   []model.Picture
		resDB []struct {
			IDStaff   int `bson:"id_staff"`
			IDPicture int `bson:"id_picture"`
		}
	)
	helpCollection := p.db.Collection("_staff_picture")
	helpFilter := bson.D{{"id_staff", idStaff}}
	helpCur, err := helpCollection.Find(context.TODO(), helpFilter)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	if err = helpCur.All(context.TODO(), &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}
	for _, r := range resDB {
		tmp, err := p.getPictureById(ctx, r.IDPicture)
		if err != nil {
			return nil, fmt.Errorf("getPictureById: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (p *PictureRepo) CreatePicture(ctx context.Context, record model.Picture) (int, error) {
	var maxID pictureModel

	collection := p.db.Collection("_picture")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	err := collection.FindOne(context.TODO(), filter, opts).Decode(&maxID)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	newPicture := pictureModel{
		ID:  maxID.ID + 1,
		URL: record.URL,
	}
	_, err = collection.InsertOne(context.TODO(), newPicture)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	return newPicture.ID, nil
}

func (p *PictureRepo) AddPictureToStaff(ctx context.Context, record model.Picture, id int) error {
	type query struct {
		IDStaff   int `bson:"id_staff"`
		IDPicture int `bson:"id_picture"`
	}
	var tmp query
	collection := p.db.Collection("_staff_picture")
	filter := bson.D{{"id_staff", id}, {"id_picture", record.Id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&tmp)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	tmp = query{id, record.Id}
	_, err = collection.InsertOne(context.TODO(), tmp)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (p *PictureRepo) AddPictureToDorama(ctx context.Context, record model.Picture, id int) error {
	type query struct {
		IDDorama  int `bson:"id_dorama"`
		IDPicture int `bson:"id_picture"`
	}
	var tmp query
	collection := p.db.Collection("_dorama_picture")
	filter := bson.D{{"id_staff", id}, {"id_picture", record.Id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&tmp)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	tmp = query{id, record.Id}
	_, err = collection.InsertOne(context.TODO(), tmp)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (p *PictureRepo) DeletePicture(ctx context.Context, id int) error {
	collection := p.db.Collection("_picture")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}
