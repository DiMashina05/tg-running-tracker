package httpapi

import (
	"net/http"

	"github.com/DiMashina05/tg-running-tracker/internal/storage"
)

func NewServer(store storage.Store) http.Handler {
	handler := &Handler{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/users/name", handler.postName)
	mux.HandleFunc("/users/profile", handler.getProfile)
	mux.HandleFunc("/users/runs", handler.postRuns)
	mux.HandleFunc("/users/stats", handler.getStats)

	return mux
}
