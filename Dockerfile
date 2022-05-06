FROM golang:1.17

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
COPY ./.config ./.config
RUN go build -o getir

CMD ["./getir"]
