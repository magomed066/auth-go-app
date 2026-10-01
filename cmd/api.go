package main

import (
	"net/http"
	"time"

	"charm.land/log/v2"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting, analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server has started"))
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 15,
		IdleTimeout: time.Minute,
	}

	log.Info("Server has started at", "addr", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	addr string // PORT
}

