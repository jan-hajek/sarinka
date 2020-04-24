package main

import (
	"flag"
	"log"
	"os"
	"path"

	"git.vsh-labs.cz/jelito/sarinka/src/app"
	"git.vsh-labs.cz/jelito/sarinka/src/downloader"
	"git.vsh-labs.cz/jelito/sarinka/src/http"
	"git.vsh-labs.cz/jelito/sarinka/src/storage"
	"git.vsh-labs.cz/jelito/sarinka/src/youtube"
)

var App *app.Handler

func main() {
	var download = flag.Bool("download", false, "help message for flagname")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.New(path.Join(pwd, "data"))

	if download != nil && *download {
		youtubeKey := "AIzaSyBOi1g2HxI_Y4mhZ6hHwHQlYczWh22IDXw"
		downloader.New(
			storage,
			youtube.New(youtubeKey, 100),
		).SaveData()

		return
	}

	App, err = app.New(storage)
	if err != nil {
		log.Fatal(err)
	}

	http.New(App).Run()

}
