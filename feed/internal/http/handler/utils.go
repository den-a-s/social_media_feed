package handler

import (
	"net/http"
	"social-media-feed/internal/fake"
)

func getUserId(r *http.Request) (int, error) {
	// Должно быть что-то такое, но пока что нет взаимодействия между сервисами, довольствуемся AdminId
	// cookie, err := r.Cookie("user_id")
	// if err != nil {
	// 	return 0, errors.New(fmt.Sprintf("Error in cookie: %s", err))
	// }

	// return id, nil

	return fake.AdminId, nil
}