package model

import "time"

type Tag struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(30)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
