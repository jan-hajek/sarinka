package main

import (
	"flag"
	"log"
	"os"

	"git.vsh-labs.cz/cml/nest/src/downloader"
	"git.vsh-labs.cz/cml/nest/src/youtube"

	"git.vsh-labs.cz/cml/nest/src/http"

	"git.vsh-labs.cz/cml/nest/src/app"
	"git.vsh-labs.cz/cml/nest/src/storage"
)

var App *app.Handler

func main() {
	var download = flag.Bool("download", false, "help message for flagname")
	flag.Parse()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.New(path)

	if download != nil && *download {
		youtubeKey := "AIzaSyB26rMJblj16B-QSGApUyvB7qpNSnI4nQI"
		downloader.New(
			storage,
			youtube.New(youtubeKey, 50),
		).SaveData(app.GetChannelIds())

		return
	}

	App, err = app.New(app.GetChannelIds(), storage)
	if err != nil {
		log.Fatal(err)
	}

	http.New(App).Run()

}
