package http

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"

	"git.vsh-labs.cz/cml/nest/src/app"
)

type Handler struct {
	app *app.Handler

	mx *sync.Mutex
}

func New(app *app.Handler) *Handler {
	return &Handler{
		app: app,
		mx:  &sync.Mutex{},
	}
}

func (h *Handler) Run() {
	r := mux.NewRouter()

	// html
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, "./www/homepage.html")
	})
	r.HandleFunc("/play/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, "./www/play.html")
	})

	// rest
	r.HandleFunc("/current/", h.currentHandler)
	r.HandleFunc("/preview/", h.previewHandler)
	r.HandleFunc("/channels/", h.channelsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getId(r *http.Request) (id string) {
	return r.URL.Query().Get("id")
}

func getChannelId(r *http.Request) (channelId string) {
	return r.URL.Query().Get("channelId")
}
