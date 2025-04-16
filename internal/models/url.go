package models

type URL struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	Original  string `json:"original_url"`
	ShortKey  string `json:"short_key"`
	Visits    int64  `json:"visits"`
	CreatedAt string `json:"created_at"`
}
