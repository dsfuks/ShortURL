package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type DbHandler struct {
	storage DBStorage
}

func NewDbHandler(db DBStorage) *DbHandler {
	return &DbHandler{storage: db}
}
func (h *DbHandler) CreateURL(c *gin.Context) {
	longURL := c.Param("url")
	_, err := h.storage.data.Query("select * from shurl where longurl = $1",
		longURL)
	if errors.Is(err, sql.ErrNoRows) {
		h.storage.Insert(longURL)
	}

	rows, err := h.storage.data.Query("select * from shurl where longurl = $1",
		longURL)
	var shortURL string
	for rows.Next() {
		user := new(struct {
			shorturl string
			longurl  string
		})
		err := rows.Scan(&user.shorturl, &user.longurl)
		if err != nil {
			log.Println(err)
			return
		}
		shortURL = user.shorturl
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortURL": shortURL,
		"longURL":  longURL,
	})
}
func (h *DbHandler) GetURL(c *gin.Context) {
	shortURL := c.Param("url")
	rows, _ := h.storage.data.Query("select * from shurl where shorturl = $1",
		shortURL)
	var longURL string
	for rows.Next() {
		user := new(struct {
			shorturl string
			longurl  string
		})
		err := rows.Scan(&user.shorturl, &user.longurl)
		if err != nil {
			log.Println(err)
			return
		}
		longURL = user.longurl
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortURL": shortURL,
		"longURL":  longURL,
	})
}
