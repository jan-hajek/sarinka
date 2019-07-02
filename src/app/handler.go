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

	channels   []*Channel
	allChannel *Channel
}

func New(
	channelIds []string,
	storage *storage.Handler,
) (*Handler, error) {
	h := &Handler{}

	var allItems []*youtube.Item

	for _, channelId := range channelIds {
		data, err := storage.LoadData(channelId)
		if err != nil {
			return nil, err
		}

		h.channels = append(h.channels,
			NewChannel(channelId, data.Name, data.Thumbnail.Url, data.Items, data.TotalResults),
		)

		allItems = append(allItems, data.Items...)
	}

	sort.Sort(&itemSorter{
		items: allItems,
	})

	h.allChannel = NewChannel("", "", "", allItems, len(allItems))

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
