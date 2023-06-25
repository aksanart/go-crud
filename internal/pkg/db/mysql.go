package db

import (
	"aksan-go-crud/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBMysqlInf interface {
	Migrate() error
	Ping() error
	Client() *gorm.DB
}

type DBMysql struct {
	DB *gorm.DB
}

func NewDB(cfg *configs.Config) (DBMysqlInf, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DBMysql{DB: database}, nil
}

func (d *DBMysql) Migrate() error {
	for _, model := range MigrateModels() {
		err := d.Client().Debug().AutoMigrate(model.Model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DBMysql) Ping() error {
	db, err := d.Client().DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (d *DBMysql) Client() *gorm.DB {
	return d.DB
}
