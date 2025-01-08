package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"social-media-feed/internal/fake"
	"strconv"

	ssov1 "github.com/username/protos/gen/go/sso"
)

func (h *Handler) auth(w http.ResponseWriter, r *http.Request) {
	auth_html, err := os.ReadFile("./web/templates/auth.html")
	if err != nil {
		err_str := fmt.Sprintf("[auth] Not read file: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	w.Write([]byte(auth_html))
}

func (h *Handler) registrate(w http.ResponseWriter, r *http.Request) {

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err_str := fmt.Sprintf("[registrate] Not get url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	login := params.Get("login")
	password := params.Get("password")

	h.logger.Debug("data:", "login", login, "password", password)

	if login == "" || password == "" {
		err_str := fmt.Sprintf("[registrate] Not parse url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	resp, err := h.authClient.Register(r.Context(), &ssov1.RegisterRequest{Email: login, Password: password})
	if err != nil {
		err_str := fmt.Sprintf("[registrate] Not get response before registrate: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	h.logger.Debug("data:", "resp.UserId", resp.UserId)
	h.logger.Debug("data stringed:", "resp.UserId", strconv.FormatInt(resp.UserId, 10))

	cookie := http.Cookie{Name: "user_id", Value: strconv.FormatInt(resp.UserId, 10)}

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		err_str := fmt.Sprintf("[login] Not get url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	login := params.Get("login")
	password := params.Get("password")

	h.logger.Debug("data:", "login", login, "password", password)

	if login == "" || password == "" {
		err_str := fmt.Sprintf("[login] Not parse url params: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	resp, err := h.authClient.Login(r.Context(), &ssov1.LoginRequset{Email: login, Password: password})
	if err != nil {
		err_str := fmt.Sprintf("[login] Not get response before registrate: %s", err)
		h.newErrorResponse(w, http.StatusInternalServerError, err_str)
		return
	}

	h.logger.Debug("[login] data:", "resp.Token", resp.Token)

	cookie_token := http.Cookie{Name: "token", Value: resp.Token}
	// TODO поменять на реальный user_id
	cookie_user_id := http.Cookie{Name: "user_id", Value: strconv.FormatInt(fake.AdminId, 10)}

	http.SetCookie(w, &cookie_token)
	http.SetCookie(w, &cookie_user_id)

	w.WriteHeader(http.StatusOK)
}
