package main

import (
	"log"
	"os"

	"git.vsh-labs.cz/cml/nest/src/http"

	"git.vsh-labs.cz/cml/nest/src/app"
	"git.vsh-labs.cz/cml/nest/src/storage"
)

var App *app.Handler

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	App, err = app.New(app.GetChannelIds(), storage.New(path))
	if err != nil {
		log.Fatal(err)
	}

	http.New(App).Run()
}
