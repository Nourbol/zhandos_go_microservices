package http

import (
	"context"
	"errors"
	"fmt"
	jsonlog "github.com/nourbol/zhandos/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Limiter struct {
	RPS     float64
	Burst   int
	Enabled bool
}

type Options struct {
	Port    int
	Env     string
	Limiter Limiter
}

type Delivery struct {
	logger *jsonlog.Logger
	router http.Handler

	options Options
}

func New(logger *jsonlog.Logger, options Options) *Delivery {
	delivery := &Delivery{
		logger:  logger,
		options: options,
	}

	delivery.router = delivery.routes()
	return delivery
}

func (d *Delivery) Run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", d.options.Port),
		Handler:      d.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		d.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		shutdownError <- srv.Shutdown(ctx)
	}()

	d.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"Env":  d.options.Env,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	d.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
	return nil
}
