package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func MustLoad() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Cannot read .env file: %s", err))
	}
}