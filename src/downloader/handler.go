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
	res, err := h.youtube.LoadData(channelId, "")
	if err != nil {
		return err
	}

	return h.storage.SaveData(channelId, res)
}
