package handler

import (
	"encoding/json"
	"fmt"
	"ms-go-initial/model"
	"net/http"
)

func ErrorHandler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		w.Header().Add(model.ContentType, model.ApplicationJSON)
		if err != nil {
			if exception, ok := err.(model.RESTCompatible); ok {
				restException := exception.GetException()
				w.WriteHeader(restException.StatusCode)
				_ = json.NewEncoder(w).Encode(&model.ErrorResponse{
					Message: restException.Message,
					Code:    restException.Code,
				})
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&model.ErrorResponse{
				Message: "Internal server error occurred",
				Code:    fmt.Sprintf("%s.internal.server.error", model.Exception),
			})
		}
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error
