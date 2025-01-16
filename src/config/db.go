package config

import (
	"fmt"
	model "simple-ecommerce/src/models"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	var dns string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Error connect to database")
	}

	db.AutoMigrate(&model.User{}, &model.Alamat{}, &model.Category{}, &model.Trx{}, &model.Toko{}, &model.DetailTrx{}, &model.FotoProduk{}, &model.LogProduk{}, &model.Produk{})

	return db
}
