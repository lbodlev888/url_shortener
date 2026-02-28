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

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	if host == "" { host = "localhost" }
	if port == "" { port = "5432" }
	if user == "" { user = "postgres" }
	if dbname == "" { dbname = "shorts" }

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
		url text not null,
		ipaddress text not null,
		clicks integer not null default 0,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		userId integer not null,

		constraint fk_user foreign key (userId) references users(id) on delete restrict,
		constraint nenegative_clicks check (clicks >= 0))`)
}
