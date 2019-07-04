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

func (h *Handler) LoadAllData() (res []youtube.Result, err error) {
	files, err := ioutil.ReadDir(h.path)
	if err != nil {
		return
	}

	for _, f := range files {
		data, err := h.loadFile(path.Join(h.path, f.Name()))
		if err != nil {
			return res, err
		}
		res = append(res, data)
	}

	return
}

func (h *Handler) loadFile(path string) (_ youtube.Result, err error) {
	bytes, err := ioutil.ReadFile(path)
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
