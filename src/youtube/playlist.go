package youtube

import (
	"errors"
	"fmt"
	"log"
)

func (h *Handler) LoadPlaylistData(playlistId string) (res Result, err error) {

	nextPageToken := ""
	for {
		itemsUrl := fmt.Sprintf(
			"https://www.googleapis.com/youtube/v3/playlistItems?key=%s&playlistId=%s&part=snippet,id&maxResults=%d",
			h.key,
			playlistId,
			50, // youtube limit
		)
		if nextPageToken != "" {
			itemsUrl += fmt.Sprintf("&pageToken=%s", nextPageToken)
		}

		log.Println(itemsUrl)

		var result playlistResult
		err = getJson(itemsUrl, &result)
		if err != nil {
			return
		}

		if len(result.Items) == 0 {
			return res, errors.New("no items in playlist")
		}

		for _, i := range result.Items {
			if i.Snippet.ResourceId.VideoId == "" {
				log.Println(fmt.Sprintf("wrong kind id [%s]", i.Snippet.ResourceId.VideoId))
				continue

			}
			if i.Snippet.ResourceId.Kind != "youtube#video" {
				log.Println(fmt.Sprintf("wrong kind for id [%s], kind [%s]", i.Snippet.ResourceId.VideoId, i.Snippet.ResourceId.Kind))
				continue
			}

			res.Items = append(res.Items, &Item{
				Id:    i.Snippet.ResourceId.VideoId,
				Title: i.Snippet.Title,
				Thumbnail: Thumbnail{
					Url:    i.Snippet.Thumbnails.High.Url,
					Width:  i.Snippet.Thumbnails.High.Width,
					Height: i.Snippet.Thumbnails.High.Height,
				},
			})
		}

		res.Channel.Id = playlistId
		res.Channel.Title = result.Items[0].Snippet.Title
		res.Channel.Thumbnail.Url = result.Items[0].Snippet.Thumbnails.High.Url
		res.Channel.Thumbnail.Height = result.Items[0].Snippet.Thumbnails.High.Height
		res.Channel.Thumbnail.Width = result.Items[0].Snippet.Thumbnails.High.Width

		res.TotalResults = result.PageInfo.TotalResults

		if len(result.Items) == 0 ||
			len(res.Items) >= h.itemsLimit ||
			result.NextPageToken == "" {
			break
		}

		nextPageToken = result.NextPageToken
	}

	return
}

type playlistResult struct {
	NextPageToken string
	PageInfo      struct {
		TotalResults int
	}
	Items []struct {
		Snippet struct {
			Title      string
			ResourceId struct {
				Kind    string
				VideoId string
			}
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
