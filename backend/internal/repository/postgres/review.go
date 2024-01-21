package postgres

import (
	"DoramaSet/internal/logic/model"
	"DoramaSet/internal/tracing"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type ReviewRepo struct {
	db *gorm.DB
}

type reviewModel struct {
	IdDorama int
	Username string
	Mark     int
	Content  string
}

func NewReviewRepo(db *gorm.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) GetAllReview(ctx context.Context, idD int) ([]model.Review, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetAllReview")
	defer span.End()
	var (
		resDB []reviewModel
		res   []model.Review
	)
	result := r.db.WithContext(ctx).Table("dorama_set.review").
		Where("id_dorama = ?", idD).Find(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}

	for _, d := range resDB {
		m := model.Review{
			Username: d.Username,
			Mark:     d.Mark,
			Content:  d.Content,
		}
		res = append(res, m)
	}
	return res, nil
}

func (r *ReviewRepo) GetReview(ctx context.Context, username string, idD int) (*model.Review, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo GetReview")
	defer span.End()
	var resDB reviewModel
	result := r.db.WithContext(ctx).Table("dorama_set.review").
		Where("id_dorama = ? and username = ?", idD, username).
		Take(&resDB)
	if result.Error != nil {
		return nil, fmt.Errorf("db: %w", result.Error)
	}
	m := model.Review{
		Username: resDB.Username,
		Mark:     resDB.Mark,
		Content:  resDB.Content,
	}
	return &m, nil
}

func (r *ReviewRepo) CreateReview(ctx context.Context, idD int, record *model.Review) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo CreateReview")
	defer span.End()
	m := reviewModel{
		IdDorama: idD,
		Username: record.Username,
		Mark:     record.Mark,
		Content:  record.Content,
	}
	result := r.db.WithContext(ctx).Table("dorama_set.review").Create(&m)
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}
func (r *ReviewRepo) DeleteReview(ctx context.Context, username string, idD int) error {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo DeleteReview")
	defer span.End()
	result := r.db.WithContext(ctx).Table("dorama_set.review").
		Where("username = ? and id_dorama = ?", username, idD).
		Delete(&reviewModel{})
	if result.Error != nil {
		return fmt.Errorf("db: %w", result.Error)
	}
	return nil
}

func (r *ReviewRepo) AggregateRate(ctx context.Context, idD int) (float64, int, error) {
	ctx, span := tracing.StartSpanFromContext(ctx, "Repo AggregateRate")
	defer span.End()
	var resDB struct {
		Avg   float64
		Count int
	}
	result := r.db.WithContext(ctx).Table("dorama_set.review").
		Where("id_dorama = ?", idD).
		Select("AVG(mark), COUNT(*)").
		Take(&resDB)
	if result.Error != nil {
		return 0, 0, fmt.Errorf("db: %w", result.Error)
	}
	return resDB.Avg, resDB.Count, nil
}
