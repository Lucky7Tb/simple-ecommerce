package model

import "time"

type FotoProduk struct {
	ID        uint
	IdProduk  uint
	Produk    Produk `gorm:"foreignKey:IdProduk;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (FotoProduk) TableName() string {
	return "foto_produk"
}
