FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY task-tracker .

RUN chmod +x task-tracker

COPY migrations ./migrations

EXPOSE 8080

CMD ["./task-tracker"]