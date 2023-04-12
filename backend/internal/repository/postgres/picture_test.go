package postgres

import (
	"DoramaSet/internal/logic/model"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestPictureRepo_CreatePicture(t *testing.T) {
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
			_ = p.DeletePicture(model.Picture{Id: got})
		})
	}
}

func TestPictureRepo_DeletePicture(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		record model.Picture
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PictureRepo{
				db: tt.fields.db,
			}
			if err := p.DeletePicture(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("DeletePicture() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPictureRepo_GetListDorama(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		idDorama int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Picture
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PictureRepo{
				db: tt.fields.db,
			}
			got, err := p.GetListDorama(tt.args.idDorama)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListDorama() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListDorama() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPictureRepo_GetListStaff(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		idStaff int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Picture
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PictureRepo{
				db: tt.fields.db,
			}
			got, err := p.GetListStaff(tt.args.idStaff)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListStaff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListStaff() got = %v, want %v", got, tt.want)
			}
		})
	}
}
