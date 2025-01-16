package model

import (
	"time"
)

type Tabler interface {
	TableName() string
}

type Alamat struct {
	ID           uint
	IdUser       uint
	User         User `gorm:"foreignKey:IdUser;references:ID"`
	JudulAlamat  string
	NamaPenerima string
	Notelp       string
	DetailAlamat string
	CreatedAt    time.Time
	UpdayedAt    time.Time
}

func (Alamat) TableName() string {
	return "alamat"
}
