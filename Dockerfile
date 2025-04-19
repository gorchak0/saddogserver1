# Используем официальный образ Go
FROM golang:1.20-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Собираем приложение
RUN go build -o app .

# Используем более лёгкий образ для запуска
FROM alpine:latest

# Копируем собранное приложение из builder
COPY --from=builder /app/app /app

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app"]
