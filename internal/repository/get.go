package repository

import "aksan-go-crud/internal/entity"

func (r *RepoStruct) Get() ([]*entity.DataGame, error) {
	var datas []*entity.DataGame
	tx := r.DB.Find(&datas)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return datas, nil
}
