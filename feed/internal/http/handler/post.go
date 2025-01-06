package handler

import (
	"fmt"
	"net/http"
	"social-media-feed/internal/fake"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !fake.IsAdmin(userId) {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.ServeFile(w, r, "web/templates/formPost.html")
	//w.Write([]byte(formPost))
	// Доделать
}

func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.PostGateway.GetAll()
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	main_html, err := h.getFilledMainTemplate(posts)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		return
	}

	w.Write([]byte(main_html))
}


func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
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