package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID            uint
	Nama          string
	KataSandi     string
	Notelp        string `gorm:"type:varchar(16);uniqueIndex"`
	TanggalLahir  time.Time
	JenisKelamiin string
	Tentang       sql.NullString
	Pekerjaan     string
	Email         string `gorm:"type:varchar(100);uniqueIndex"`
	IdProvinsi    uint
	IdKota        uint
	IsAdmin       bool
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}

func (User) TableName() string {
	return "user"
}
