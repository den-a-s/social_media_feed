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

	// Доделать
}

func (h *Handler) mainPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
    if err != nil {
        http.Error(w, "cookie 'user_id' не найден", http.StatusUnauthorized)
        return
    }
    // Преобразуем значение cookie в целое число
    userId, err := strconv.Atoi(cookie.Value)
    if err != nil {
        http.Error(w, "некорректный user_id в cookie", http.StatusBadRequest)
        return
    }
	postWithLike, err := h.repo.PostLikeGatewayPostgres.JoinPostWithLike(userId)
	if err != nil {
        http.Error(w, "ошибка при чтении постов с лайками", http.StatusBadRequest)
        return
    }
	

	main_html, err := h.getFilledMainTemplate(posts)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		return
	}

	w.Write([]byte(main_html))
}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	
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