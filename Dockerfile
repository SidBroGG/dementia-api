# Stage 1 building
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o todo-api ./cmd/api/main.go

# Stage 2 running
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todo-api .

EXPOSE 8080

CMD [ "./todo-api" ]