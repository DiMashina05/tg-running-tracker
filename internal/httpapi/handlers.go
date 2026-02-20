package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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

		name, err := service.SetName(h.store, req.Name, req.UserID)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrAlreadyRegistered):

				http.Error(w, err.Error(), http.StatusConflict)

			case errors.Is(err, service.ErrInvalidName):

				http.Error(w, err.Error(), http.StatusBadRequest)

			default:

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		resp := postNameReqest{Name: name, UserID: req.UserID}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h *Handler) getProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idStr := r.URL.Query().Get("user_id")

		userID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}

		name := h.store.GetName(userID)

		if name == "" {
			http.Error(w, "Not registered", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		resp := getProfileResponce{UserID: userID, Name: name}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func (h *Handler) postRuns(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req postRunsReqest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		distance, err := service.AddRun(h.store, req.Dist, req.UserID)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrInvalidDistance):

				http.Error(w, err.Error(), http.StatusBadRequest)

			case errors.Is(err, service.ErrNotRegistered):

				http.Error(w, err.Error(), http.StatusNotFound)

			default:

				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		resp := postRunsResponse{UserID: req.UserID, Dist: distance}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println(err)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idStr := r.URL.Query().Get("user_id")

		userID, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid user_id", http.StatusBadRequest)
			return
		}

		stats, err := service.GetStats(h.store, userID)

		if err != nil {
			switch {
			case errors.Is(err, service.ErrNotRegistered):
				http.Error(w, err.Error(), http.StatusNotFound)

			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(stats); err != nil {
			log.Println(err)
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
