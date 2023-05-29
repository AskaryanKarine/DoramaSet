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
)

type DoramaRepo struct {
	db      *mongo.Database
	picRepo repository.IPictureRepo
	epRepo  repository.IEpisodeRepo
	revRepo repository.IReviewRepo
}

type doramaModel struct {
	ID          int    `bson:"id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	ReleaseYear int    `bson:"release_year"`
	Status      string `bson:"status"`
	Genre       string `bson:"genre"`
}

func NewDoramaRepo(db *mongo.Database, PR repository.IPictureRepo, ER repository.IEpisodeRepo, RR repository.IReviewRepo) *DoramaRepo {
	return &DoramaRepo{db, PR, ER, RR}
}

func (d *DoramaRepo) getDoramaLogicModel(m doramaModel) (*model.Dorama, error) {
	ep, err := d.epRepo.GetList(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListEp: %w", err)
	}
	photo, err := d.picRepo.GetListDorama(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getListDoramaPic: %w", err)
	}
	review, err := d.revRepo.GetAllReview(m.ID)
	if err != nil {
		return nil, fmt.Errorf("getAllReview: %w", err)
	}
	rate, cnt, err := d.revRepo.AggregateRate(m.ID)
	if err != nil {
		return nil, fmt.Errorf("aggreagateRate: %w", err)
	}
	tmp := model.Dorama{
		Id:          m.ID,
		Name:        m.Name,
		Description: strings.Replace(m.Description, ";", ",", -1),
		Genre:       m.Genre,
		Status:      m.Status,
		ReleaseYear: m.ReleaseYear,
		Rate:        rate,
		CntRate:     cnt,
		Posters:     photo,
		Episodes:    ep,
		Reviews:     review,
	}

	return &tmp, nil
}

func (d *DoramaRepo) GetList() ([]model.Dorama, error) {
	var (
		resDB []doramaModel
		res   []model.Dorama
	)
	collection := d.db.Collection("dorama")
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
		tmp, err := d.getDoramaLogicModel(r)
		if err != nil {
			return nil, err
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
	collection := d.db.Collection("dorama")
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
		tmp, err := d.getDoramaLogicModel(r)
		if err != nil {
			return nil, err
		}
		res = append(res, *tmp)
	}
	return res, nil
}

func (d *DoramaRepo) GetDorama(id int) (*model.Dorama, error) {
	var (
		redDB doramaModel
	)
	collection := d.db.Collection("dorama")
	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&redDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	res, err := d.getDoramaLogicModel(redDB)
	if err != nil {
		return nil, fmt.Errorf("getDoramaLogicModel: %w", err)
	}

	return res, nil
}

func (d *DoramaRepo) CreateDorama(dorama model.Dorama) (int, error) {
	var (
		maxID doramaModel
	)
	collection := d.db.Collection("dorama")
	filter := bson.D{{}}
	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	err := collection.FindOne(nil, filter, opts).Decode(&maxID)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	newStaff := doramaModel{
		ID:          maxID.ID + 1,
		Name:        dorama.Name,
		Description: dorama.Description,
		ReleaseYear: dorama.ReleaseYear,
		Status:      dorama.Status,
		Genre:       dorama.Genre,
	}
	_, err = collection.InsertOne(nil, newStaff)
	if err != nil {
		return -1, fmt.Errorf("db: %w", err)
	}
	return newStaff.ID, nil
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
	collection := d.db.Collection("staff")
	filter := bson.D{{"id", dorama.Id}}
	_, err := collection.ReplaceOne(nil, filter, m)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (d *DoramaRepo) DeleteDorama(id int) error {
	collection := d.db.Collection("dorama")
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (d *DoramaRepo) AddStaff(idD, idS int) error {
	type query struct {
		IdDorama int `bson:"id_dorama"`
		IdStaff  int `bson:"id_staff"`
	}
	var resDB query

	collection := d.db.Collection("_dorama_staff")
	filter := bson.D{{"id_dorama", idD}, {"id_staff", idS}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf("db: %w", err)
	}
	tmp := query{
		IdDorama: idD,
		IdStaff:  idS,
	}
	_, err = collection.InsertOne(nil, tmp)
	if err != nil {
		return fmt.Errorf("db: %w", err)
	}
	return nil
}

func (d *DoramaRepo) GetListByListId(idL int) ([]model.Dorama, error) {
	var (
		resDB struct {
			Dorama []int `bson:"dorama"`
		}
		res []model.Dorama
	)
	collection := d.db.Collection("list")
	filter := bson.D{{"id", idL}}
	err := collection.FindOne(nil, filter).Decode(&resDB)
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	for _, r := range resDB.Dorama {
		dorama, err := d.GetDorama(r)
		if err != nil {
			return nil, fmt.Errorf("getDoramaById: %w", err)
		}
		res = append(res, *dorama)
	}
	return res, nil
}
