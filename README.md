### Реализация новостной ленты

Есть посты, пользователи и реакции на постах (расписать поподробнее данные фичи)

Для локальной разработки подкиньте нужный файл конфигурации. Обычно его называют local.yaml
Подкинуть его нужно в папку config/

Для того чтобы запустить приложение необходимо создать .env файл в котором указать следующие переменные окружения

```
DB_PASSWORD=<пароль от базы>
# или prod.yaml
CONFIG_PATH="./config/local.yaml"
```

Для запуска использовать

```
make build && make local_start # или _prod
```
