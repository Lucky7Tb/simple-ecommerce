package model

import "time"

type DetailTrx struct {
	ID          uint
	IdTrx       uint
	Trx         Trx `gorm:"foreignKey:IdTrx;references:ID"`
	IdLogProduk uint
	LogProduk   LogProduk `gorm:"foreignKey:IdLogProduk;references:ID"`
	IdToko      uint
	Toko        Toko `gorm:"foreignKey:IdToko;references:ID"`
	Kuantitas   uint
	HargaTotal  uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (DetailTrx) TableName() string {
	return "detail_trx"
}
