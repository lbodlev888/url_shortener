package models

import "time"

type ShortUrl struct {
	Id int `json:"id" db:"id"`
	Path string `json:"path" db:"path"`
	Url string `json:"url" db:"url"`
	IpAddress string `json:"ipaddress" db:"ipaddress"`
	Clicks int `json:"clicks" db:"clicks"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UserId int `json:"userId" db:"userId"`
}
