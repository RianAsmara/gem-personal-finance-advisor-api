FROM golang:1.22.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates
RUN adduser -D ftuser

COPY --from=build /app/server .
COPY .env .

RUN chown ftuser:ftuser server .env
RUN chmod 755 server
RUN chmod 644 .env

USER ftuser

EXPOSE 9999

CMD ["./server"]