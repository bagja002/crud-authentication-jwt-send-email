package database

import (
	"fmt"
	"pertemuan4/enitity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Konfigurasi database
	username := "user"
	password := "password"
	host := "127.0.0.1"
	port := "3306"
	databaseName := "contohdb"

	// Buat DSN dari variabel
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, databaseName)

	// Koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi database:", err)
		return
	}

	DB = db

	// Auto Migrate
	db.AutoMigrate(
		&enitity.Users{},
		&enitity.Admin{},
		&enitity.Pelanggan{},
	)
}
