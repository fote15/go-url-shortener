CREATE TABLE urls (
                      id SERIAL PRIMARY KEY,
                      original TEXT NOT NULL,
                      short_key TEXT UNIQUE NOT NULL,
                      visits INTEGER DEFAULT 0,
                      created_at TIMESTAMP DEFAULT NOW()
);
