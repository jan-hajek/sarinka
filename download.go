package main

import (
	"log"
	"os"

	"git.vsh-labs.cz/cml/nest/src/app"

	"git.vsh-labs.cz/cml/nest/src/downloader"
	"git.vsh-labs.cz/cml/nest/src/storage"
	"git.vsh-labs.cz/cml/nest/src/youtube"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	saveData(storage.New(path), app.GetChannelIds())
}

func saveData(storage *storage.Handler, channelIds []string) {
	youtubeKey := "AIzaSyB26rMJblj16B-QSGApUyvB7qpNSnI4nQI"
	downloader := downloader.New(storage, youtube.New(youtubeKey, 50))
	for _, channelId := range channelIds {
		err := downloader.SaveChannelData(channelId)
		if err != nil {
			log.Fatal(err)
		}
	}
}
