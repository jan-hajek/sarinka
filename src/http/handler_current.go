package http

import (
	"encoding/json"
	"net/http"

	"git.vsh-labs.cz/cml/nest/src/app"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

func (h *Handler) currentHandler(w http.ResponseWriter, r *http.Request) {
	h.mx.Lock()
	defer h.mx.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	channelId := getChannelId(r)
	id := getId(r)

	channel := h.app.GetChannel(channelId)
	_, item := channel.GetItems(id, 2)

	result := PlayResponse{
		Channel: channel,
		Current: item[0],
		Next:    item[1],
	}

	data, err := json.Marshal(&result)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

type PlayResponse struct {
	Channel *app.Channel
	Current *youtube.Item
	Next    *youtube.Item
}
