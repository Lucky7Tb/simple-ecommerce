package model

import "time"

type Trx struct {
	ID               uint
	IdUser           uint
	User             User `gorm:"foreignKey:IdUser;references:ID"`
	AlamatPengiriman uint
	Alamat           Alamat `gorm:"foreignKey:AlamatPengiriman;references:ID"`
	HargaTotal       uint
	KodeInvoice      string
	MethodBayar      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Trx) TableName() string {
	return "trx"
}
