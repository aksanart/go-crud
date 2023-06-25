package service

import (
	"aksan-go-crud/internal/entity"
)

func (s *serviceStruct) SaveService(req *entity.SaveRequest) error {
	data := &entity.DataGame{
		Stadium: req.Stadium,
		Mdate:   req.Mdate,
		Team1:   req.Team1,
		Team2:   req.Team2,
	}
	return s.Repo.Save(data)
}
