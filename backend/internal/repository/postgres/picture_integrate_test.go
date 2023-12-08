//go:build integration

package postgres

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/logic/model"
	"context"
	"gorm.io/gorm"
	"testing"
)

func TestPictureRepo_CreatePicture(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		record model.Picture
		id     int
		tbl    string
	}
	pic := model.Picture{URL: "test"}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(id int) error
	}{
		{
			name:    "success",
			fields:  fields{db: db},
			args:    args{record: pic},
			wantErr: false,
			check: func(id int) error {
				res := db.Table("dorama_set.picture").Where("id = ?", id).Take(&model.Picture{})
				return res.Error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PictureRepo{
				db: tt.fields.db,
			}
			got, err := p.CreatePicture(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePicture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := tt.check(got); (err != nil) != tt.wantErr {
				t.Errorf("CreatePicture() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
