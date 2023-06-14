package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	db      *pgxpool.Pool
	options Options
}

type Options struct {
}

func New(db *pgxpool.Pool, options Options) *Repository {
	return &Repository{db: db, options: options}
}
