package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"

	"git.vsh-labs.cz/cml/nest/src/app"
)

type Handler struct {
	app *app.Handler
}

func New(app *app.Handler) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/current/", h.currentHandler)
	r.HandleFunc("/next/", h.nextHandler)
	r.HandleFunc("/preview/", h.previewHandler)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := path.Dir("./index.html")
		// set header
		w.Header().Set("Content-type", "text/html")
		http.ServeFile(w, r, p)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func (h *Handler) currentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, fmt.Sprintf(`{"videoId":"%s"}`, h.app.Current().Id))
}

func (h *Handler) nextHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, fmt.Sprintf(`{"videoId":"%s"}`, h.app.Next().Id))
}
