package model

import "time"

type Produk struct {
	ID            uint
	IdToko        uint
	Toko          Toko `gorm:"foreignKey:IdToko;references:ID"`
	IdCategory    uint
	Category      Category `gorm:"foreignKey:IdCategory;references:ID"`
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Stok          uint
	Deskripsi     string
	CreatedAt     time.Time
	UpdayedAt     time.Time
}

func (Produk) TableName() string {
	return "produk"
}
