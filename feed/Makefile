build:
	docker build -t social-media-feed .

local_start:
	docker-compose up

prod_start:
	docker run -it -d -p 80:80 social-media-feed

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:6432/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:6432/postgres?sslmode=disable' down
