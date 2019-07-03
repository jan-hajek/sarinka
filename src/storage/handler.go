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

func (h *Handler) LoadData(channelId string) (_ youtube.Result, err error) {
	bytes, err := ioutil.ReadFile(path.Join(h.path, channelId+".json"))
	if err != nil {
		return
	}

	var data youtube.Result
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return
	}

	return data, err
}

func (h *Handler) SaveData(channelId string, v youtube.Result) (err error) {
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
