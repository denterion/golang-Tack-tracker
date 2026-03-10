FROM alpine:3.18

WORKDIR /app

COPY task-tracker .

COPY migrations ./migrations

EXPOSE 8080

CMD ["./task-tracker"]