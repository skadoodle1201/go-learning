package handlers

import (
	"github.com/go-chi/chi"
	chimiddler "github.com/go-chi/chi/middleware"
	"github.com/skadoodle1201/goAPI/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddler.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
