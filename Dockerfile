FROM golang:1.19-alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY . .
RUN go mod tidy
RUN go test -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o bot .


FROM alpine:latest
WORKDIR /app
COPY --from=build /build/bot .

ENV TELEGRAM_API_TOKEN token
ENV SENTRY_DSN dsn
ENV DEBUG_JSON json

CMD ["./bot"]