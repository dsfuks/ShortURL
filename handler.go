package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	storage MemoryStorage
}

func NewHandler(m MemoryStorage) *Handler {
	return &Handler{storage: m}
}
func (h *Handler) CreateURL(c *gin.Context) {
	var url URL
	url.Value = c.Param("url")
	h.storage.Insert(url)
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortURL": h.storage.data[url].Value,
		"longURL":  c.Param("url"),
	})
}

func (h *Handler) GetURL(c *gin.Context) {
	var url URL
	url.Value = c.Param("url")
	get, err := h.storage.Get(url)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"shortURL": c.Param("url"),
		"longURL":  get.Value,
	})
}
