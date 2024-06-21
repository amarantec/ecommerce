FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

COPY . .

EXPOSE 8080

CMD [ "go", "run", "cmd/api/main.go" ]