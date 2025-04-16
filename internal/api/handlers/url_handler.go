package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/fote15/go-url-shortener/internal/api/middleware"
	"github.com/fote15/go-url-shortener/internal/repository"
	"github.com/fote15/go-url-shortener/internal/utils"
	"net/http"

	"github.com/gorilla/mux"
	"strconv"
)

type UrlHandler struct {
	repo *repository.URLRepository
}

func NewUrlHandler(db *sql.DB) *UrlHandler {
	repo := repository.NewURLRepository(db)
	return &UrlHandler{repo: repo}
}

type shortenRequest struct {
	Original string `json:"original_url"`
	Custom   string `json:"custom_key"`
}

type shortenResponse struct {
	Short string `json:"short_url"`
}

func (h *UrlHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest

	userID, ok := r.Context().Value(middlewares.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err, newOriginal := utils.IsValidURL(req.Original)
	if err != nil {
		http.Error(w, "Invalid URL"+newOriginal, http.StatusBadRequest)
		return
	}
	req.Original = newOriginal

	key := req.Custom
	if key == "" {
		key = utils.GenerateKey()
	}
	url, err := h.repo.Create(req.Original, key, userID)
	if err != nil {
		http.Error(w, "key most likely already exists", http.StatusConflict)
		return
	}

	resp := shortenResponse{Short: "/" + url.ShortKey}
	json.NewEncoder(w).Encode(resp)
}

func (h *UrlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["shortKey"]
	url, err := h.repo.GetByShortKey(key)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	go h.repo.IncrementVisits(url.ID)
	http.Redirect(w, r, url.Original, http.StatusFound)
}

func (h *UrlHandler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err = h.repo.Delete(int64(id), int64(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UrlHandler) EditURL(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID ", http.StatusBadRequest)
		return
	}
	var req shortenRequest
	json.NewDecoder(r.Body).Decode(&req)
	err, newOriginal := utils.IsValidURL(req.Original)
	if err != nil {
		http.Error(w, "Invalid URL"+err.Error(), http.StatusBadRequest)
		return
	}
	req.Original = newOriginal
	err = h.repo.Update(int64(id), req.Original)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Data string `json:"data"`
	}{Data: "OK"})
}

func (h *UrlHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]
	url, err := h.repo.GetByShortKey(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(url)
}

func (h *UrlHandler) ListUserURLs(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middlewares.UserIDKey).(int)
	if !ok {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
		return
	}

	list, err := h.repo.ListByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}
