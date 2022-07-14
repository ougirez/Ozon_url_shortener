package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"shorty/base63"
	"shorty/model"
)

const (
	HOST = "database"
	PORT = "5432"
)

type PostgresHandler struct {
	db *gorm.DB
}

func (p *PostgresHandler) Setup() {
	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	dbUser = "admin"
	dbPassword = "admin"
	dbName = "admin"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, dbUser, dbPassword, dbName)
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = p.db.AutoMigrate(&model.Shorty{})
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PostgresHandler) Exists(id int64) bool {
	var exists bool
	p.db.Where("id = ?", id).Find(&exists)
	return exists
}

func (p *PostgresHandler) Save(url string) (string, error) {
	id := rand.Int63()
	for p.Exists(id) {
		id = rand.Int63()
	}
	fmt.Println(id)
	shortUrl := base63.Encode(id)
	var shorty = model.Shorty{
		ID:       id,
		URL:      url,
		ShortUrl: shortUrl,
	}
	err := p.db.Create(&shorty).Error
	if err != nil {
		fmt.Println(err)
	}
	return "http://localhost:3000/" + shortUrl, err
}

func (p *PostgresHandler) Get(shortUrl string) (string, error) {
	var shorty model.Shorty
	tx := p.db.Where("short_url = ?", shortUrl).First(&shorty)
	if tx.Error != nil {
		return "", tx.Error
	}
	return shorty.URL, nil
}
