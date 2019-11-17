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
	_, items := channel.GetItems(id, 2)

	result := PlayResponse{
		Channel: channel,
	}
	if len(items) > 0 {
		result.Current = items[0]
	}
	if len(items) > 1 {
		result.Next = items[1]
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
