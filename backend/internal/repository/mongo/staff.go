package mongo

import (
	"DoramaSet/internal/interfaces/repository"
	"DoramaSet/internal/logic/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type StaffRepo struct {
	db      *mongo.Database
	picRepo repository.IPictureRepo
}

type staffModel struct {
	ID       int       `bson:"id"`
	Name     string    `bson:"name"`
	Birthday time.Time `bson:"birthday"`
	Gender   string    `bson:"gender"`
	Type     string    `bson:"type"`
}

func (s *StaffRepo) getStaffLogicModel(ctx context.Context, st staffModel) (*model.Staff, error) {
	photo, err := s.picRepo.GetListStaff(ctx, st.ID)
	if err != nil {
		return nil, fmt.Errorf("getListStaff: %w", err)
	}
	return &model.Staff{
		Id:       st.ID,
		Name:     st.Name,
		Birthday: st.Birthday,
		Type:     st.Type,
		Gender:   st.Gender,
		Photo:    photo,
	}, nil
}

func NewStaffRepo(db *mongo.Database, pr repository.IPictureRepo) *StaffRepo {
	return &StaffRepo{db: db, picRepo: pr}
}

func (s *StaffRepo) GetList(ctx context.Context) ([]model.Staff, error) {
	var (
		resDB []staffModel
		res   []model.Staff
	)

	collection := s.db.Collection("staff")
	filter := bson.D{{}}
	opts := options.Find().SetSort(bson.D{{"id", 1}})

	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db_find: %w", err)
	}

	if err := cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	for _, r := range resDB {
		tmp, err := s.getStaffLogicModel(ctx, r)
		if err != nil {
			return nil, err
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (s *StaffRepo) GetListName(ctx context.Context, name string) ([]model.Staff, error) {
	var (
		resDB []staffModel
		res   []model.Staff
	)
	collection := s.db.Collection("staff")
	filter := bson.D{{"name", bson.D{{"$regex", strings.TrimRight(name, "\r\n")}}}}
	opts := options.Find().SetSort(bson.D{{"id", 1}})

	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db_find: %w", err)
	}

	if err := cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	for _, r := range resDB {
		tmp, err := s.getStaffLogicModel(ctx, r)
		if err != nil {
			return nil, err
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (s *StaffRepo) GetListDorama(ctx context.Context, idDorama int) ([]model.Staff, error) {
	var (
		resDB []struct {
			IdDorama int `bson:"id_dorama"`
			IdStaff  int `bson:"id_staff"`
		}
		res []model.Staff
	)
	collection := s.db.Collection("_dorama_staff")
	filter := bson.D{{"id_dorama", idDorama}}
	opts := options.Find().SetSort(bson.D{{"id_staff", 1}})

	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("db_find: %w", err)
	}

	if err := cur.All(nil, &resDB); err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	for _, r := range resDB {
		tmp, err := s.GetStaffById(ctx, r.IdStaff)
		if err != nil {
			return nil, fmt.Errorf("getStaffById: %w", err)
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (s *StaffRepo) CreateStaff(ctx context.Context, record model.Staff) (int, error) {
	var (
		maxID staffModel
	)
	collection := s.db.Collection("staff")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	err := collection.FindOne(nil, filter, opts).Decode(&maxID)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	newStaff := staffModel{
		ID:       maxID.ID + 1,
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Type:     record.Type,
	}
	_, err = collection.InsertOne(nil, newStaff)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	return newStaff.ID, nil
}

func (s *StaffRepo) UpdateStaff(ctx context.Context, record model.Staff) error {
	m := staffModel{
		ID:       record.Id,
		Name:     record.Name,
		Birthday: record.Birthday,
		Gender:   record.Gender,
		Type:     record.Type,
	}
	collection := s.db.Collection("staff")
	filter := bson.D{{"id", record.Id}}
	_, err := collection.ReplaceOne(nil, filter, m)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (s *StaffRepo) DeleteStaff(ctx context.Context, id int) error {
	collection := s.db.Collection("staff")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (s *StaffRepo) GetStaffById(ctx context.Context, id int) (*model.Staff, error) {
	var (
		resDB staffModel
	)
	collection := s.db.Collection("staff")
	filter := bson.D{{"id", id}}

	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db_find: %w", err)
	}

	res, err := s.getStaffLogicModel(ctx, resDB)
	if err != nil {
		return nil, err
	}

	return res, nil
}
