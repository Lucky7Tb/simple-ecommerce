package model

import "time"

type Category struct {
	ID           uint
	NamaCategory string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Category) TableName() string {
	return "category"
}
