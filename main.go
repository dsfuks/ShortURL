package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.Default()
	if os.Args[1] == "memory" {
		var memory = NewMemoryStorage()
		handler := NewHandler(*memory)
		router.POST("/:url", handler.CreateURL)
		router.GET("/:url", handler.GetURL)
	} else {
		router.POST("/:url")
	}
	router.Run()
}
