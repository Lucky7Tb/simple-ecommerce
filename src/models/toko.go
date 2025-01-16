package model

import "time"

type Toko struct {
	ID        uint
	IdUser    uint
	User      User `gorm:"foreignKey:IdUser;references:ID"`
	NamaToko  string
	UrlFoto   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Toko) TableName() string {
	return "toko"
}
