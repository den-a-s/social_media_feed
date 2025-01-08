package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"social-media-feed/internal/feed_data"
	"strconv"
)

func (h *Handler) changingLike(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err_str := fmt.Sprintf("[like] Not get url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	like_id := params.Get("like_id")
	post_id := params.Get("post_id")
	int_post_id, err := strconv.Atoi(post_id)
	if err != nil {
		err_str := fmt.Sprintf("[like] Error with converting post_id: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	h.logger.Debug("data:", "like_id", like_id)

	if like_id == "" {
		like := feed_data.Like{PostId: int_post_id, UserId: userId}
		newID, err := h.repo.LikeGateway.Create(like)
		fmt.Println("new id is = ")
		fmt.Println(newID)
		if err != nil {
			h.logger.Error(err.Error())
			http.Error(w, "Error with creating like", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	int_like_id, err := strconv.Atoi(like_id)
	if err != nil {
		err_str := fmt.Sprintf("[like] Error with converting like_id: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	err = h.repo.LikeGateway.Delete(int_like_id)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, "Error with deleting like", http.StatusInternalServerError)
		return
	}

}

// func (h *Handler) deleteLike(w http.ResponseWriter, r *http.Request) {
// 	userId, err := getUserId(r)
// 	if err != nil {
// 		h.newErrorResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	// Доделать
// }
