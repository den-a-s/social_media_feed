package handler

import (
	"fmt"
	"net/http"
	"social-media-feed/internal/fake"
	//"strconv"
)

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
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
	// cookie, err := r.Cookie("user_id")
    // if err != nil {
    //     http.Error(w, "cookie 'user_id' не найден", http.StatusUnauthorized)
    //     return
    // }
    // // Преобразуем значение cookie в целое число
    // userId, err := strconv.Atoi(cookie.Value)
    // if err != nil {
    //     http.Error(w, "некорректный user_id в cookie", http.StatusBadRequest)
    //     return
    // }
	postWithLike, err := h.repo.PostsWithLikeGateway.JoinPostWithLike(fake.AdminId)
	if err != nil {
        http.Error(w, fmt.Sprintf("ошибка при чтении постов с лайками: %s", err), http.StatusBadRequest)
        return
    }

	main_html, err := h.getFilledMainTemplate(postWithLike)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("Не смогли получить сформированный шаблон: %s", err))
		return
	}

	w.Write([]byte(main_html))
}

func (h *Handler) getPostById(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) updatePost(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
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