package model

type Language struct {
	Name  string `json:"name" gorm:"-"`
	Ratio int    `json:"ratio" gorm:"-"`
}

type Commit struct {
	Date  string `json:"date" gorm:"-"`
	Count int    `json:"count" gorm:"-"`
}

type Git interface {
	[]*Language | []*Commit
}
