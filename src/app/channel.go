package app

import (
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type Channel struct {
	id         string
	downloaded []*youtube.Item
	totalCount int

	ids    map[string]struct{}
	actual int
}

func NewChannel(id string, downloaded []*youtube.Item, totalCount int) *Channel {
	ch := &Channel{
		id:         id,
		ids:        make(map[string]struct{}),
		totalCount: totalCount,
	}

	ch.add(downloaded)

	return ch
}
func (ch *Channel) add(items []*youtube.Item) {
	for _, i := range items {
		if _, exists := ch.ids[i.Id]; !exists {
			ch.downloaded = append(ch.downloaded, i)
			ch.ids[i.Id] = struct{}{}
		} else {
			println("duplicate entry,", i.Id)
		}
	}
}

func (ch *Channel) Current() *youtube.Item {
	if ch.actual >= len(ch.downloaded) {
		return nil
	}

	return ch.downloaded[ch.actual]
}

func (ch *Channel) Next() *youtube.Item {
	ch.actual++

	return ch.Current()
}

func (ch *Channel) Reset() *youtube.Item {
	ch.actual = 0

	return ch.Current()
}

func (ch *Channel) Preview(limit int) (res []*youtube.Item) {
	limit = ch.actual + limit
	if limit > len(ch.downloaded) {
		limit = len(ch.downloaded)
	}

	return ch.downloaded[ch.actual:limit]
}
