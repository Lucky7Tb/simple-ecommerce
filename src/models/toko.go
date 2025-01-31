package model

import (
	"database/sql"
	"time"
)

type Toko struct {
	ID        uint
	IdUser    uint
	User      User `gorm:"foreignKey:IdUser;references:ID"`
	NamaToko  string
	UrlFoto   sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Toko) TableName() string {
	return "toko"
}
