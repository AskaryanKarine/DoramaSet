package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
	"time"
)

func connect() *gorm.DB {
	dsn := "host=localhost user=karine password=12346 dbname=DoramaSet sslmode=disable"
	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return pureDB
}

var db = connect()

func TestSubscritionRepo_GetList(t *testing.T) {
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
				{2, "standart subs", 100, month}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubscritionRepo{
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

func TestSubscritionRepo_GetSubscription(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
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
			s := SubscritionRepo{
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

func TestSubscritionRepo_GetSubscriptionByPrice(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		price int
	}
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
			s := SubscritionRepo{
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