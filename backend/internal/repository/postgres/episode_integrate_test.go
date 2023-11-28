//go:build integration

package postgres

import (
	"DoramaSet/internal/container"
	"DoramaSet/internal/logic/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestEpisodeRepo_GetEpisode(t *testing.T) {
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
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Episode
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{1},
			want:    &model.Episode{Id: 1, NumSeason: 1, NumEpisode: 1},
			wantErr: false,
		},
		{
			name:    "don't exist in db",
			fields:  fields{db},
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			got, err := e.GetEpisode(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpisode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEpisode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEpisodeRepo_GetList(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

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
		want    []model.Episode
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{1},
			want:    []model.Episode{{1, 1, 1}},
			wantErr: false,
		},
		{
			name:    "dont exists in db",
			fields:  fields{db},
			args:    args{-1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			got, err := e.GetList(tt.args.idDorama)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaffList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaffList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEpisodeRepo_MarkEpisode(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		idEp     int
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		check   func(id int, username string) error
	}{
		{
			name:    "success",
			fields:  fields{db: db},
			args:    args{idEp: 1, username: "test"},
			wantErr: false,
			check: func(id int, username string) error {
				res := db.Table("dorama_set.userepisode").
					Where("id_episode = ? and username = ?", id, username).Take(&episodeModel{})
				return res.Error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			if err := e.MarkEpisode(tt.args.idEp, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("MarkEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(tt.args.idEp, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("MarkEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEpisodeRepo_CreateEpisode(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		episode model.Episode
		idD     int
	}
	ep := model.Episode{Id: 2, NumEpisode: 2, NumSeason: 1}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		check   func(id int) error
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{episode: ep, idD: 1},
			want:    0,
			wantErr: false,
			check: func(id int) error {
				var m episodeModel
				res := db.Table("dorama_set.episode").
					Where("id = ?", id).Take(&m)
				if res.Error != nil {
					return res.Error
				}
				ep1 := episodeModel{ID: id, NumSeason: ep.NumSeason, NumEpisode: ep.NumEpisode, IdDorama: 1}
				if !reflect.DeepEqual(m, ep1) {
					return errors.New("error")
				}
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			got, err := e.CreateEpisode(tt.args.episode, tt.args.idD)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(got); (err != nil) != tt.wantErr {
				t.Errorf("CreateEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEpisodeRepo_DeleteEpisode(t *testing.T) {
	dbContainer, db, err := container.SetupTestDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer dbContainer.Terminate(context.Background())

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
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
			args:    args{1},
			wantErr: false,
			check: func(id int) error {
				res := db.Table("dorama_set.episode").Where("id = ?", id).Take(&episodeModel{})
				return res.Error
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			if err := e.DeleteEpisode(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tt.check(tt.args.id); (err != nil) != !tt.wantErr {
				t.Errorf("CreateEpisode() error = %v, wantErr %v", err, !tt.wantErr)
			}
		})
	}
}
