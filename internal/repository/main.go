package repository

import (
	"aksan-go-crud/internal/entity"
	"aksan-go-crud/internal/pkg/db"

	"gorm.io/gorm"
)

type RepoInf interface {
	Save(data *entity.DataGame) error
	Get() ([]*entity.DataGame, error)
}

type RepoStruct struct {
	DB *gorm.DB
}

func NewRepository(db db.DBMysqlInf) RepoInf {
	return &RepoStruct{
		DB: db.Client(),
	}
}
