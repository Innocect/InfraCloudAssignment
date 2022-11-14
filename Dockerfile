FROM golang:1.17.2-alpine3.14 AS Assignment
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main cmd/main.go
EXPOSE 2802
CMD [ "/app/main" ]