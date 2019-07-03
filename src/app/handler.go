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
		"UCgNqMjgfN6YQYCYOxhp78xQ", // pisnicky pro deti
		"UCyTOz12LseUJhDtl8K6-vFw", // baby studio
		"UCnUExytdcrLyl_ccqe2dxpg", // ceske pohadky
		"UCVGHomAxlCzmBH7T5SN0_ag", // ciperkove
		"UC20HC7Rj2_pyyUq4xAjtw0Q", // baby smile
		"UCRdYoBTTupd3svI9kV36UvQ", // nakladak
		"UCjuM-aDNfzbVShImGfj0WOA", // odtahove auto
		"UClroLesWYk7cc9Tctv8HrkA", // monster auto
		"UCZBqWU1GgUHBTPZ-FMPZ8-Q", // byl jednou
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
