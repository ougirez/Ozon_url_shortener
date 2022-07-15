package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shorty/model"
	"strings"
	"testing"
)

func testSetup() *PostgresInstance {
	var p = new(PostgresInstance)
	dbUser := "postgres"
	dbPassword := "0000"
	dbName := "shortyTest"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", PORT, dbUser, dbPassword, dbName)
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = p.db.AutoMigrate(&model.Shorty{})
	if err != nil {
		log.Println(err)
	}
	return p
}

func TestExists(t *testing.T) {
	p := testSetup()
	p.db.Exec("DELETE FROM shorties")
	err := p.db.Create(&model.Shorty{
		ID:       52345234312,
		URL:      "someUrl",
		ShortUrl: "someShortUrl",
	}).Error
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	var id int64 = 52345234312
	want := true
	res := p.Exists(id)
	if !want == res {
		t.Fatalf("got %v, want %v", res, want)
	}
}

func TestNotExists(t *testing.T) {
	p := testSetup()
	p.db.Exec("DELETE FROM shorties")
	err := p.db.Create(&model.Shorty{
		ID:       52345234312,
		URL:      "someUrl",
		ShortUrl: "someShortUrl",
	}).Error
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	var id int64 = 165215242342
	want := false
	res := p.Exists(id)
	if !want == res {
		t.Fatalf("got %v, want %v", res, want)
	}
}

func TestGet(t *testing.T) {
	p := testSetup()
	p.db.Exec("DELETE FROM shorties")
	err := p.db.Create(&model.Shorty{
		ID:       52345234312,
		URL:      "https://www.youtube.com/watch?v=8YJxWV5yPx0",
		ShortUrl: "G1rW1b_wva",
	}).Error
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	wantRes := "https://www.youtube.com/watch?v=8YJxWV5yPx0"
	res, err := p.Get("G1rW1b_wva")
	if err != nil {
		t.Fatalf("unexpected error uccured: %s", err.Error())
	}
	if res != wantRes {
		t.Fatalf("wrong shortUrl: got %s, want %s", res, wantRes)
	}
}

func TestGetNotFound(t *testing.T) {
	p := testSetup()
	p.db.Exec("DELETE FROM shorties")
	err := p.db.Create(&model.Shorty{
		ID:       52345234312,
		URL:      "https://www.youtube.com/watch?v=8YJxWV5yPx0",
		ShortUrl: "G1rW1b_wva",
	}).Error
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	wantRes, wantErrMsg := "", "shorty not found"
	res, err := p.Get("27I5tQ6N4E")
	if err == nil {
		t.Fatalf("error is %s, want nil", err.Error())
	}
	if err.Error() != "shorty not found" {
		t.Fatalf("wrong error: got %s, want %s", err.Error(), wantErrMsg)
	}
	if res != wantRes {
		t.Fatalf("wrong url: got %s, want %s", res, wantRes)
	}
}

func TestSave(t *testing.T) {
	p := testSetup()
	p.db.Exec("DELETE FROM shorties")
	url := "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	short, resErr := p.Save(url)
	short = strings.TrimPrefix(short, "http://localhost:3000/")
	if resErr != nil {
		t.Fatal(resErr)
	}
	getResult, getErr := p.Get(short)
	if getErr != nil {
		t.Fatalf("shortUrl have to exists but it doesn't")
	}
	if getResult != url {
		t.Fatalf("url saved wrong: got %s, want %s", getResult, url)
	}
}
