package http

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) previewHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	preview := preview{
		TotalCount: h.app.TotalCount(),
		Current:    h.app.CurrentIndex() + 1,
	}

	for _, i := range h.app.Preview() {
		preview.Items = append(
			preview.Items,
			Item{
				Url:    i.Thumbnail.Url,
				Width:  i.Thumbnail.Width,
				Height: i.Thumbnail.Height,
			},
		)
	}

	data, err := json.Marshal(&preview)
	if err != nil {
		panic(err)
	}

	w.Write(data)

}

type preview struct {
	TotalCount int
	Current    int
	Items      []Item
}

type Item struct {
	Url    string
	Width  int
	Height int
}
