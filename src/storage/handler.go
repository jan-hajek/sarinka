package storage

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"git.vsh-labs.cz/cml/nest/src/youtube"
)

type Handler struct {
	path string
}

func New(path string) *Handler {
	return &Handler{
		path: path,
	}
}

type Data struct {
	Name         string
	Thumbnail    Thumbnail
	TotalResults int
	Items        []*youtube.Item
}

type Thumbnail struct {
	Url string
}

func (h *Handler) LoadData(channelId string) (_ Data, err error) {
	bytes, err := ioutil.ReadFile(path.Join(h.path, channelId+".json"))
	if err != nil {
		return
	}

	var data Data
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return
	}

	return data, err
}

func (h *Handler) SaveData(channelId string, v *Data) (err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(path.Join(h.path, channelId+".json"), data, 0664)
	if err != nil {
		return
	}

	return
}
