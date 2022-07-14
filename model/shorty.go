package model

type Shorty struct {
	ID       int64  `json:"id" gorm:"primaryKey"`
	URL      string `json:"url" gorm:"not null"`
	ShortUrl string `json:"short_url" gorm:"unique not null"`
}
