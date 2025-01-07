FROM golang:alpine3.21

ENV SERVICE_NAME=sso

RUN apk --update upgrade && \
    apk add sqlite musl-dev gcc && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY . .
COPY --from=protos . ../protos

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o bin/${SERVICE_NAME} ./cmd/${SERVICE_NAME}

EXPOSE 44044
# Интерполяция не работает в exec режиме (массив аргументов)
RUN go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
CMD ./bin/$SERVICE_NAME --config=./config/local.yaml
