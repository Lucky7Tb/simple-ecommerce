package model

import (
	"time"
)

type User struct {
	ID            uint
	Nama          string
	KataSandi     string
	Notelp        string
	TanggalLahir  time.Time
	JenisKelamiin string
	Tentang       string
	Pekerjaan     string
	Email         string
	IdProvinsi    uint
	IdKota        uint
	IsAdmin       bool
	CreatedAt     time.Time
	UpdayedAt     time.Time
}

func (User) TableName() string {
	return "user"
}
