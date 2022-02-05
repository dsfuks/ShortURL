package main

import (
	"errors"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789")

type URL struct {
	Value string `json:"value"`
}
type MemoryStorage struct {
	data map[URL]URL
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[URL]URL),
	}
}

func (s *MemoryStorage) Insert(e URL) {
	var shLink URL
	if _, ok := s.data[e]; !ok {
		b := make([]rune, 10)
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		shLink.Value = string(b)
		s.data[e] = shLink
	}
}

func (s *MemoryStorage) Get(e URL) (URL, error) {
	if val, ok := reverseMap(s.data)[e]; ok {
		return val, nil
	}
	return URL{}, errors.New("No such url")
}

func reverseMap(m map[URL]URL) map[URL]URL {
	n := make(map[URL]URL)
	for k, v := range m {
		n[v] = k
	}
	return n
}
