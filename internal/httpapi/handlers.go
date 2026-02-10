package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/DiMashina05/tg-running-tracker/internal/service"
	"github.com/DiMashina05/tg-running-tracker/internal/storage"
)

type Handler struct {
	store storage.Store
}

func (h *Handler) postName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req postNameReqest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		service.SetName(h.store, req.Name, req.UserID)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getprofile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *Handler) postRuns(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
