FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o myapp

FROM alpine:latest

WORKDIR /root/

RUN apk add --no-cache libc6-compat

COPY --from=builder /app/myapp .
COPY --from=builder /app/docs ./docs 

RUN chmod +x /root/myapp

EXPOSE 8080

CMD ["./myapp"]