package service

import (
	"aksan-go-crud/internal/entity"
	"aksan-go-crud/internal/repository"
)

type ServiceIntf interface {
	SaveService(req *entity.SaveRequest) error
	GetService() (interface{}, error)
}

type serviceStruct struct {
	Repo repository.RepoInf
}

func NewService(repo repository.RepoInf) ServiceIntf {
	return &serviceStruct{
		Repo: repo,
	}
}
