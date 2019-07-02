package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) channelsHandler(w http.ResponseWriter, r *http.Request) {
	h.mx.Lock()
	defer h.mx.Unlock()

	var result channelsResponse
	for _, ch := range h.app.GetChannels() {
		result.Channels = append(
			result.Channels,
			channelsResponseITem{
				Id:           ch.Id,
				Name:         ch.Name,
				ThumbnailUrl: ch.ThumbnailUrl,
			},
		)
	}

	data, err := json.Marshal(&result)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

type channelsResponse struct {
	Channels []channelsResponseITem
}

type channelsResponseITem struct {
	Id           string
	Name         string
	ThumbnailUrl string
}
