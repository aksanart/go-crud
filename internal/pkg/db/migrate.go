package db

import "aksan-go-crud/internal/entity"

type Model struct {
	Model interface{}
}

func MigrateModels() []Model {
	return []Model{
		{Model: entity.DataGame{}},
	}
}
