package postgres

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/logic/model"
	"context"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func TestSubscriptionRepo_GetList(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	type fields struct {
		db *gorm.DB
	}
	month := 720 * time.Hour
	tests := []struct {
		name    string
		fields  fields
		want    []model.Subscription
		wantErr bool
	}{
		{
			name:   "success",
			fields: fields{db},
			want: []model.Subscription{{1, "free", 0, month},
				{2, "not free", 50, month}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubscriptionRepo{
				db: tt.fields.db,
			}
			got, err := s.GetList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionRepo_GetSubscription(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	month := 720 * time.Hour
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Subscription
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{1},
			want:    &model.Subscription{1, "free", 0, month},
			wantErr: false,
		},
		{
			name:    "dont exists in db",
			fields:  fields{db},
			args:    args{5},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubscriptionRepo{
				db: tt.fields.db,
			}
			got, err := s.GetSubscription(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionRepo_GetSubscriptionByPrice(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		price int
	}
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	month := 720 * time.Hour
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Subscription
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{0},
			want:    &model.Subscription{1, "free", 0, month},
			wantErr: false,
		},
		{
			name:    "dont exists in db",
			fields:  fields{db},
			args:    args{200},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubscriptionRepo{
				db: tt.fields.db,
			}
			got, err := s.GetSubscriptionByPrice(tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubscriptionByPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubscriptionByPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
