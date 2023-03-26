package controllers

import (
	"fmt"
	"math/rand"
	"simple_short_url_server/models"
	"time"

	"github.com/catinello/base62"
)

func SetShortURL(original_url string) string {
	t := time.Now().UnixNano()
	r1 := rand.New(rand.NewSource(t))
	short_url := base62.Encode(r1.Int())

	var existed_url string = models.GetOriginalURL(short_url)
	if existed_url != "" {
		fmt.Println("short URL is existed")
		return existed_url
	}

	models.SetShortURL(short_url, original_url)
	return short_url
}

func GetOriginalURL(short_url string) string {
	var original_url string = models.GetOriginalURL(short_url)
	return original_url
}
