package downloader

import (
	"log"

	"git.vsh-labs.cz/cml/nest/src/storage"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

func getChannelIds() []string {
	return []string{
		"UC-Gm4EN7nNNR3k67J8ywF4g", // blippi toys

	}
}

func getPlaylistIds() []string {
	return []string{}
}

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

func (h *Handler) SaveData() {
	for _, channelId := range getChannelIds() {
		err := h.saveChannelData(channelId)
		if err != nil {
			log.Println(err)
		}
	}
	for _, playlistId := range getPlaylistIds() {
		err := h.savePlaylistData(playlistId)
		if err != nil {
			log.Println(err)
		}
	}
}

func (h *Handler) saveChannelData(channelId string) error {
	res, err := h.youtube.LoadChannelData(channelId)
	if err != nil {
		return err
	}

	return h.storage.SaveData(channelId, res)
}

func (h *Handler) savePlaylistData(playlistId string) error {
	res, err := h.youtube.LoadPlaylistData(playlistId)
	if err != nil {
		return err
	}

	return h.storage.SaveData(playlistId, res)
}
