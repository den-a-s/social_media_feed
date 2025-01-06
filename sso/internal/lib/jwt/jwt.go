package jwt

import (
	"sso/internal/domain/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: подумать где его хранить
const secret = "test-secret"

func NewToken(user model.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
