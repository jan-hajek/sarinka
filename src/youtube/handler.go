package youtube

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	key       string
	pageLimit int
}

func New(key string, pageLimit int) *Handler {
	return &Handler{
		key:       key,
		pageLimit: pageLimit,
	}
}

func (h *Handler) LoadData(channelId string, nextPageToken string) (res Result, err error) {
	url := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/search?key=%s&channelId=%s&part=snippet,id&order=date&maxResults=%d",
		h.key,
		channelId,
		h.pageLimit,
	)
	if nextPageToken != "" {
		url += fmt.Sprintf("&pageToken=%s", nextPageToken)
	}

	log.Println(url)

	var result result
	err = getJson(url, &result)
	if err != nil {
		return
	}

	var items []*Item
	for _, i := range result.Items {
		items = append(items, &Item{
			Id: i.Id.VideoId,
			Thumbnail: Thumbnail{
				Url:    i.Snippet.Thumbnails.Medium.Url,
				Width:  i.Snippet.Thumbnails.Medium.Width,
				Height: i.Snippet.Thumbnails.Medium.Height,
			},
		})
	}

	return Result{
		Items:         items,
		TotalResults:  result.PageInfo.TotalResults,
		NextPageToken: result.NextPageToken,
	}, nil
}

type Result struct {
	Items         []*Item
	TotalResults  int
	NextPageToken string
}

type Item struct {
	Id        string
	Thumbnail Thumbnail
}

type Thumbnail struct {
	Url    string
	Width  int
	Height int
}

type result struct {
	NextPageToken string
	PageInfo      struct {
		TotalResults int
	}
	Items []struct {
		Id struct {
			VideoId string
		}
		Snippet struct {
			Thumbnails struct {
				Medium struct {
					Url    string
					Width  int
					Height int
				}
			}
		}
	}
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
