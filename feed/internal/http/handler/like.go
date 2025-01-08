package handler

import (
	"net/http"
	"social-media-feed/internal/fake"
)

func (h *Handler) createLike(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Доделать
}

func (h *Handler) deleteLike(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Доделать
}