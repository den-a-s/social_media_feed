### Репозиторий с proto-файлами

Здесь лежит контракт сервиса авторизации. Для генерации нужно сначала уставновить две библиотеки:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

После нужно обновить PATH:
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

Теперь можем выполнить генерацию командой (из директории ./protos)
```
protoc -I proto proto/sso/sso.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
```