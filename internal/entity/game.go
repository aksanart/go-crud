package entity

import "time"

type Game struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
}

type SaveRequest struct {
	Mdate   string `json:"mdate" binding:"required"`
	Stadium string `json:"stadium" binding:"required"`
	Team1   string `json:"team1" binding:"required"`
	Team2   string `json:"team2" binding:"required"`
}

type DataGame struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Mdate     string    `json:"mdate" binding:"required"`
	Stadium   string    `json:"stadium" binding:"required"`
	Team1     string    `json:"team1" binding:"required"`
	Team2     string    `json:"team2" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updatedAt"`
}

// TableName isrename nama table
func (DataGame) TableName() string {
	return "data_game"
}
