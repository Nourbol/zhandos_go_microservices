package main

import (
	"flag"
	jsonlog "github.com/nourbol/zhandos/pkg/logger"
	"github.com/nourbol/zhandos/pkg/postgres"
	deliveryHTTP "github.com/nourbol/zhandos/services/shelters/internal/delivery/http"
	"os"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxConns    int
		maxIdleTime string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("ZHANDOS_SHELTERS_DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.db.maxConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	store, err := postgres.New(postgres.Settings{
		DSN:         cfg.db.dsn,
		MaxConns:    cfg.db.maxConns,
		MaxIdleTime: cfg.db.maxIdleTime,
	})
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer store.Pool.Close()
	logger.PrintInfo("database connection pool established", nil)

	var (
		httpListener = deliveryHTTP.New(
			logger,
			deliveryHTTP.Options{
				Port: cfg.port,
				Env:  cfg.env,
				Limiter: deliveryHTTP.Limiter{
					RPS:     cfg.limiter.rps,
					Burst:   cfg.limiter.burst,
					Enabled: cfg.limiter.enabled,
				}})
	)

	err = httpListener.Run()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
