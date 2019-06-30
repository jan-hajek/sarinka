package downloader

import (
	"log"

	"git.vsh-labs.cz/cml/nest/src/storage"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type Handler struct {
	storage *storage.Handler
	youtube *youtube.Handler
}

func New(storage *storage.Handler, youtube *youtube.Handler) *Handler {
	return &Handler{
		storage: storage,
		youtube: youtube,
	}
}

func (h *Handler) SaveData(channelIds []string) {
	for _, channelId := range channelIds {
		err := h.saveChannelData(channelId)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (h *Handler) saveChannelData(channelId string) error {
	var data storage.Data

	res, err := h.youtube.LoadData(channelId, "")
	if err != nil {
		return err
	}

	data.TotalResults = res.TotalResults
	data.Items = append(data.Items, res.Items...)

	return h.storage.SaveData(channelId, &data)
}

//func (h *Handler) SaveChannelData(channelId string) error {
//	var nextPageToken string
//	var i int
//
//	var data storage.Data
//
//	for {
//		i++
//		if i > 20 {
//			panic("xxxx")
//		}
//		res, err := h.youtube.LoadData(channelId, nextPageToken)
//		if err != nil {
//			return err
//		}
//		nextPageToken = res.NextPageToken
//
//		data.TotalResults = res.TotalResults
//		data.Items = append(data.Items, res.Items...)
//
//		if nextPageToken == "" {
//			break
//		}
//	}
//
//	return h.storage.SaveData(channelId, &data)
//}
