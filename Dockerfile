# Используем базовый образ с Go
FROM golang:latest AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN go build -o main main.go

EXPOSE 8080


# Команда для запуска
CMD ["./main", "-s", "database"]
