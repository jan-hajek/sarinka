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
	channelUrl := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/channels?key=%s&id=%s&part=snippet,id",
		h.key,
		channelId,
	)
	log.Println(channelUrl)

	var channelResult channelResult
	err = getJson(channelUrl, &channelResult)
	if err != nil {
		return
	}

	if len(channelResult.Items) == 0 {
		log.Fatal("no data fo channel")
	}

	res.Channel.Id = channelResult.Items[0].Id
	res.Channel.Title = channelResult.Items[0].Snippet.Title
	res.Channel.Thumbnail.Url = channelResult.Items[0].Snippet.Thumbnails.High.Url
	res.Channel.Thumbnail.Height = channelResult.Items[0].Snippet.Thumbnails.High.Height
	res.Channel.Thumbnail.Width = channelResult.Items[0].Snippet.Thumbnails.High.Width

	itemsUrl := fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/search?key=%s&channelId=%s&part=snippet,id&order=date&maxResults=%d",
		h.key,
		channelId,
		h.pageLimit,
	)
	if nextPageToken != "" {
		itemsUrl += fmt.Sprintf("&pageToken=%s", nextPageToken)
	}

	log.Println(itemsUrl)

	var itemsResult itemsResult
	err = getJson(itemsUrl, &itemsResult)
	if err != nil {
		return
	}

	for _, i := range itemsResult.Items {
		if i.Id.VideoId == "" {
			log.Println(fmt.Sprintf("wrong kind id [%s]", i.Id.VideoId))
			continue

		}
		if i.Kind != "youtube#searchResult" {
			log.Println(fmt.Sprintf("wrong kind for id [%s], kind [%s]", i.Id.VideoId, i.Kind))
			continue
		}

		res.Items = append(res.Items, &Item{
			Id:    i.Id.VideoId,
			Title: i.Snippet.Title,
			Thumbnail: Thumbnail{
				Url:    i.Snippet.Thumbnails.Medium.Url,
				Width:  i.Snippet.Thumbnails.Medium.Width,
				Height: i.Snippet.Thumbnails.Medium.Height,
			},
		})
	}

	res.TotalResults = itemsResult.PageInfo.TotalResults
	res.NextPageToken = itemsResult.NextPageToken

	return
}

type Result struct {
	Channel       Channel
	Items         []*Item
	TotalResults  int
	NextPageToken string
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

type channelResult struct {
	Items []struct {
		Id      string
		Snippet struct {
			Title      string
			Thumbnails struct {
				High struct {
					Url    string
					Width  int
					Height int
				}
			}
		}
	}
}

type itemsResult struct {
	NextPageToken string
	PageInfo      struct {
		TotalResults int
	}
	Items []struct {
		Kind string
		Id   struct {
			VideoId string
		}
		Snippet struct {
			Title      string
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
