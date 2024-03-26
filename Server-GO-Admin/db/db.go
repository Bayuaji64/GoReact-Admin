package db

import (
	"fmt"
	"os"

	"example.com/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Deklarasi variabel DB sebagai pointer ke *gorm.DB
var DB *gorm.DB

func InitDB() error { // Fungsi mengembalikan error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	// Membuat DSN dari nilai .env
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s", dbUser, dbPassword, dbHost, dbName, dbCharset, dbParseTime, dbLoc)

	// Membuka koneksi ke database menggunakan GORM
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not connect with the database: %w", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{})

	return nil
}
