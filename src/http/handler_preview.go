package http

import (
	"encoding/json"
	"net/http"

	"git.vsh-labs.cz/jelito/sarinka/src/youtube"
)

func (h *Handler) previewHandler(w http.ResponseWriter, r *http.Request) {
	h.mx.Lock()
	defer h.mx.Unlock()

	channelId := getChannelId(r)
	id := getId(r)

	channel := h.app.GetChannel(channelId)
	position, items := channel.GetItems(id, 10)

	var result previewResponse
	result.Items = items
	result.TotalCount = channel.GetDownloadedCount()
	if position == 0 {
		result.Position = result.TotalCount
	} else {
		result.Position = position
	}

	data, err := json.Marshal(&result)
	if err != nil {
		panic(err)
	}

	w.Write(data)
}

type previewResponse struct {
	Position   int
	TotalCount int
	Items      []*youtube.Item
}
