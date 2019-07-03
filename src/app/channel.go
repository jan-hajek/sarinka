package app

import (
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type Channel struct {
	Id           string
	Name         string
	ThumbnailUrl string
	downloaded   []*youtube.Item
	totalCount   int

	ids map[string]int
}

func NewChannel(result youtube.Result) *Channel {
	ch := &Channel{
		Id:           result.Channel.Id,
		Name:         result.Channel.Title,
		ThumbnailUrl: result.Channel.Thumbnail.Url,
		ids:          make(map[string]int),
		totalCount:   result.TotalResults,
	}

	ch.add(result.Items)

	return ch
}

func NewChannelWithAll(items []*youtube.Item, totalCount int) *Channel {
	ch := &Channel{
		ids:        make(map[string]int),
		totalCount: totalCount,
	}

	ch.add(items)

	return ch
}

func (ch *Channel) add(items []*youtube.Item) {
	for _, i := range items {
		if i.Id == "" {
			continue
		}
		if _, exists := ch.ids[i.Id]; !exists {
			ch.downloaded = append(ch.downloaded, i)
			ch.ids[i.Id] = len(ch.downloaded) - 1
		} else {
			println("duplicate entry,", i.Id)
		}
	}
}

func (ch *Channel) GetDownloadedCount() int {
	return len(ch.downloaded)
}

func (ch *Channel) getPosition(id string) int {
	if pos, exists := ch.ids[id]; exists {
		return pos
	}
	return 0
}

func (ch *Channel) GetItems(startId string, limit int) (position int, res []*youtube.Item) {
	position = ch.getPosition(startId)

	x := position
	for i := 0; i < limit; i++ {
		res = append(res, ch.downloaded[x])
		x++
		if x == len(ch.downloaded) {
			x = 0
		}
	}
	return
}
