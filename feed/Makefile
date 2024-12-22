setup:
	go mod tidy

local_start:
	CONFIG_PATH="./config/local.yaml" go run cmd/feed/main.go

prod_start:
	CONFIG_PATH="./config/prod.yaml" go run cmd/feed/main.go

start:
	setup
	prod_start