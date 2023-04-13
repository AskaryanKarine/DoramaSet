package postgres

import (
	"DoramaSet/internal/logic/model"
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
	db := connect()
	dr := DoramaRepo{db: db}
	idD, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	idE, _ := EpisodeRepo{db: db}.CreateEpisode(model.Episode{}, idD)
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
			args:    args{idE},
			want:    &model.Episode{Id: idE, NumSeason: 0, NumEpisode: 0},
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
	_ = EpisodeRepo{db: db}.DeleteEpisode(idE)
	_ = dr.DeleteDorama(idD)
}

func TestEpisodeRepo_GetList(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		idDorama int
	}
	db := connect()
	dr := DoramaRepo{db: db}
	idD, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	idE, _ := EpisodeRepo{db: db}.CreateEpisode(model.Episode{}, idD)
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
			args:    args{idD},
			want:    []model.Episode{{idE, 0, 0}},
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
				t.Errorf("GetList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() got = %v, want %v", got, tt.want)
			}
		})
	}
	_ = EpisodeRepo{db: db}.DeleteEpisode(idE)
	_ = dr.DeleteDorama(idD)
}

func TestEpisodeRepo_MarkEpisode(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		idEp     int
		username string
	}
	db := connect()
	ur := UserRepo{db: db, subRepo: SubscriptionRepo{db: db}}
	_ = ur.CreateUser(model.User{Username: "qwerty"})
	dr := DoramaRepo{db: db}
	idD, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	idE, _ := EpisodeRepo{db: db}.CreateEpisode(model.Episode{}, idD)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db},
			args:    args{idEp: idE, username: "qwerty"},
			wantErr: false,
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
		})
	}
	_ = ur.DeleteUser("qwerty")
	_ = dr.DeleteDorama(idD)
}

func TestEpisodeRepo_CreateEpisode(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		episode model.Episode
		idD     int
	}
	db := connect()
	dr := DoramaRepo{db: db}
	ep := model.Episode{}
	id, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db},
			args:    args{episode: ep, idD: id},
			want:    0,
			wantErr: false,
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
				return
			}
			_ = e.DeleteEpisode(got)
		})
	}
	_ = dr.DeleteDorama(id)
}

func TestEpisodeRepo_DeleteEpisode(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		id int
	}
	db := connect()
	dr := DoramaRepo{db: db}
	ep := model.Episode{}
	id, _ := dr.CreateDorama(model.Dorama{Status: "finish"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			fields:  fields{db: db},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EpisodeRepo{
				db: tt.fields.db,
			}
			idE, _ := e.CreateEpisode(ep, id)
			tt.args.id = idE
			if err := e.DeleteEpisode(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEpisode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	_ = dr.DeleteDorama(id)
}
