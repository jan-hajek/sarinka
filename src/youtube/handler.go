package youtube

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	key        string
	itemsLimit int
}

func New(key string, itemsLimit int) *Handler {
	return &Handler{
		key:        key,
		itemsLimit: itemsLimit,
	}
}

type Result struct {
	Channel      Channel
	Items        []*Item
	TotalResults int
}

type Item struct {
	Id        string
	Title     string
	Thumbnail Thumbnail
}

type Channel struct {
	Id        string
	Title     string
	Thumbnail Thumbnail
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
}

func getJson(url string, target interface{}) error {
	myClient := &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
