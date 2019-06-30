package app

import (
	"sort"

	"git.vsh-labs.cz/cml/nest/src/storage"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

func GetChannelIds() []string {
	return []string{
		"UCbCmjCuTUZos6Inko4u57UQ", //coco melon
		"UC-Gm4EN7nNNR3k67J8ywF4g", // blippi toys
		"UC5PYHgAzJ1wLEidB58SK6Xw", // blippi
		"UCH5TSOM_0xZbbLImGqigtsA", // vroom vroom
	}
}

type Handler struct {
	storage *storage.Handler

	channels     []*Channel
	previewLimit int

	//actualChannel int

	currentIndex int
	items        []*youtube.Item
}

func New(
	channelIds []string,
	storage *storage.Handler,
) (*Handler, error) {
	h := &Handler{
		previewLimit: 4,
	}

	for _, channelId := range channelIds {
		data, err := storage.LoadData(channelId)
		if err != nil {
			return nil, err
		}

		h.channels = append(h.channels,
			NewChannel(channelId, data.Items, data.TotalResults),
		)

		h.items = append(h.items, data.Items...)
	}

	sort.Sort(&itemSorter{
		items: h.items,
	})

	return h, nil
}

func (h *Handler) TotalCount() int {
	return len(h.items)
}

func (h *Handler) CurrentIndex() int {
	return h.currentIndex
}

func (h *Handler) Current() *youtube.Item {
	return h.items[h.currentIndex]
}

func (h *Handler) Next() *youtube.Item {
	h.currentIndex++

	if h.currentIndex >= len(h.items) {
		h.currentIndex = 0
	}

	return h.Current()
}

func (h *Handler) Preview() (result []*youtube.Item) {

	current := h.currentIndex
	for i := 0; i < h.previewLimit; i++ {
		result = append(result, h.items[current])

		current++
		if current >= len(h.items) {
			current = 0
		}
	}

	return
}

//
//func (h *Handler) Current() *youtube.Item {
//	return h.channels[h.actualChannel].Current()
//}
//
//func (h *Handler) Next() *youtube.Item {
//	next := h.channels[h.actualChannel].Next()
//	if next == nil {
//		if h.actualChannel >= len(h.channels)-1 {
//			h.actualChannel = 0
//		} else {
//			h.actualChannel++
//		}
//		return h.channels[h.actualChannel].Reset()
//	}
//	return next
//}
//
//func (h *Handler) Preview() (result []*youtube.Item) {
//
//	result = h.channels[h.actualChannel].Preview(h.previewLimit)
//
//	return
//}
