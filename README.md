1. В cmd лежит основной main файл (точка входа)
2. В pkg хранится вся логика нашего приложения
3. Установка Gin фреймворка для эндпоинтов
   go get -u github.com/gin-gonic/gin
4. Для работы с файлом конфигурации устанавливаем пакет
   go get -u github.com/spf13/viper
5. Создание инстанса docker postgres
   docker run --name=rest-go -e POSTGRES_PASSWORD=123123 -p 5432:5432 -d --rm postgres
6. Пакет для миграций
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
7. Инициализация миграций
   migrate create -ext sql -dir ./schema -seq init
8. Выполнить миграции
   migrate -path ./schema -database 'postgres://postgres:123123@localhost:5432/postgres?sslmode=disable' up
9. Откат миграции
   migrate -path ./schema -database 'postgres://postgres:123123@localhost:5432/postgres?sslmode=disable' down
10. Подключение к контейнеру Postgres Docker
    docker exec -ti rest-go /bin/bash
    psql -U postgres
    вывести все таблицы \d
11. Установка пакета для работы с БД
    go get -u github.com/jmoiron/sqlx
12. Для работы с переменными окружения .ENV установим пакет
    go get -u github.com/joho/godotenv
