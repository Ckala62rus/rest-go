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
13. Установка MySQL пакета (драйвера)
    go get -u github.com/go-sql-driver/mysql
14. Установка пакета для логирования Logrus
    go get -u github.com/sirupsen/logrus
15. Установка пакета для JWT аутентификации
    go get -u github.com/dgrijalva/jwt-go

Примечание:
для смены бд postgres на mysql необходимо скачать драйвер, создать новую обертку для репозитория,
а так же выполнить эти команды(для конкретной бд). Без них будет ругаться, что драйвера нет!
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

выполнение миграции для mysql
migrate -path ./schema -database 'mysql://root:000000@tcp(0.0.0.0:3306)/go?query' up

обновление модулей(пакетов)
go get -u all

docker pull postgres
