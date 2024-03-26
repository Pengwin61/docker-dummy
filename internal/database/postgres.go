package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ClientPG struct {
	conDB *sql.DB
}

func NewPostgresClient(host, port, user, pass, db string) (*ClientPG, error) {

	conn, err := connectDB(host, port, user, pass, db)

	if err != nil {
		return nil, err
	}

	return &ClientPG{conDB: conn}, nil
}

func connectDB(host, port, user, password, dbname string) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to", host)
	return db, nil
}

var PostgresIp = "192.168.140.185:5432"

func Init() {
	PostgresIp = "localhost:5432"
}
