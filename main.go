package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const (
	port = ":8090"
)

func main() {
	router := gin.Default()
	log.Println("Docker works..")
	if os.Args[1] == "memory" {
		var memory = NewMemoryStorage()
		handler := NewHandler(*memory)
		router.POST("/:url", handler.CreateURL)
		router.GET("/:url", handler.GetURL)
	} else {
		log.Println("This is db")
		var db = NewDBStorage()
		dbHandler := NewDbHandler(*db)
		if db.err != nil {
			log.Println("Can't connect to db")
		}
		router.POST("/:url", dbHandler.CreateURL)
		router.GET("/:url", dbHandler.GetURL)
	}
	router.Run(port)
}
