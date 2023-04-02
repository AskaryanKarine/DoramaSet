package interfaces

import "DoramaSet/internal/logic/model"

type IEpisodeRepo interface {
	GetList(idDorama int) ([]model.Episode, error)
	GetEpisode(id int) (model.Episode, error)
	MarkEpisode(idEp int, username string) error
}
