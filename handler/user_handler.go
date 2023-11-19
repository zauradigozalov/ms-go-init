package handler

import (
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"ms-go-initial/config"
	"ms-go-initial/dto"
	"ms-go-initial/service"
	"ms-go-initial/util"
	"net/http"
	"strconv"
)

type handler struct {
	service service.IService
}

func Handler(router *chi.Mux, service service.IService) *chi.Mux {
	router.Mount(config.RootPath, Router(service))
	return router
}

func Router(service service.IService) http.Handler {
	h := &handler{service: service}
	r := chi.NewRouter()

	r.Post("/users", ErrorHandler(h.saveUser))
	r.Put("/users/{userId}", ErrorHandler(h.updateUser))

	return r
}

func (h *handler) saveUser(w http.ResponseWriter, r *http.Request) error {

	var userDTO dto.UserRequest

	if err := util.DecodeBody(r, &userDTO); err != nil {
		return err
	}

	err := h.service.CreateUser(userDTO.UserName, userDTO.Status)

	if err != nil {
		log.Error("ActionLog.saveUser.error ", err)
		return err
	}

	util.HandleResponse(w, nil, http.StatusCreated)
	return err

}

func (h *handler) updateUser(w http.ResponseWriter, r *http.Request) error {

	var userDTO dto.UserRequest

	userId, _ := strconv.ParseInt(chi.URLParam(r, "userId"), 10, 64)

	if err := util.DecodeBody(r, &userDTO); err != nil {
		return err
	}

	err := h.service.UpdateUser(uint(userId), userDTO.UserName, userDTO.Status)

	if err != nil {
		log.Error("ActionLog.updateUser.error ", err)
		return err
	}

	util.HandleResponse(w, nil, http.StatusOK)
	return err
}
