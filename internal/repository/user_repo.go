package repository

import (
	"database/sql"
	"errors"

	"github.com/fote15/go-url-shortener/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`
	return r.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `SELECT id, email, password FROM users WHERE email = $1`
	var user models.User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
