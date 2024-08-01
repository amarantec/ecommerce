FROM docker.io/library/golang:1.22.5
 
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
 
RUN go build -o app cmd/api/main.go
 
CMD ["./app"]

