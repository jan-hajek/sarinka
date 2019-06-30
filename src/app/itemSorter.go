package app

import (
	"strings"

	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type itemSorter struct {
	items []*youtube.Item
}

func (s *itemSorter) Len() int {
	return len(s.items)
}

// Swap is part of sort.Interface.
func (s *itemSorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *itemSorter) Less(i, j int) bool {
	return strings.Compare(s.items[i].Id, s.items[j].Id) > 0
}
