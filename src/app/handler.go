package app

import (
	"sort"

	"git.vsh-labs.cz/cml/nest/src/storage"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type Handler struct {
	storage *storage.Handler

	channels   []*Channel
	allChannel *Channel
}

func New(
	storage *storage.Handler,
) (*Handler, error) {
	h := &Handler{}

	var allItems []*youtube.Item

	allData, err := storage.LoadAllData()
	if err != nil {
		return nil, err
	}

	for _, data := range allData {
		h.channels = append(h.channels,
			NewChannel(data),
		)

		allItems = append(allItems, data.Items...)
	}

	sort.Sort(&itemSorter{
		items: allItems,
	})

	h.allChannel = NewChannelWithAll(allItems, len(allItems))

	return h, nil
}

func (h *Handler) GetChannel(channelId string) *Channel {
	for _, ch := range h.channels {
		if ch.Id == channelId {
			return ch
		}
	}

	return h.allChannel
}

func (h *Handler) GetChannels() []*Channel {
	return h.channels
}
