package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	"github.com/lbodlev888/url_shortener/models"
)

func NewUrl(raw_token, url, ipaddress string) error {
	token, _ := models.ParseToken(raw_token)
	userId, err := getUserIdByUsername(token.Username)
	if err != nil {
		return err
	}

	rand.Read(path)
	pathHex := hex.EncodeToString(path)

	_, err = db.Exec("INSERT INTO shorts (path, url, ipaddress, userid) values ($1, $2, $3, $4)", pathHex, url, ipaddress, userId)
	if err != nil {
		return err
	}

	ctx := context.Background()
	rdb.Del(ctx, "urls:all:"+token.Username).Result()
	return nil
}

func GetLongUrl(path string) (string, error) {
	ctx := context.Background()
	key := "urls:" + path
	var url string

	val, err := rdb.Get(ctx, key).Result()
	if err == nil {
		log.Println("got from redis")
		return val, nil
	}

	err = db.Get(&url, "SELECT url FROM shorts WHERE path=$1", path)
	if err != nil {
		return "", err
	}
	rdb.Set(ctx, key, url, 10*time.Minute)

	log.Println("got from postgres")
	return url, nil
}

func GetAllShorts(raw_token string) ([]models.ShortUrl, error) {
	token, _ := models.ParseToken(raw_token)

	ctx := context.Background()
	key := "urls:all:" + token.Username
	var shorts []models.ShortUrl
	val, err := rdb.Get(ctx, key).Result()
	if err == nil {
		if json.Unmarshal([]byte(val), &shorts) == nil {
			log.Println("got from redis")
			return shorts, nil
		}
	}

	err = db.Select(&shorts, "SELECT path, url, clicks, created_at FROM shorts JOIN users ON shorts.userid=users.id WHERE username=$1 OR email=$1", token.Username)
	if err != nil {
		return nil, err
	}

	if len(shorts) > 0 {
		data, _ := json.Marshal(shorts)
		rdb.Set(ctx, key, data, 10*time.Minute)
	}
	log.Println("got from postgres")
	return shorts, err
}

func DeleteShort(raw_token, path string) error {
	token, _ := models.ParseToken(raw_token)
	ctx := context.Background()
	rdb.Del(ctx, "urls:all:"+token.Username, "urls:"+path)

	_, err := db.Exec("DELETE FROM shorts WHERE path=$1", path)
	return err
}

func Increment(path string) {
	db.Exec("UPDATE shorts SET clicks=clicks+1 WHERE path=$1", path)
}
