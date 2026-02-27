package services

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	dbname := os.Getenv("DBNAME")

	if host == "" { host = "localhost" }
	if port == "" { port = "5432" }

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err = sqlx.Connect("postgres", connStr)
	if err != nil { panic(err) }

	db.Exec(`CREATE TABLE users (
		id serial primary key,
		username varchar(50) not null,
		email text not null,
		password text not null,
		salt text not null)`)
	db.Exec(`CREATE TABLE shorts (
		id serial primary key,
		path varchar(15) unique not null,
		link text not null,
		ipaddress text not null,
		userId integer not null,
		constraint fk_user foreign key (userId) references users(id) on delete restrict)`)
}
