package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Store struct {
	Pool *pgxpool.Pool
}

type Settings struct {
	DSN         string
	MaxConns    int
	MaxIdleTime string
}

func New(settings Settings) (*Store, error) {
	poolCfg, err := configurePool(settings.DSN, settings.MaxConns, settings.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{Pool: conn}, nil
}

func configurePool(dsn string, maxConns int, maxIdleTime string) (*pgxpool.Config, error) {
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	poolCfg.MaxConns = int32(maxConns)
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	poolCfg.MaxConnIdleTime = duration
	return poolCfg, nil
}
