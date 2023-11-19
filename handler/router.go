package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"ms-go-initial/db"
	"net/http"
)

type baseHandler struct{}

func BaseHandler(router *chi.Mux) *baseHandler {
	h := &baseHandler{}
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.StripSlashes)

	router.Get("/readiness", h.Health)
	router.Get("/health", h.Health)
	return h
}

// Health @Summary Health endpoint for kubernetes health and readiness check
// @Tags health-handler
// @Success 200 {} http.Response
// @Router /health [get]
func (*baseHandler) Health(w http.ResponseWriter, r *http.Request) {
	tempDb := db.GetDb().Exec("select 1")
	if tempDb.Error != nil {
		log.Errorf("Got error when selecting 1: %+v", tempDb.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
