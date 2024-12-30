stop:
	docker-compose down

start:
	docker-compose up -d --build social-media-feed

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:6432/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:6432/postgres?sslmode=disable' down
