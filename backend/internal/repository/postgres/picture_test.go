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
	dr := DoramaRepo{db: db}
	idD, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success in doramapicture",
			fields:  fields{db: db},
			args:    args{record: model.Picture{}, id: idD, tbl: "dorama"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PictureRepo{
				db: tt.fields.db,
			}
			got, err := p.CreatePicture(tt.args.record, tt.args.id, tt.args.tbl)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePicture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			_ = p.DeletePicture(got)
		})
	}
}
