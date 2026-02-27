package services

import (
	"encoding/base64"

	"github.com/lbodlev888/url_shortener/models"
)

func RegisterUser(user models.User) error {
	user.Password, user.Salt = deriveKey([]byte(user.Password))

	_, err := db.NamedExec("INSERT INTO users (username, email, password, salt) values (:username, :email, :password, :salt)", user)
	return err
}

func LoginUser(user models.User) (bool, error) {
	plainPassword := user.Password
	err := db.Get(&user, "SELECT password, salt FROM users WHERE email=$1 OR username=$1", user.Username)
	if err != nil { return false, err }

	saltB, err := base64.StdEncoding.DecodeString(user.Salt)
	if err != nil { return false, err }

	hashB, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil { return false, err }

	return checkKey([]byte(plainPassword), hashB, saltB), nil
}
