FROM golang:1.26

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/http/main.go"]