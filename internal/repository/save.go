package repository

import "aksan-go-crud/internal/entity"

func (r *RepoStruct) Save(data *entity.DataGame) error {
	tx := r.DB.Create(data)
	return tx.Error
}
