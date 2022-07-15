package postgres

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"os"
	"shorty/base63"
	"shorty/model"
)

const (
	HOST = "database"
	PORT = "5432"
)

type PostgresInstance struct {
	db *gorm.DB
}

func (p *PostgresInstance) Setup() {
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
		log.Println(err)
	}
}

func (p *PostgresInstance) Exists(id int64) bool {
	var shorty model.Shorty
	r := p.db.Where("id = ?", id).Find(&shorty)
	return r.RowsAffected > 0
}

func (p *PostgresInstance) Save(url string) (string, error) {
	id := rand.Int63()
	for p.Exists(id) {
		id = rand.Int63()
	}
	shortUrl := base63.Encode(id)
	var shorty = model.Shorty{
		ID:       id,
		URL:      url,
		ShortUrl: shortUrl,
	}
	err := p.db.Create(&shorty).Error
	if err != nil {
		log.Println(err)
	}
	return "http://localhost:3000/" + shortUrl, err
}

func (p *PostgresInstance) Get(shortUrl string) (string, error) {
	var shorty model.Shorty
	tx := p.db.Where("short_url = ?", shortUrl).First(&shorty)
	if tx.Error != nil {
		return "", errors.New("shorty not found")
	}
	return shorty.URL, nil
}
