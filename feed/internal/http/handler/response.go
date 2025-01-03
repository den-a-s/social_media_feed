package handler

import (
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func(h *Handler) newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	h.logger.Error(message)
	w.WriteHeader(statusCode)
}