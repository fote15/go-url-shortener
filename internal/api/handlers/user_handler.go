package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/fote15/go-url-shortener/internal/models"
	"github.com/fote15/go-url-shortener/internal/repository"
	"github.com/fote15/go-url-shortener/internal/utils"
	_ "github.com/fote15/go-url-shortener/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{
		repo: repository.NewUserRepository(db),
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Basic validation
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	// Email format validation
	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashed)

	if err := h.repo.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID, // assuming you assign an ID (e.g., from DB)
		"email": user.Email,
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByEmail(input.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
