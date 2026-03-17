package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lbodlev888/url_shortener/models"
)

func NewUrl(token *models.Token, url, ipaddress string) error {
	status, err := testURL(url)
	if err != nil {
		return err
	}
	if !status {
		return fmt.Errorf("Malicious URL detected")
	}

	userId, err := getUserIdByUsername(token.Username)
	if err != nil {
		return err
	}

	rand.Read(path)
	pathHex := hex.EncodeToString(path)

	result, err := db.Exec("INSERT INTO shorts (path, url, ipaddress, userid) values ($1, $2, $3, $4)", pathHex, url, ipaddress, userId)
	if err != nil {
		return err
	}

	rowsAf, _ := result.RowsAffected()
	if rowsAf > 0 {
		ctx := context.Background()
		rdb.Del(ctx, "urls:all:"+token.Username).Result()
	}

	return nil
}

func GetLongUrl(path string) (string, error) {
	ctx := context.Background()
	key := "urls:" + path
	var url string

	val, err := rdb.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}

	err = db.Get(&url, "UPDATE shorts SET clicks = clicks + 1 WHERE path = $1 RETURNING url", path)
	if err != nil {
		return "", err
	}
	rdb.Set(ctx, key, url, 10*time.Minute)

	return url, nil
}

func GetAllShorts(token *models.Token) ([]models.ShortUrl, error) {
	ctx := context.Background()
	key := "urls:all:" + token.Username
	var shorts []models.ShortUrl
	val, err := rdb.Get(ctx, key).Result()
	if err == nil {
		if json.Unmarshal([]byte(val), &shorts) == nil {
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
	return shorts, err
}

func DeleteShort(token *models.Token, path string) error {
	result, err := db.Exec("DELETE FROM shorts USING users WHERE shorts.userid=users.id AND path=$1", path)
	rowsAf, _ := result.RowsAffected()

	if rowsAf > 0 {
		ctx := context.Background()
		rdb.Del(ctx, "urls:all:"+token.Username, "urls:"+path)
	}
	return err
}
