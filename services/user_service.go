package services

import (
	"encoding/base64"
	"fmt"

	"github.com/lbodlev888/url_shortener/models"
)

func RegisterUser(user models.User) error {
	user.Password, user.Salt = deriveKey([]byte(user.Password))

	_, err := db.NamedExec("INSERT INTO users (username, email, password, salt) values (:username, :email, :password, :salt)", user)
	return err
}

func LoginUser(user models.User) (string, error) {
	plainPassword := user.Password
	err := db.Get(&user, "SELECT password, salt FROM users WHERE username=$1", user.Username)
	if err != nil {
		return "", err
	}

	saltB, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil {
		return "", err
	}

	hashB, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return "", err
	}

	if checkKey([]byte(plainPassword), hashB, saltB) {
		t := models.NewToken(user.Username)
		return t.Issue(key), nil
	} else {
		return "", fmt.Errorf("invalid_creds")
	}
}

func getUserIdByUsername(username string) (int, error) {
	var id int
	err := db.Get(&id, "SELECT id FROM users WHERE username=$1", username)
	return id, err
}
