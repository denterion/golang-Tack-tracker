# Берём минимальный образ
FROM alpine:3.18

WORKDIR /app

# Копируем уже собранный бинарник
COPY task-tracker .

# Копируем папку с миграциями (если есть)
COPY migrations ./migrations

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["./task-tracker"]