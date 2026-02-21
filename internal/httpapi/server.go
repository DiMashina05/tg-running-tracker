package httpapi

import (
	"net/http"

	"github.com/DiMashina05/tg-running-tracker/internal/storage"
	
	swgui "github.com/swaggest/swgui/v5"
)

func NewServer(store storage.Store) http.Handler {
	handler := &Handler{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/users/name", handler.postName)
	mux.HandleFunc("/users/profile", handler.getProfile)
	mux.HandleFunc("/users/runs", handler.postRuns)
	mux.HandleFunc("/users/stats", handler.getStats)
	mux.HandleFunc("/openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/openapi.yml")
	})

	mux.Handle("/swagger/", swgui.NewHandler("tg-running-tracker API", "/openapi.yml", "/swagger/"))
	return mux
}
