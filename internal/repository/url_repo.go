// url_repo.go
package repository

import (
	_ "context"
	"database/sql"
	"errors"
	"time"

	"github.com/fote15/go-url-shortener/internal/models"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(originalURL, shortKey string, userId int) (*models.URL, error) {
	query := `INSERT INTO urls (original, short_key, visits, created_at, user_id)
	          VALUES ($1, $2, 0, NOW(), $3) RETURNING id, created_at`
	var id int64
	var createdAt time.Time

	err := r.db.QueryRow(query, originalURL, shortKey, userId).Scan(&id, &createdAt)
	if err != nil {
		return nil, err
	}
	return &models.URL{
		ID:        id,
		Original:  originalURL,
		ShortKey:  shortKey,
		Visits:    0,
		CreatedAt: createdAt.Format(time.RFC3339),
	}, nil
}

func (r *URLRepository) GetByShortKey(key string) (*models.URL, error) {
	query := `SELECT id, original, short_key, visits, created_at FROM urls WHERE short_key = $1`
	var u models.URL
	err := r.db.QueryRow(query, key).Scan(&u.ID, &u.Original, &u.ShortKey, &u.Visits, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *URLRepository) IncrementVisits(id int64) error {
	_, err := r.db.Exec(`UPDATE urls SET visits = visits + 1 WHERE id = $1`, id)
	return err
}

func (r *URLRepository) Delete(id int64, userID int64) error {
	rs, err := r.db.Exec(`DELETE FROM urls WHERE id = $1 AND user_id = $2`, id, userID)
	if err != nil {
		return err
	}

	affected, err := rs.RowsAffected()
	if err != nil || affected == 0 {
		return errors.New("no rows deleted (maybe not your URL?)")
	}

	return nil
}

func (r *URLRepository) Update(id int64, newOriginal string) error {
	result, err := r.db.Exec(`UPDATE urls SET original = $1 WHERE id = $2`, newOriginal, id)
	rows_count, err := result.RowsAffected()
	if err != nil || rows_count == 0 {
		return errors.New("no rows updated (maybe not your URL?)")
	}
	return err
}

func (r *URLRepository) ListByUserID(userID int) ([]*models.URL, error) {
	rows, err := r.db.Query(`
		SELECT id, original, short_key, visits, created_at 
		FROM urls 
		WHERE user_id = $1 
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.URL
	for rows.Next() {
		u := new(models.URL)
		err := rows.Scan(&u.ID, &u.Original, &u.ShortKey, &u.Visits, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}
	return urls, nil
}
