package model

import "time"

type LogProduk struct {
	ID            uint
	IdProduk      uint
	Produk        Produk `gorm:"foreignKey:IdProduk;references:ID"`
	IdToko        uint
	Toko          Toko `gorm:"foreignKey:IdToko;references:ID"`
	IdCategory    uint
	Category      Category `gorm:"foreignKey:IdCategory;references:ID"`
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Deskripsi     string
	CreatedAt     time.Time
	UpdayedAt     time.Time
}

func (LogProduk) TableName() string {
	return "log_produk"
}
