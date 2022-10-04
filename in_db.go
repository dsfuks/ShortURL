package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"time"
)

const (
	psqlInfo = "host=localhost port=8080 user=postgres password=1 " +
		"dbname=short sslmode=disable"
	driverName = "postgres"
)

type DBStorage struct {
	data *sql.DB
	err  error
}

func NewDBStorage() *DBStorage {
	var DB DBStorage
	DB.data, DB.err = sql.Open(driverName, psqlInfo)
	return &DB
}
func (s *DBStorage) Insert(e string) {
	b := make([]rune, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	_, err := s.data.Exec(query, string(b), e)
	if err != nil {
		log.Println(err)
	}
}

const query = `
	insert into shurl (shorturl,longurl) values ($1, $2) 
	on conflict (shorturl) do update 
	set shorturl = excluded.shorturl,
		longurl = excluded.longurl`
