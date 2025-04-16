#!/bin/sh

echo "⏳ Waiting for Postgres to be ready..."

# Wait for DB to be available
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  >&2 echo "🛑 Postgres is unavailable - sleeping"
  sleep 2
done

echo "✅ Postgres is up!"

echo "🛠 Running DB migrations..."
./migrate

echo "🚀 Starting app..."
./app
