package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	psqlInfo = "host=host port=8080 user=postgres password=1 " +
		"dbname=postgres sslmode=disable"
	driverName = "postgres"
)

func Insert(c *gin.Context) {
	db, err := sql.Open(driverName, psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	result, err := db.Exec("insert into shurl (shortURL, longURL) values ($1, $2)",
		"short", "long")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
	fmt.Println("Successfully connected!")
}
