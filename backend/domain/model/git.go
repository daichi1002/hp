package model

type Language struct {
	Name  string `json:"name" gorm:"-"`
	Ratio int    `json:"ratio" gorm:"-"`
}
